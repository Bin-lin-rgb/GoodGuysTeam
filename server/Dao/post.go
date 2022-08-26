package Dao

import (
	"server/Model"
	"time"
)

func (mgr manager) StoreArticle(postInfo Model.Post) error {
	if err := mgr.db.Create(&postInfo).Error; err != nil {
		return err
	}
	return nil
}

func (mgr manager) GetList(latestTime int64) ([]Model.Post, error) {
	var List []Model.Post
	timeLayout := "2006-01-02 15:04:05"
	var count int64
	mgr.db.Model(&Model.Post{}).Count(&count)
	result := mgr.db.Model(&Model.Post{}).Where("created_at <= ?", time.Unix(latestTime, 0).Format(timeLayout)).
		Order("created_at DESC").Limit(10).Count(&count).Find(&List)

	if 1 < count && count <= 9 {
		result := mgr.db.Model(&Model.Post{}).Where("created_at <= ?", time.Unix(latestTime, 0).Format(timeLayout)).
			Order("created_at DESC").Count(&count).Find(&List)
		return List, result.Error
	}

	if count == 1 {
		//	是小于等于最早时间只能查到1个文章记录, 证明已经刷完了
		result := mgr.db.Model(&Model.Post{}).Order("created_at DESC").
			Preload("Author").Limit(10).Count(&count).Find(&List)
		return List, result.Error
	}
	return List, result.Error
}

func (mgr manager) GetPostById(postId int64) (Model.Post, error) {
	var postDetail Model.Post
	if err := mgr.db.Where("id = ?", postId).Find(&postDetail).Error; err != nil {
		return postDetail, err
	}
	return postDetail, nil
}

func (mgr manager) GetUsernameById(AuthorId uint64) (author string, err error) {
	var result Model.User
	if err := mgr.db.Table("user").Select("username").Where("user_id = ?", AuthorId).Find(&result).Error; err != nil {
		return "", err
	}
	return result.Username, nil
}

func (mgr manager) GetCommunityNameById(CommunityId uint64) (CommunityName string, err error) {
	var result Model.Community
	if err := mgr.db.Table("community").Select("community_name").Where("community_id = ?", CommunityId).Find(&result).Error; err != nil {
		return "", err
	}
	return result.CommunityName, nil
}

func (mgr manager) GetPostListByUserId(userId uint64) ([]Model.Post, error) {
	var postLists []Model.Post
	if err := mgr.db.Where("author_id = ?", userId).Find(&postLists).Error; err != nil {
		return postLists, err
	}
	return postLists, nil
}

func (mgr manager) DeletePostById(postId int64) error {
	var post Model.Post
	if err := mgr.db.Where("id = ?", postId).Delete(&post).Error; err != nil {
		return err
	}
	return nil
}
