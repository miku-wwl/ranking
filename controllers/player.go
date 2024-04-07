package controllers

import (
	"ranking/cache"
	"ranking/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PalyerController struct{}

func (p PalyerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	rs, err := models.GetPlayers(aid, "id asc")
	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}
	ReturnSuccess(c, 0, "success", rs, 1)
}

func (p PalyerController) GetRanking(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	redisKey := "ranking:" + aidStr
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()

	if err == nil && len(rs) > 0 {
		var players []models.Player
		for _, value := range rs {
			id, _ := strconv.Atoi(value)
			rsInfo, _ := models.GetPlayerInfo(id)
			if rsInfo.Id > 0 {
				players = append(players, rsInfo)
			}
		}
		ReturnSuccess(c, 0, "successRedis", players, 1)
		return
	}

	rsDb, errDB := models.GetPlayers(aid, "score desc")
	if errDB == nil {
		for _, value := range rsDb {
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.Id, value.Score))
		}
		//设置过期时间
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)
		ReturnSuccess(c, 0, "success", rsDb, 1)
		return
	}
	ReturnError(c, 4004, "没有相关信息")
	return
}
