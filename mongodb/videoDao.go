package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"bspider/model"
	"log"
	"time"
)

func UpsertVideoWithMidAidChannelsToDb(model model.VideoWithMidAidChannelsDo) []string {
	urlPre := "https://api.bilibili.com/x/article/archives?ids="
	var rt []string
	video := DbCollection.Collection("video")
	for _, aid := range model.Aid {
		filter := bson.M{
			"aid": aid,
		}
		update := bson.M{
			"$set": bson.M{
				"aid":   aid,
				"focus": true,
			},
		}
		option := options.Update()
		option.SetUpsert(true)

		_, err := video.UpdateOne(context.TODO(), filter, update, option)
		if err != nil {
			panic(err)
		}

		rt = append(rt, urlPre+aid)
		log.Printf("upsert video into mongodb success aid is %s", aid)
	}
	return rt
}

func UpsertVideoToDb(modelList []model.VideoDo) {
	video := DbCollection.Collection("video")
	for _, model := range modelList {
		filter := bson.M{
			"aid": model.Aid,
		}
		update := bson.M{
			"$set": bson.M{
				"cView":      model.CurrentView,
				"cFavorite":  model.CurrentFavorite,
				"cDanmaku":   model.CurrentDanmaku,
				"cCoin":      model.CurrentCoin,
				"cShare":     model.CurrentShare,
				"cLike":      model.CurrentLike,
				"cDatetime":  model.CurrentDatetime,
				"author":     model.Author,
				"subChannel": model.SubChannel,
				"channel":    model.Channel,
				"mid":        model.Mid,
				"pic":        model.Pic,
				"title":      model.Title,
				"datetime":   time.Now().Unix(),
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
		_, err := video.UpdateOne(context.TODO(), filter, update, op)
		if err != nil {
			panic(err)
		}
		log.Printf("upsert video into mongodb success video is %v", model)
	}
}

func SelectAidListForVideoUpdate(focus bool) []string {
	var rt []string
	video := DbCollection.Collection("video")
	filter := bson.M{
		"focus": focus,
	}
	op := options.Find()
	op.SetBatchSize(200)
	op.SetProjection(bson.M{"aid": 1, "_id": 0})
	result, err := video.Find(context.TODO(), filter, op)

	if err != nil {
		panic(err)
	}
	for result.Next(context.TODO()) {
		var s interface{}
		err := result.Decode(&s)
		if err != nil {
			panic(err)
		}
		rt = append(rt, s.(primitive.D).Map()["aid"].(string))
	}
	return rt
}

func SelectTitileOfVideoSortByKey(key string) []string {
	video := DbCollection.Collection("video")
	op := options.Find()
	op.SetLimit(200)
	op.SetBatchSize(200)
	op.SetProjection(bson.M{"title": 1, "_id": 0})
	op.SetSort(bson.M{key: -1})
	result, err := video.Find(context.TODO(), bson.M{}, op)
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

		entity, ok := s.(primitive.D).Map()["title"]
		if ok {
			rt = append(rt, entity.(string))
		}
	}
	return rt
}

func SelectKeyByCondition(key string, limitCount int64, skip int64, lastValue int32) []int32 {
	video := DbCollection.Collection("video")
	filter := bson.M{
		key: bson.M{
			"$lt": lastValue,
		},
	}
	op := options.Find()
	op.SetLimit(limitCount)
	op.SetSkip(skip)
	op.SetProjection(bson.M{key: 1})
	op.SetSort(bson.M{key: -1})
	result, err := video.Find(context.TODO(), filter, op)

	if err != nil {
		panic(err)
	}

	var rt []int32
	for result.Next(context.TODO()) {
		var s interface{}
		err := result.Decode(&s)
		if err != nil {
			panic(err)
		}
		tmpValue := s.(primitive.D).Map()[key].(int32)
		rt = append(rt, tmpValue)
	}
	return rt
}

func EstimatedDocumentCount() int64 {
	video := DbCollection.Collection("video")
	rt, err := video.EstimatedDocumentCount(context.TODO())
	if err != nil {
		panic(err)
	}
	return rt
}
