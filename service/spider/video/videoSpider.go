package video

import (
	"bytes"
	"github.com/gocolly/colly"
	"github.com/thedevsaddam/gojsonq"
	"bspider/constant"
	"bspider/engine"
	"bspider/model"
	"bspider/mongodb"
	"strconv"
	"time"
)

func CatchFromFucus(w engine.Worker) {
	c := colly.NewCollector()

	c.OnScraped(func(e *colly.Response) {
		jq := gojsonq.New().Reader(bytes.NewBuffer(e.Body))
		data := jq.Find("data").(map[string]interface{})
		vListOrigion, ok := data["vlist"]
		if !ok {
			return
		}
		tList := data["tlist"].(map[string]interface{})
		vList := vListOrigion.([]interface{})
		var aidList []string
		var mid string

		for _, v := range vList {
			aid := strconv.Itoa(int(v.(map[string]interface{})["aid"].(float64)))
			aidList = append(aidList, aid)
			mid = strconv.Itoa(int(v.(map[string]interface{})["mid"].(float64)))
		}
		model := model.VideoWithMidAidChannelsDo{
			Mid:      mid,
			Aid:      aidList,
			Channels: tList,
		}
		urlList := mongodb.UpsertVideoWithMidAidChannelsToDb(model)
		mongodb.UpsertChannelToDb(model)
		for _, url := range urlList {
			w.Queue.AddWork(engine.Worker{Url: url, Method: CatchFromFucusSecond, Queue: w.Queue})
		}
	})

	c.Visit(w.Url)
}

func CatchFromFucusSecond(w engine.Worker) {
	c := colly.NewCollector()

	c.OnScraped(func(e *colly.Response) {
		jq := gojsonq.New().Reader(bytes.NewBuffer(e.Body))
		dataMap := jq.Find("data").(map[string]interface{})
		var modelList []model.VideoDo
		for _, idata := range dataMap {
			data := idata.(map[string]interface{})
			stat := data["stat"].(map[string]interface{})
			owner := data["owner"].(map[string]interface{})

			model := model.VideoDo{
				Aid:             strconv.Itoa(int(stat["aid"].(float64))),
				Mid:             strconv.Itoa(int(owner["mid"].(float64))),
				Pic:             data["pic"].(string),
				Author:          owner["name"].(string),
				Title:           data["title"].(string),
				Datetime:        int(data["pubdate"].(float64)),
				CurrentView:     int(stat["view"].(float64)),
				CurrentFavorite: int(stat["favorite"].(float64)),
				CurrentDanmaku:  int(stat["danmaku"].(float64)),
				CurrentCoin:     int(stat["coin"].(float64)),
				CurrentShare:    int(stat["share"].(float64)),
				CurrentLike:     int(stat["like"].(float64)),
				CurrentDatetime: time.Now().Unix(),
			}

			model.Data = map[string]interface{}{
				"view":     model.CurrentView,
				"favorite": model.CurrentFavorite,
				"danmaku":  model.CurrentDanmaku,
				"coin":     model.CurrentCoin,
				"share":    model.CurrentShare,
				"like":     model.CurrentLike,
				"datetime": model.CurrentDatetime,
			}

			tid := int(data["tid"].(float64))
			subChannel, ok := data["tname"]
			if ok {
				model.SubChannel = subChannel.(string)
				model.Channel = constant.SubChannel2Channel[model.SubChannel]
			} else {
				if tid == 51 {
					model.Channel = "番剧"
				} else if tid == 170 {
					model.Channel = "国创"
				} else if tid == 159 {
					model.Channel = "娱乐"
				}
			}
			modelList = append(modelList, model)
		}
		mongodb.UpsertVideoToDb(modelList)
	})

	c.Visit(w.Url)
}
