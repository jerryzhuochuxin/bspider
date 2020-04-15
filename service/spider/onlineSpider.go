package spider

import (
	"bspider/mongodb/dao"
	"bspider/mongodb/dao/model"
	"bspider/object/engineBo"
	"bspider/util"
	"bytes"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"gopkg.in/xmlpath.v1"
	"time"
)

func CatchOnlineFromWorker(w engineBo.WorkerBo) {
	c := colly.NewCollector()

	const xPathPre = `//*[@id="app"]/div/div[2]/div`
	const titleXpath = xPathPre + `/a/p`
	const watchXpath = xPathPre + `/p/b`
	const authorXpath = xPathPre + `/div/a`
	const hrefXpath = xPathPre + `/a/@href`

	c.OnScraped(func(e *colly.Response) {
		node, err := xmlpath.ParseHTML(bytes.NewBuffer(e.Body))
		if err != nil {
			panic(err)
		}
		titles := util.ParseNodeToStringSlice(titleXpath, node)
		watches := util.ParseNodeToIntSlice(watchXpath, node)
		authors := util.ParseNodeToStringSlice(authorXpath, node)
		hrefs := util.ParseNodeToStringSlice(hrefXpath, node)
		ln := len(titles)

		for i := 0; i < ln; i++ {
			ml := model.VideoOnlineDo{
				Title:  titles[i],
				Author: authors[i],
				Aid:    hrefs[i][25:],
			}
			data := make(map[string]interface{})
			data["datetime"] = time.Now().Unix()
			data["number"] = watches[i]
			ml.Data = data
			wk := engineBo.WorkerBo{Url: "https:" + hrefs[i], Method: CatchOnlineFromWorkerSecond, Queue: w.Queue, Data: ml}
			w.Queue.AddWork(wk)
		}
	})

	c.Visit(w.Url)
}

func CatchOnlineFromWorkerSecond(w engineBo.WorkerBo) {
	c := colly.NewCollector()
	const channelXpath = `//*[@id="viewbox_report"]/div[1]/span[1]/a[1]`
	const subChannelXpath = `//*[@id="viewbox_report"]/div[1]/span[1]/a[2]`

	c.OnScraped(func(e *colly.Response) {
		root, err := htmlquery.Parse(bytes.NewBuffer(e.Body))
		ml := w.Data.(model.VideoOnlineDo)
		if err != nil {
			panic(err)
		}

		channel := util.ParseNodeToStringUseHtmlquery(root, channelXpath)
		if channel != "" {
			ml.Channel = channel
		} else {
			ml.Channel = "番剧"
		}

		subChannel := util.ParseNodeToStringUseHtmlquery(root, subChannelXpath)
		if subChannel != "" {
			ml.SubChannel = subChannel
		} else {
			ml.SubChannel = "番剧"
		}
		dao.UpsertVideoOnlineToDb(ml)
	})

	c.Visit(w.Url)
}
