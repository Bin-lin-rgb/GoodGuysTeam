package Model

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id            int64  `json:"id"`
	UserId        uint64 `json:"user_id"`
	Username      string `json:"username"`
	Password      string `json:"password,omitempty"`
	Email         string `json:"email"`
	Gender        int8   `json:"gender"`
	Company       string `gorm:"column:company" json:"company"`               //公司
	Position      string `gorm:"column:position" json:"position"`             //职位
	UserPage      string `gorm:"column:user_page" json:"user_page"`           //个人主页
	UserIntroduce string `gorm:"column:user_introduce" json:"user_introduce"` //个人介绍
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
	PostId      uint64 `json:"post_id"`
	Title       string `json:"title,omitempty" binding:"required"`        // 标题
	Content     string `json:"content,omitempty" binding:"required"`      // 内容
	AuthorId    uint64 `json:"author_id,omitempty"`                       // 作者的用户id
	CommunityId uint64 `json:"community_id,omitempty" binding:"required"` // 所属社区
	Status      int8   `json:"status,omitempty"`                          // 帖子状态
}

// UnmarshalJSON 为Post类型实现自定义的UnmarshalJSON方法
func (p *Post) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Title       string `json:"title"`
		Content     string `json:"content"`
		CommunityID int64  `json:"community_id"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Title) == 0 {
		err = errors.New("帖子标题不能为空")
	} else if len(required.Content) == 0 {
		err = errors.New("帖子内容不能为空")
	} else if required.CommunityID == 0 {
		err = errors.New("未指定版块")
	} else {
		p.Title = required.Title
		p.Content = required.Content
		p.CommunityId = uint64(required.CommunityID)
	}
	return
}

type Community struct {
	Id            int32  `json:"id,omitempty"`
	CommunityId   uint64 `json:"community_id,omitempty"`
	CommunityName string `json:"community_name,omitempty"`
	Introduction  string `json:"introduction,omitempty"`
}

type ApiPostList struct {
	PostId        int64  `json:"post_id,omitempty"`
	Title         string `json:"title,omitempty"`          // 标题
	Content       string `json:"content,omitempty"`        // 内容
	CommunityName string `json:"community_name,omitempty"` // 嵌入社区分类信息
	AuthorName    string `json:"author_name,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	CreatedAt2    string `json:"created_at2,omitempty"`
}
