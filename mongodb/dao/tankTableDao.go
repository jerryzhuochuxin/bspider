package dao

import (
	"bspider/mongodb"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertRankTableToDb(name string, object map[string]interface{}) {
	rankTable := mongodb.DbCollection.Collection("rank_table")
	op := options.Update()
	op.SetUpsert(true)
	_, err := rankTable.UpdateOne(context.TODO(), bson.M{"name": name}, bson.M{"$set": object}, op)
	if err != nil {
		panic(err)
	} else {
		logrus.Info("upsert ranktabl into mongodb success object is ", object)
	}
}
