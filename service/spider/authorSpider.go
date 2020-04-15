package spider

import (
	"bspider/mongodb/dao"
	"bspider/mongodb/dao/model"
	"bspider/object/engineBo"
	"bytes"
	"github.com/gocolly/colly"
	"github.com/thedevsaddam/gojsonq"
	"gopkg.in/xmlpath.v1"
	"time"
)

func CatchAuthorFromTank(w engineBo.WorkerBo) {
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})

	xpath := "//*[@id='app']/div[1]/div/div[1]/div[2]/div[3]/ul/li/div[2]/div[2]/div/a/@href"
	urlPre := "https://api.bilibili.com/x/web-interface/card?mid="
	c.OnScraped(func(e *colly.Response) {
		node, err := xmlpath.ParseHTML(bytes.NewBuffer(e.Body))
		if err != nil {
			panic(err)
		}
		path := xmlpath.MustCompile(xpath)
		it := path.Iter(node)
		for it.Next() {
			mid := it.Node().String()[21:]
			w.Queue.AddWork(engineBo.WorkerBo{Url: urlPre + mid, Method: CatchAuthorFromTankSecond})
		}
	})

	err := c.Visit(w.Url)
	if err != nil {
		panic(err)
	}
}

func CatchAuthorFromTankSecond(w engineBo.WorkerBo) {
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})

	c.OnScraped(func(e *colly.Response) {
		jq := gojsonq.New().Reader(bytes.NewBuffer(e.Body))

		data := jq.Find("data").(map[string]interface{})
		card := data["card"].(map[string]interface{})

		fans := int(card["fans"].(float64))
		if fans > 1000 {
			tp := model.AuthorDo{
				CFans:      fans,
				Focus:      true,
				Level:      int(card["level_info"].(map[string]interface{})["current_level"].(float64)),
				Mid:        card["mid"].(string),
				Name:       card["name"].(string),
				Face:       card["face"].(string),
				Official:   card["Official"].(map[string]interface{})["title"].(string),
				Sex:        card["sex"].(string),
				CAttention: int(card["attention"].(float64)),
				CArchive:   int(data["archive_count"].(float64)),
				CArticle:   int(data["article_count"].(float64)),
			}
			tp.Data = map[string]interface{}{
				"fans":      tp.CFans,
				"attention": tp.CAttention,
				"archive":   tp.CArchive,
				"article":   tp.CArticle,
				"datetime":  time.Now().Unix(),
			}
			dao.UpsertAuthorToDb(tp)
		}
	})

	err := c.Visit(w.Url)
	if err != nil {
		panic(err)
	}
}

func CatchAuthorFromFucus(w engineBo.WorkerBo) {
	c := colly.NewCollector()

	c.Limit(&colly.LimitRule{
		DomainRegexp: "*",
		RandomDelay:  1 * time.Second,
	})

	c.OnScraped(func(e *colly.Response) {
		jq := gojsonq.New().Reader(bytes.NewBuffer(e.Body))

		data := jq.Find("data").(map[string]interface{})
		card := data["card"].(map[string]interface{})
		fansTmp, err := card["fans"]
		if err == true {
			model := model.AuthorDo{
				Mid:          card["mid"].(string),
				Name:         card["name"].(string),
				Face:         card["face"].(string),
				Official:     card["Official"].(map[string]interface{})["title"].(string),
				Sex:          card["sex"].(string),
				Data:         nil,
				Level:        int(card["level_info"].(map[string]interface{})["current_level"].(float64)),
				Focus:        false,
				Pts:          "",
				CFans:        int(fansTmp.(float64)),
				CArchive:     int(data["archive_count"].(float64)),
				CArticle:     int(data["article_count"].(float64)),
				CArchiveView: 0,
				CArticleView: 0,
				CLike:        0,
				CDatetime:    0,
			}

			attention, ok := card["attention"]
			if ok {
				model.CAttention = int(attention.(float64))
			}

			model.Data = map[string]interface{}{
				"fans":      model.CFans,
				"attention": model.CAttention,
				"archive":   model.CArchive,
				"article":   model.CArticle,
				"datetime":  time.Now().Unix(),
			}
			urlPre := "https://api.bilibili.com/x/space/upstat?mid="
			nw := engineBo.WorkerBo{Url: urlPre + model.Mid, Data: model, Queue: w.Queue, Method: CatchAuthorFromFucusSecond}
			w.Queue.AddWork(nw)
		}
	})

	c.Visit(w.Url)
}

func CatchAuthorFromFucusSecond(w engineBo.WorkerBo) {
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{
		DomainRegexp: "*",
		RandomDelay:  1 * time.Second,
	})
	c.OnScraped(func(e *colly.Response) {
		jq := gojsonq.New().Reader(bytes.NewBuffer(e.Body))
		data := jq.Find("data").(map[string]interface{})

		model := w.Data.(model.AuthorDo)
		model.CArchiveView = int(data["archive"].(map[string]interface{})["view"].(float64))
		model.CArticleView = int(data["article"].(map[string]interface{})["view"].(float64))
		model.CLike = int(data["likes"].(float64))
		model.CRate = 0
		re := dao.AggregateForAuthorByMid(model)
		for _, v := range re {
			deltaSeconds := float32(time.Now().Unix() - v["datetime"].(int64))
			deltaFans := float32(int32(model.CFans) - v["fans"].(int32))
			model.CRate = int(deltaFans / deltaSeconds * 86400)
		}
		dao.UpsertAuthorToDb(model)
	})
	c.Visit(w.Url)
}
