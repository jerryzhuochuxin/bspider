package online

import (
	"bspider/engine"
	"testing"
)

func TestCatchOnlineFromWorker(t *testing.T) {
	url := "https://www.bilibili.com/video/online.html"
	w := engine.Worker{Url: url}
	CatchOnlineFromWorker(w)
}
