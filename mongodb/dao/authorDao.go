package dao

import (
	"bspider/mongodb"
	"bspider/mongodb/dao/model"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func UpsertAuthorToDb(model model.AuthorDo) {
	author := mongodb.DbCollection.Collection("author")
	if model.CFans != 0 {
		filter := bson.M{"mid": model.Mid}
		update := bson.M{
			"$set": bson.M{
				"focus":         true,
				"sex":           model.Sex,
				"name":          model.Name,
				"face":          model.Face,
				"level":         model.Level,
				"cFans":         model.CFans,
				"cLike":         model.CLike,
				"cRate":         model.CRate,
				"official":      model.Official,
				"c_attention":   model.CAttention,
				"c_archiveView": model.CArchiveView,
				"c_articleView": model.CArticleView,
			},
			"$push": bson.M{
				"data": bson.M{
					"$each":     bson.A{model.Data},
					"$position": 0,
				},
			},
		}
		op := options.Update()
		op.SetUpsert(true)
		_, err := author.UpdateOne(context.TODO(), filter, update, op)
		if err != nil {
			panic(err)
		}
		logrus.Info("upsert author into mongodb ", model)
	}
}

func UpsertChannelToDb(model model.VideoWithMidAidChannelsDo) {
	author := mongodb.DbCollection.Collection("author")
	filter := bson.M{
		"mid": model.Mid,
	}
	update := bson.M{
		"$set": bson.M{
			"channels": model.Channels,
		},
	}
	option := options.Update()
	option.SetUpsert(true)
	author.UpdateOne(context.TODO(), filter, update, option)
	logrus.Info("upsert video channel to mongodb ", model)
}

func UpsertRankToDb(rank map[string]interface{}, mid string) {
	author := mongodb.DbCollection.Collection("author")
	filter := bson.M{
		"mid": mid,
	}
	update := bson.M{
		"$set": bson.M{
			"rank": rank,
		},
	}
	op := options.Update()
	op.SetUpsert(true)
	_, err := author.UpdateOne(context.TODO(), filter, update, op)
	if err != nil {
		panic(err)
	}
	logrus.Info("upsert author rank to mongodb ", rank)
}

func SelectMidForAuthorByFucus() []string {
	collection := mongodb.DbCollection.Collection("author")
	filter := bson.M{"$or": bson.A{
		bson.M{"focus": true},
		bson.M{"forceFocus": true}},
	}
	op := options.Find()
	op.SetProjection(bson.M{"mid": 1, "_id": 0})
	result, err := collection.Find(context.TODO(), filter, op)

	if err != nil {
		panic(err)
	}

	var rt []string
	for result.Next(context.TODO()) {
		var s interface{}
		err := result.Decode(&s)
		if err != nil {
			panic(err)
		}
		rt = append(rt, s.(primitive.D).Map()["mid"].(string))
	}
	return rt
}

func AggregateForAuthorByMid(model model.AuthorDo) []primitive.M {
	author := mongodb.DbCollection.Collection("author")

	agg := bson.A{
		bson.M{"$match": bson.M{"mid": model.Mid}},
		bson.M{"$unwind": "$data"},
		bson.M{"$match": bson.M{"data.datetime": bson.M{"$gt": time.Now().AddDate(0, 0, -1).Unix()}}},
		bson.M{"$sort": bson.M{"data.datetime": 1}},
		bson.M{"$limit": 1},
		bson.M{"$project": bson.M{"datetime": "$data.datetime", "like": "$data.like", "fans": "$data.fans", "archiveView": "$data.archiveView", "articleView": "$data.articleView"}},
	}
	result, err := author.Aggregate(context.TODO(), agg)
	if err != nil {
		panic(err)
	}

	var rt []primitive.M
	for result.Next(context.TODO()) {
		var s interface{}
		err := result.Decode(&s)
		if err != nil {
			panic(err)
		}
		rt = append(rt, s.(primitive.D).Map())
	}
	return rt
}

func CountAuthorCountByKey(key string) int64 {
	author := mongodb.DbCollection.Collection("author")
	result, err := author.CountDocuments(context.TODO(), bson.M{key: bson.M{"$exists": 1}})
	if err != nil {
		panic(err)
	}
	return result
}

func SelectKey(key string) []map[string]interface{} {
	author := mongodb.DbCollection.Collection("author")
	op := options.Find()
	op.SetProjection(bson.M{"mid": 1, "rank": 1, key: 1})
	op.SetSort(bson.M{key: -1})
	op.SetBatchSize(300)
	result, err := author.Find(context.TODO(), bson.M{key: bson.M{"$exists": 1}}, op)

	if err != nil {
		panic(err)
	}

	var rt []map[string]interface{}
	for result.Next(context.TODO()) {
		var s interface{}
		err := result.Decode(&s)
		if err != nil {
			panic(err)
		}
		rt = append(rt, s.(primitive.D).Map())
	}
	return rt
}
