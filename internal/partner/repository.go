package partner

import (
	"context"

	"github.com/caiowWillian/partner-service/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	CreatePartner(partner Partner) error
	GetById(id string) (Partner, error)
	GetNearPartner(latLong []float32) (Partner, error)
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

func (r *repository) CreatePartner(partner Partner) error {
	_, err := r.repo.GetCollection(settings).InsertOne(context.Background(), partner)

	return err
}

func (r *repository) GetById(id string) (Partner, error) {
	var partner Partner
	filter := bson.D{{"id", id}}
	err := r.repo.GetCollection(settings).FindOne(context.Background(), filter).Decode(&partner)
	return partner, err
}

func (r *repository) GetNearPartner(latLong []float32) (Partner, error) {
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
