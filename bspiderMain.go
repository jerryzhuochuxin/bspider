package main

import (
	"bspider/mongodb"
	"bspider/object/engineBo"
	"bspider/service/schedule"
	"bspider/util"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

var (
	debug      bool
	banner     bool
	mongodbUrl string
	dbName     string
)

func getQueueBoSlice(debug bool) []engineBo.QueueBo {
	var rt []engineBo.QueueBo
	if debug {
		logrus.Info("debug mode")
		logrus.SetLevel(logrus.DebugLevel)

		rt = schedule.DebugSpider()
		schedule.CommonDebug()
	} else {
		logrus.Info("product mode")
		logrus.SetLevel(logrus.InfoLevel)

		rt = schedule.RunSpider()
		schedule.CommonRun()
	}

	if len(rt) == 0 && !debug {
		panic("no task to run")
	}

	return rt
}

func runQueueBo(queueBoSlice []engineBo.QueueBo) {
	logrus.Info(util.GetCurrentFuncName()+" start task count is", len(queueBoSlice))
	for _, queueBo := range queueBoSlice {
		tmpQueueBo := queueBo
		go func() {
			for {
				tmpQueueBo.RunNextWork()
				time.Sleep(time.Second / 4)
			}
		}()
	}
}

func checkHeartBeat() {
	for {
		logrus.Info("bspider is alive")
		time.Sleep(10 * time.Second)
	}
}

func printBanner() {
	if !banner {
		return
	}
	content, err := ioutil.ReadFile("./conf/banner.txt")
	if err != nil {
		logrus.Error("banner.txt not exist")
	}
	fmt.Println(string(content) + "\n\n")
}

func main() {
	flag.BoolVar(&debug, "debug", true, "set debug")
	flag.BoolVar(&banner, "banner", true, "set banner")
	flag.StringVar(&mongodbUrl, "dbUrl", "mongodb://localhost:27017", "set dbUrl")
	flag.StringVar(&dbName, "dbName", "bspider", "set dbName")
	flag.Parse()

	printBanner()
	mongodb.SetDb(mongodbUrl,dbName)
	runQueueBo(getQueueBoSlice(debug))

	checkHeartBeat()
}
