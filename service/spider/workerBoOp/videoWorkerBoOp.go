package workerBoOp

import (
	"bspider/mongodb/dao"
	"bspider/object/engineBo"
	"bspider/service/spider"
	"fmt"
)

func AddVideoFromFucusWorkerBo(q *engineBo.QueueBo) {
	urlPre := "https://space.bilibili.com/ajax/member/getSubmitVideos?mid=%s&pagesize=10&page=1&order=pubdate"
	for _, mid := range dao.SelectMidForAuthorByFucus() {
		wo := engineBo.WorkerBo{Url: fmt.Sprintf(urlPre, mid), Method: spider.CatchVideoFromFucus, Queue: q}
		q.AddWork(wo)
	}
}
func AddVideoByFocusWorkerBo(q *engineBo.QueueBo) {
	urlPre := "https://api.bilibili.com/x/article/archives?ids="
	for _, aid := range dao.SelectAidListForVideoUpdate(true) {
		wo := engineBo.WorkerBo{Url: urlPre + aid, Method: spider.CatchVideoFromFucus, Queue: q}
		q.AddWork(wo)
	}
}
func AddVideoByUnFocusWorkerBo(q *engineBo.QueueBo) {
	urlPre := "https://api.bilibili.com/x/article/archives?ids="
	for _, aid := range dao.SelectAidListForVideoUpdate(false) {
		wo := engineBo.WorkerBo{Url: urlPre + aid, Method: spider.CatchVideoFromFucus, Queue: q}
		q.AddWork(wo)
	}
}

