package siteinfo

import (
	"bytes"
	"github.com/gocolly/colly"
	"github.com/thedevsaddam/gojsonq"
	"bspider/engine"
	"bspider/model"
	"bspider/mongodb"
)

func CatchFromWorker(w engine.Worker) {
	c := colly.NewCollector()

	c.OnScraped(func(e *colly.Response) {
		jq := gojsonq.New().Reader(bytes.NewBuffer(e.Body))
		data := jq.Find("data").(map[string]interface{})

		object := model.SiteInfoDo{
			RegionCount: data["region_count"].(map[string]interface{}),
			AllCount:    int(data["all_count"].(float64)),
			WebOnline:   int(data["web_online"].(float64)),
			PlayOnline:  int(data["play_online"].(float64)),
		}
		mongodb.InsertSiteInfoToDb(object)
	})

	c.Visit(w.Url)
}
