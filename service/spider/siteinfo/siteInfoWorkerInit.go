package siteinfo

import (
	"bspider/engine"
	"log"
)

func InitWorkerForSiteInfo(q *engine.Queue) {
	url := "https://api.bilibili.com/x/web-interface/online"
	w := engine.Worker{Url: url, Method: CatchFromWorker, Queue: q}
	q.AddWork(w)
	log.Printf("add worker %v", w)
}
