package siteinfo

import (
	"bspider/engine"
	"testing"
)

func TestCatchFromWorker(t *testing.T) {
	url := "https://api.bilibili.com/x/web-interface/online"
	w := engine.Worker{Url: url}
	CatchFromWorker(w)
}
