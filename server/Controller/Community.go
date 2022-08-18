package Controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/Dao"
)

func GetCommunity(c *gin.Context) {
	list, err := Dao.Mgr.GetCommunityList()
	if err != nil {
		zap.L().Error("GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, list)
}
