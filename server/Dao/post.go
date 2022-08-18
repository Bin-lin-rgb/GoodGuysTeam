package Dao

import "server/Model"

func (mgr manager) StoreArticle(postInfo Model.Post) error {
	if err := mgr.db.Create(&postInfo).Error; err != nil {
		return err
	}
	return nil
}
