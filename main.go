package main

import (
	"bspider/engine"
	"bspider/schedule"
	"bspider/util"
	"flag"
	"log"
	"time"
)

var debug = true

const clock = 8

func main() {
	flag.BoolVar(&debug, "debug", true, "set debug")
	flag.Parse()

	var qList []engine.Queue
	if debug {
		qList = schedule.SpiderDebug()
		schedule.CommonDebug()
	} else {
		qList = schedule.SpiderRun()
		schedule.CommonRun()
	}

	if len(qList) == 0 && !debug {
		panic("no task to run")
	}

	log.Printf(util.GetCurrentFuncName()+" start task count is %d", len(qList))
	for _, q := range qList {
		ky := q
		go func() {
			for {
				w := ky.GetWork()
				time.Sleep(time.Second / clock)
				w.Method(w)
				log.Printf("remove a worker %v", w)
			}
		}()
	}

	for {
		log.Printf(util.GetCurrentFuncName() + " lauching")
		time.Sleep(10 * time.Second)
	}
}
