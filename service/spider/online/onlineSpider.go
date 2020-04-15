package online

import (
	"bspider/engine"
	"bytes"
	"github.com/gocolly/colly"
	"gopkg.in/xmlpath.v1"
	"log"
)

func CatchOnlineFromWorker(w engine.Worker) {
	c := colly.NewCollector()

	//xPathPre := `//*[@id="app"]/div/div[2]/div`
	titleX := `//*[@id="app"]/div/div[2]/div/a/p/text`
	//watchX := xPathPre + `./p/b/text()`
	//authorX := xPathPre + `./div[1]/a/text`
	//hrefX  := xPathPre + `./a/@href`
	c.OnHTML()
	c.OnScraped(func(e *colly.Response) {
		node, err := xmlpath.ParseHTML(bytes.NewBuffer(e.Body))
		if err != nil {
			panic(err)
		}
		//log.Printf("#%v", node)
		xPath := xmlpath.MustCompile(titleX)
		it := xPath.Iter(node)
		for it.Next() {
			log.Printf("%s\n", it.Node().String())
		}
	})

	c.Visit(w.Url)
}
