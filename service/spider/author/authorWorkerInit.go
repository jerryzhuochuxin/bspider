package author

import (
	"bspider/engine"
	"bspider/mongodb"
	"log"
)

func InitWorkerForAuthorFromTank(q *engine.Queue) {
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
		wo := engine.Worker{Url: url, Method: CatchFromTank, Queue: q}
		q.AddWork(wo)
		log.Printf("add a worker is %v", wo)
	}
}
func InitWorkerForAuthorFromFucus(q *engine.Queue) {
	urlPre := "https://api.bilibili.com/x/web-interface/card?mid="
	for _, mid := range mongodb.SelectMidForAuthorByFucus() {
		wo := engine.Worker{Url: urlPre + mid, Method: CatchFromFucus, Queue: q}
		q.AddWork(wo)
		log.Printf("add a worker is %v", wo)
	}
}
