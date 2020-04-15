package schedule

import (
	"github.com/jasonlvhit/gocron"
	"bspider/engine"
	"bspider/service/spider/author"
	"bspider/service/spider/siteinfo"
	"bspider/service/spider/video"
	"bspider/util"
	"log"
	"time"
)

func SpiderRun() []engine.Queue {
	var rt []engine.Queue
	aopSpiderFunc(gocron.Every(1).Day().At("01:00").Do, author.InitWorkerForAuthorFromFucus, 1000, &rt)
	aopSpiderFunc(gocron.Every(1).Day().At("07:00").Do, video.InitWorkerForVideoByFocus, 1000, &rt)
	aopSpiderFunc(gocron.Every(1).Day().At("14:00").Do, author.InitWorkerForAuthorFromTank, 1000, &rt)
	aopSpiderFunc(gocron.Every(1).Day().At("22:00").Do, video.InitWorkerForVideoFromFucus, 1000, &rt)

	aopSpiderFunc(gocron.Every(1).Hour().Do, siteinfo.InitWorkerForSiteInfo, 1000, &rt)
	aopSpiderFunc(gocron.Every(1).Week().Do, video.InitWorkerForVideoByFocus, 1000, &rt)
	go func() {
		for {
			<-gocron.Start()
		}
	}()
	return rt
}

func SpiderDebug() []engine.Queue {
	te := time.Now().Add(1 * time.Second).Format("15:04:05")

	funcList := []interface{}{
		author.InitWorkerForAuthorFromFucus,
		video.InitWorkerForVideoByFocus,
		video.InitWorkerForVideoByUnFocus,
		author.InitWorkerForAuthorFromTank,
		video.InitWorkerForVideoFromFucus,
		siteinfo.InitWorkerForSiteInfo,
	}

	var rt []engine.Queue
	for _, v := range funcList {
		aopSpiderFunc(gocron.Every(1).Day().At(te).Do, v, 3000, &rt)
	}
	go func() {
		for {
			<-gocron.Start()
		}
	}()
	return rt
}

func aopSpiderFunc(fun util.ScheduleFunc, jobFun interface{}, batch int, q *[]engine.Queue) {
	qe := engine.Queue(make(chan engine.Worker, batch))
	err := fun(jobFun, &qe)
	funcName := util.GetFuncName(jobFun)
	if err != nil {
		log.Printf(funcName+" fail ", err)
	} else {
		log.Printf(funcName + " success ")
		*q = append(*q, qe)
	}
}
