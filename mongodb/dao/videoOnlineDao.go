package dao

import (
	"bspider/mongodb"
	"bspider/mongodb/dao/model"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertVideoOnlineToDb(online model.VideoOnlineDo) {
	videoOnline := mongodb.DbCollection.Collection("video_online")
	filter := bson.M{
		"title": online.Title,
	}
	update := bson.M{
		"$set": bson.M{
			"title":      online.Title,
			"author":     online.Author,
			"channel":    online.Channel,
			"subChannel": online.SubChannel,
		},
		"$addToSet": bson.M{
			"data": online.Data,
		},
	}
	op := options.Update()
	op.SetUpsert(true)

	_, err := videoOnline.UpdateOne(context.TODO(), filter, update, op)
	if err != nil {
		panic(err)
	}
	logrus.Info("upsert videoOnline into mongodb ", online)
}
