package Dao

import "server/Model"

func (mgr manager) AddComment(comment *Model.Comment) {
	mgr.db.Create(comment)
}
