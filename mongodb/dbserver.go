package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbCollection *mongo.Database

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://user:123456@localhost:27017/biliob?authMechanism=SCRAM-SHA-256")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	DbCollection = client.Database("biliob")
}
