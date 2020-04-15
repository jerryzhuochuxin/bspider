package schedule

import (
	"github.com/jasonlvhit/gocron"
	"bspider/service"
	"bspider/util"
	"log"
	"time"
)

func CommonRun() {
	aopFunc(gocron.Every(1).Monday().At("03:20").Do, service.ComputeAuthorTankTable)
	aopFunc(gocron.Every(1).Wednesday().At("03:20").Do, service.ComputeVideoTankTable)
}
func CommonDebug() {
	te := time.Now().Add(1 * time.Second).Format("15:04:05")
	funcList := []interface{}{
		service.ComputeVideoTankTable,
		service.ComputeAuthorTankTable,
	}
	for _, v := range funcList {
		aopFunc(gocron.Every(1).Day().At(te).Do, v)
	}
}

func aopFunc(fun util.ScheduleFunc, jobFun interface{}) {
	err := fun(jobFun)
	funcName := util.GetFuncName(jobFun)
	if err != nil {
		log.Printf(funcName+" fail ", err)
	} else {
		log.Printf(funcName + " success ")
	}
}
