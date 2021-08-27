package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetCollection(operation MongoOperation) *mongo.Collection
}

type MongoOperation struct {
	Database   string
	Collection string
}

type repository struct {
	db  *mongo.Client
	err error
}

var repo Repository

func NewMongo() Repository {
	options := options.Client().ApplyURI("mongodb://localhost:27017/teste")
	client, err := mongo.Connect(context.TODO(), options)
	return &repository{client, err}
}

func init() {

	options := options.Client().ApplyURI("mongodb://localhost:27017/teste")
	client, err := mongo.Connect(context.TODO(), options)

	repo = &repository{client, err}
}

func (repo *repository) GetCollection(operation MongoOperation) *mongo.Collection {
	return (repo.db.Database(operation.Database)).Collection(operation.Collection)
}

func Repo() Repository {
	return repo
}
