package mongodb

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var DbCollection *mongo.Database

func SetDb(url string, dbName string) {
	clientOptions := options.Client().ApplyURI(url)
	clientOptions.SetConnectTimeout(3 * time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	DbCollection = client.Database(dbName)
	logrus.Debug("dbUrl is:" + url + " dbName is:" + dbName)
}
