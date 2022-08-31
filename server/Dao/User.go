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
	if err := mgr.db.Select("user_id", "username", "company", "user_page", "user_introduce", "position", "created_at").
		Where("user_id=?", userId).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (mgr manager) IsExist(username string) (Model.User, error) {
	var userinfo Model.User
	result := mgr.db.Where("username=?", username).Find(&userinfo)

	return userinfo, result.Error
}

func (mgr manager) UpdateProfileById(userId uint64, data *Model.User) error {
	var userinfo Model.User

	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["user_page"] = data.UserPage
	maps["company"] = data.Company
	maps["user_introduce"] = data.UserIntroduce
	maps["position"] = data.Position
	err := mgr.db.Model(&userinfo).Where("user_id = ? ", userId).Updates(maps).Error

	return err
}
