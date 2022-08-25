package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"server/Dao"
	"server/Model"
	"server/pkg/snowflake"
	"strconv"
	"time"
)

type PostListResponse struct {
	Code     MyCode      `json:"code"`
	Data     interface{} `json:"data,omitempty"`
	NextTime int64       `json:"next_time,omitempty"`
}

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

func PostList(c *gin.Context) {
	latestTimeString := c.Query("latest_time")

	latestTime, err := strconv.ParseInt(latestTimeString, 10, 64)
	if err != nil {
		zap.L().Error("Get post list detail with invalid param (latestTime) .", zap.Error(err))
		//ResponseError(c, CodeInvalidParams)
		latestTime = time.Now().Unix()
	}

	list, err := Dao.Mgr.GetList(latestTime)
	if err != nil {
		zap.L().Error("Dao.Mgr.GetList(latestTime) failed.", zap.Error(err))
		//ResponseError(c, CodeServerBusy)
	}

	ListLength := len(list)

	data := make([]*Model.ApiPostList, 0, ListLength)
	timeLayoutDate := "01-02"
	timeLayoutymdm := "2006-01-02 15:04"
	for _, post := range list {
		// 根据作者id查询作者信息
		authorName, err := Dao.Mgr.GetUsernameById(post.AuthorId)
		if err != nil {
			zap.L().Error("Dao.Mgr.GetAuthorById() failed",
				zap.Uint64("postID", post.AuthorId),
				zap.Error(err))
			continue
		}
		// 根据社区id查询社区详细信息
		communityName, err := Dao.Mgr.GetCommunityNameById(post.CommunityId)
		if err != nil {
			zap.L().Error("Dao.Mgr.GetCommunityNameById() failed",
				zap.Uint64("community_id", post.CommunityId),
				zap.Error(err))
			continue
		}
		// 接口数据拼接

		postData := &Model.ApiPostList{
			PostId:        post.Id,
			Title:         post.Title,
			Content:       post.Content,
			CommunityName: communityName,
			AuthorName:    authorName,
			CreatedAt:     post.CreatedAt.Format(timeLayoutDate),
			CreatedAt2:    post.CreatedAt.Format(timeLayoutymdm),
		}
		data = append(data, postData)
	}
	// 本次结构体数组最后一个的下标，
	Thelastone := ListLength - 1

	c.JSON(http.StatusOK, PostListResponse{
		Code:     CodeSuccess,
		Data:     data,
		NextTime: list[Thelastone].CreatedAt.Unix(),
	})
}

func PostDetail(c *gin.Context) {
	// 1、获取参数(从URL中获取帖子的id)
	postIdStr := c.Param("id")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
	}

	// 2、根据id取出id帖子数据(查数据库)
	post, err := Dao.Mgr.GetPostById(postId)
	if err != nil {
		zap.L().Error("Dao.Mgr.GetPostById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	timeLayoutymdm := "2006-01-02 15:04"
	authorName, err := Dao.Mgr.GetUsernameById(post.AuthorId)
	if err != nil {
		zap.L().Error("Dao.Mgr.GetAuthorById() failed",
			zap.Uint64("postID", post.AuthorId),
			zap.Error(err))
	}
	// 根据社区id查询社区详细信息
	communityName, err := Dao.Mgr.GetCommunityNameById(post.CommunityId)
	if err != nil {
		zap.L().Error("Dao.Mgr.GetCommunityNameById() failed",
			zap.Uint64("community_id", post.CommunityId),
			zap.Error(err))
	}
	// 接口数据拼接

	postData := &Model.ApiPostList{
		PostId:        post.Id,
		Title:         post.Title,
		Content:       post.Content,
		CommunityName: communityName,
		AuthorName:    authorName,
		CreatedAt2:    post.CreatedAt.Format(timeLayoutymdm),
	}

	c.JSON(http.StatusOK, PostListResponse{
		Code: CodeSuccess,
		Data: postData,
	})
}

func UserBlogList(c *gin.Context) {
	// 1、获取当前请求的UserID(从c取到当前发请求的用户ID)
	userId, err := GetCurrentUserID(c)
	fmt.Println(userId)
	if err != nil {
		zap.L().Error("GetCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
	}

	// 2、根据id取出id帖子数据(查数据库)
	postList, err := Dao.Mgr.GetPostListByUserId(userId)
	if err != nil {
		zap.L().Error("Dao.Mgr.GetPostListByUserId(userId)", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	data := make([]*Model.ApiPostList, 0, len(postList))
	timeLayoutymdm := "2006-01-02 15:04"

	for _, post := range postList {
		postData := &Model.ApiPostList{
			Title:      post.Title,
			Content:    post.Content,
			CreatedAt2: post.CreatedAt.Format(timeLayoutymdm),
		}
		data = append(data, postData)
	}

	c.JSON(http.StatusOK, PostListResponse{
		Code: CodeSuccess,
		Data: data,
	})
}
