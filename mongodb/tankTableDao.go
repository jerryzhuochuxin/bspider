package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func UpsertRankTableToDb(name string, object map[string]interface{}) {
	rankTable := DbCollection.Collection("rank_table")
	op := options.Update()
	op.SetUpsert(true)
	_, err := rankTable.UpdateOne(context.TODO(), bson.M{"name": name}, bson.M{"$set": object}, op)
	if err != nil {
		panic(err)
	} else {
		log.Printf("upsert ranktabl into mongodb success object is %v", object)
	}
}
