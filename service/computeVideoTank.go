package service

import (
	"bspider/mongodb"
	"time"
)

func ComputeVideoTankTable() {
	keys := []string{"cView", "cLike", "cDanmaku", "cFavorite", "cCoin", "cShare"}
	var mp = make(map[string]interface{})
	mp["name"] = "video_rank"

	docCount := mongodb.EstimatedDocumentCount()
	skipCount := docCount / 100
	var lastValue int32 = 999999999
	for _, key := range keys {
		titles := mongodb.SelectTitileOfVideoSortByKey(key)
		var tmpMp = make(map[string]interface{})

		for top := 1; top <= len(titles); top++ {
			tmpMp[titles[top-1]] = top
		}
		var tmpRate []int32
		for i := 1; i <= 60; i++ {
			keySelect := mongodb.SelectKeyByCondition(key, 1, skipCount, lastValue)
			if len(keySelect) == 0 {
				continue
			}
			lastValue = keySelect[0]
			tmpRate = append(tmpRate, keySelect[0])
		}
		tmpMp["rate"] = tmpRate

		mp[key] = tmpMp
	}
	mp["update_time"] = time.Now().Unix()
	mongodb.UpsertRankTableToDb("video_rank", mp)
}
