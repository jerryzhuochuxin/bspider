package workerBoOp

import (
	"bspider/object/engineBo"
	"bspider/service/spider"
)

func AddSiteInfoWorkerBo(q *engineBo.QueueBo) {
	url := "https://api.bilibili.com/x/web-interface/videoOnline"
	w := engineBo.WorkerBo{Url: url, Method: spider.CatchFromWorker, Queue: q}
	q.AddWork(w)
}
