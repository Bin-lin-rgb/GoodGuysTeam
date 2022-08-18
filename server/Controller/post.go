package Controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/Dao"
	"server/Model"
	"server/pkg/snowflake"
)

func PostArticle(c *gin.Context) {
	// 1、获取参数及校验参数
	var post Model.Post
	if err := c.ShouldBindJSON(&post); err != nil { // validator --> binding tag
		zap.L().Debug("c.ShouldBindJSON(post) err", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}

	// 获取作者ID，当前请求的UserID(从c取到当前发请求的用户ID)
	userID, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	post.AuthorId = userID
	// 1、 生成post_id(生成帖子ID)
	postID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		return
	}
	post.PostId = postID
	// 2、存入文章
	if err := Dao.Mgr.StoreArticle(post); err != nil {
		zap.L().Error("Dao.Mgr.StoreArticle failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3、返回响应
	ResponseSuccess(c, nil)
}
