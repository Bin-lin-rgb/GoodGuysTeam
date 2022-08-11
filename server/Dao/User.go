package Dao

import (
	"server/Model"
)

func (mgr manager) Register(userinfo Model.User) error {
	result := mgr.db.Create(&userinfo)
	return result.Error
}

func (mgr manager) GetUser() ([]Model.User, error) {
	var users []Model.User
	if err := mgr.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (mgr manager) IsExist(username string) (Model.User, error) {
	var userinfo Model.User
	result := mgr.db.Model(&Model.User{}).Where("username=?", username).Find(&userinfo)

	return userinfo, result.Error
}
