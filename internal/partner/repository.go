package partner

import (
	"context"

	"github.com/caiowWillian/partner-service/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodrive "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type Repository interface {
	CreatePartner(partner PartnerPostRequest) (string, error)
	GetById(id string) (Partner, error)
	GetNearPartner(latLong []float64) (Partner, error)
	CreateIndexes() error
}

type repository struct {
	repo mongo.Repository
}

var settings = mongo.MongoOperation{
	Database:   "test",
	Collection: "partner",
}

func NewRepository(repo mongo.Repository) Repository {
	return &repository{repo: repo}
}

func (r *repository) CreateIndexes() error {

	index := []mongodrive.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "document", Value: bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bsonx.Doc{{Key: "address", Value: bsonx.String("2dsphere")}},
		},
	}

	_, err := r.repo.GetCollection(settings).Indexes().CreateMany(context.Background(), index)

	return err
}

func (r *repository) CreatePartner(partner PartnerPostRequest) (string, error) {
	result, err := r.repo.GetCollection(settings).InsertOne(context.Background(), partner)

	if err != nil {
		return "", err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func (r *repository) GetById(id string) (Partner, error) {
	var partner Partner
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return partner, err
	}

	filter := bson.D{{"_id", objectId}}
	err = r.repo.GetCollection(settings).FindOne(context.Background(), filter).Decode(&partner)
	return partner, err
}

func (r *repository) GetNearPartner(latLong []float64) (Partner, error) {
	var partner Partner

	filter := bson.M{
		"$and": []bson.M{
			{
				"coverageArea": bson.M{
					"$geoIntersects": bson.M{
						"$geometry": bson.M{
							"type":        "Point",
							"coordinates": latLong,
						},
					},
				},
			},
			{
				"address": bson.M{
					"$near": bson.M{
						"$geometry": bson.M{
							"type":        "Point",
							"coordinates": latLong,
						},
					},
				},
			},
		},
	}
	err := r.repo.GetCollection(settings).FindOne(context.Background(), filter).Decode(&partner)
	return partner, err
}
