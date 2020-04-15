package dao

import (
	"bspider/mongodb"
	"bspider/mongodb/dao/model"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func InsertSiteInfoToDb(object model.SiteInfoDo) {
	siteInfo := mongodb.DbCollection.Collection("siteInfo")
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
	logrus.Info("insert siteInto to mongodb ", object)
}
