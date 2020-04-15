package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"bspider/model"
	"log"
	"time"
)

func InsertSiteInfoToDb(object model.SiteInfoDo) {
	siteInfo := DbCollection.Collection("siteInfo")
	inJson := bson.M{
		"region_count": object.RegionCount,
		"all_count":    object.AllCount,
		"web_online":   object.WebOnline,
		"play_online":  object.PlayOnline,
		"datetime":     time.Now().Unix(),
	}
	_, err := siteInfo.InsertOne(context.TODO(), inJson)
	if err != nil {
		panic(err)
	}
	log.Printf("insert siteInto to mongodb %v", object)
}
