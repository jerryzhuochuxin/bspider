package schedule

import (
	"bspider/object/engineBo"
	"bspider/service/spider/workerBoOp"
	"bspider/util"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"time"
)

func RunSpider() []engineBo.QueueBo {
	var rt []engineBo.QueueBo
	aopSpiderFunc(gocron.Every(1).Day().At("01:00").Do, workerBoOp.AddAuthorFromFucusWorkerBo, 1000, &rt)
	aopSpiderFunc(gocron.Every(1).Day().At("07:00").Do, workerBoOp.AddVideoByFocusWorkerBo, 1000, &rt)
	aopSpiderFunc(gocron.Every(1).Day().At("14:00").Do, workerBoOp.AddAuthorFromTankWorkerBo, 1000, &rt)
	aopSpiderFunc(gocron.Every(1).Day().At("22:00").Do, workerBoOp.AddVideoFromFucusWorkerBo, 1000, &rt)

	aopSpiderFunc(gocron.Every(1).Minute().Do, workerBoOp.AddVideoOnlineWorkerBoOp, 1000, &rt)
	aopSpiderFunc(gocron.Every(1).Hour().Do, workerBoOp.AddSiteInfoWorkerBo, 1000, &rt)
	aopSpiderFunc(gocron.Every(1).Week().Do, workerBoOp.AddVideoByUnFocusWorkerBo, 1000, &rt)
	go func() {
		for {
			<-gocron.Start()
		}
	}()
	return rt
}

func DebugSpider() []engineBo.QueueBo {
	te := time.Now().Add(1 * time.Second).Format("15:04:05")

	funcList := []interface{}{
		workerBoOp.AddAuthorFromFucusWorkerBo,
		workerBoOp.AddVideoByFocusWorkerBo,
		workerBoOp.AddVideoByUnFocusWorkerBo,
		workerBoOp.AddAuthorFromTankWorkerBo,
		workerBoOp.AddVideoFromFucusWorkerBo,
		workerBoOp.AddSiteInfoWorkerBo,
		workerBoOp.AddVideoOnlineWorkerBoOp,
	}

	var rt []engineBo.QueueBo
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

func aopSpiderFunc(fun util.ScheduleFunc, jobFun interface{}, batch int, q *[]engineBo.QueueBo) {
	qe := engineBo.QueueBo(make(chan engineBo.WorkerBo, batch))
	err := fun(jobFun, &qe)
	funcName := util.GetFuncName(jobFun)
	if err != nil {
		logrus.Error(funcName+" fail ", err)
	} else {
		logrus.Info(funcName + " success ")
		*q = append(*q, qe)
	}
}
