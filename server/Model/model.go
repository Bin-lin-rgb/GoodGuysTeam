package Model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int64  `json:"id"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Gender   int8   `json:"gender,omitempty"`
}

type Comment struct {
	gorm.Model
	Id       int64  `json:"id,omitempty"`
	Content  string `json:"content,omitempty"`
	PostId   int64  `json:"post_id,omitempty"`
	AuthorId int64  `json:"author_id,omitempty"`
	ParentId int64  `json:"parent_id,omitempty"`
	Status   uint8  `json:"status,omitempty"`
}

type Post struct {
	gorm.Model
	Id          int64  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`        // 标题
	Content     string `json:"content,omitempty"`      // 内容
	AuthorId    int64  `json:"author_id,omitempty"`    // 作者的用户id
	CommunityId int64  `json:"community_id,omitempty"` // 所属社区
	Status      int8   `json:"status,omitempty"`       // 帖子状态
}
