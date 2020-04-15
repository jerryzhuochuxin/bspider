package workerBoOp

import (
	"bspider/object/engineBo"
	"bspider/service/spider"
)

func AddVideoOnlineWorkerBoOp(q *engineBo.QueueBo) {
	url := "https://www.bilibili.com/video/online.html"
	w := engineBo.WorkerBo{Url: url, Method: spider.CatchOnlineFromWorker, Queue: q}
	q.AddWork(w)
}
