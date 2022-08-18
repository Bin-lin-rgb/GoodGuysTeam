package Dao

import (
	"server/Model"
)

func (mgr manager) Register(userinfo Model.User) error {
	result := mgr.db.Create(&userinfo)
	return result.Error
}

func (mgr manager) GetUser(userId uint64) (Model.User, error) {
	var user Model.User
	if err := mgr.db.Select("user_id", "username", "email", "gender", "created_at").
		Where("user_id=?", userId).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (mgr manager) IsExist(username string) (Model.User, error) {
	var userinfo Model.User
	result := mgr.db.Model(&Model.User{}).Where("username=?", username).Find(&userinfo)

	return userinfo, result.Error
}
