package workerBoOp

import (
	"bspider/mongodb/dao"
	"bspider/object/engineBo"
	"bspider/service/spider"
)

func AddAuthorFromTankWorkerBo(q *engineBo.QueueBo) {
	var rt = []string{"https://www.bilibili.com/ranking",
		"https://www.bilibili.com/ranking/all/1/0/3",
		"https://www.bilibili.com/ranking/all/168/0/3",
		"https://www.bilibili.com/ranking/all/3/0/3",
		"https://www.bilibili.com/ranking/all/129/0/3",
		"https://www.bilibili.com/ranking/all/188/0/3",
		"https://www.bilibili.com/ranking/all/4/0/3",
		"https://www.bilibili.com/ranking/all/36/0/3",
		"https://www.bilibili.com/ranking/all/160/0/3",
		"https://www.bilibili.com/ranking/all/119/0/3",
		"https://www.bilibili.com/ranking/all/155/0/3",
		"https://www.bilibili.com/ranking/all/5/0/3",
		"https://www.bilibili.com/ranking/all/181/0/3"}
	for _, url := range rt {
		wo := engineBo.WorkerBo{Url: url, Method: spider.CatchAuthorFromTank, Queue: q}
		q.AddWork(wo)
	}
}
func AddAuthorFromFucusWorkerBo(q *engineBo.QueueBo) {
	urlPre := "https://api.bilibili.com/x/web-interface/card?mid="
	for _, mid := range dao.SelectMidForAuthorByFucus() {
		wo := engineBo.WorkerBo{Url: urlPre + mid, Method: spider.CatchAuthorFromFucus, Queue: q}
		q.AddWork(wo)
	}
}
