package Dao

import "server/Model"

func (mgr manager) GetCommunityList() ([]Model.Community, error) {
	var list []Model.Community
	if err := mgr.db.Model(Model.Community{}).Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}
