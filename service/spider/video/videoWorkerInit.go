package video

import (
	"fmt"
	"bspider/engine"
	"bspider/mongodb"
	"log"
)

func InitWorkerForVideoFromFucus(q *engine.Queue) {
	urlPre := "https://space.bilibili.com/ajax/member/getSubmitVideos?mid=%s&pagesize=10&page=1&order=pubdate"
	for _, mid := range mongodb.SelectMidForAuthorByFucus() {
		wo := engine.Worker{Url: fmt.Sprintf(urlPre, mid), Method: CatchFromFucus, Queue: q}
		q.AddWork(wo)
		log.Printf("add a worker is %v", wo)
	}
}
func InitWorkerForVideoByFocus(q *engine.Queue) {
	urlPre := "https://api.bilibili.com/x/article/archives?ids="
	for _, aid := range mongodb.SelectAidListForVideoUpdate(true) {
		wo := engine.Worker{Url: urlPre + aid, Method: CatchFromFucus, Queue: q}
		q.AddWork(wo)
		log.Printf("add a worker is %v", wo)
	}
}

func InitWorkerForVideoByUnFocus(q *engine.Queue) {
	urlPre := "https://api.bilibili.com/x/article/archives?ids="
	for _, aid := range mongodb.SelectAidListForVideoUpdate(false) {
		wo := engine.Worker{Url: urlPre + aid, Method: CatchFromFucus, Queue: q}
		q.AddWork(wo)
		log.Printf("add a worker is %v", wo)
	}
}
