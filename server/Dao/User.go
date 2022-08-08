package Dao

import (
	"server/Model"
)

func (mgr manager) AddUser(user *Model.User) {
	mgr.db.Create(user)
}

func (mgr manager) GetUser() ([]Model.User, error) {
	var users []Model.User
	if err := mgr.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
