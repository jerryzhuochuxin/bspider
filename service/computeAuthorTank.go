package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"bspider/mongodb"
	"bspider/util"
	"time"
)

func ComputeAuthorTankTable() {
	keys := []string{"cFans", "cArchive_view", "cArticle_view"}
	keysHelp := map[string]string{keys[0]: "fans", keys[1]: "archiveView", keys[2]: "articleView"}
	count := mongodb.CountAuthorCountByKey(keys[0])

	for _, eachKey := range keys {

		authors := mongodb.SelectKey(eachKey)
		eachRank := keysHelp[eachKey] + "Rank"
		eachDRank := "d" + keysHelp[eachKey] + "Rank"
		eachPRank := "p" + keysHelp[eachKey] + "Rank"

		for i, eachAuthor := range authors {
			rank := make(map[string]interface{})
			rankInterface, tmpHasKey := eachAuthor["rank"]
			//eachDRank设置
			if tmpHasKey {
				rank = rankInterface.(primitive.D).Map()
				_, tmpHasKey2 := rank[eachRank]
				if tmpHasKey2 {
					rank[eachDRank] = int(rank[eachRank].(int32)) + i
				} else {
					rank[eachDRank] = 0
				}
			} else {
				rank[eachDRank] = 0
			}
			rank[eachRank] = i
			rank[eachPRank] = util.Round2(float64(i) / float64(count) * 100)

			keyCount := eachAuthor[eachKey]

			if keyCount.(int32) == 0 {
				rank[eachRank] = -1
				rank[eachPRank] = -1
				rank[eachDRank] = 0
			}

			if eachKey == "cArchive_view" {
				rank["updateTime"] = time.Now().Unix()
			}
			mongodb.UpsertRankToDb(rank, eachAuthor["mid"].(string))
		}
	}
}
