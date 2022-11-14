package Dao

import (
	"fmt"
	"log"
	"server/Model"
	"server/Utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Manager interface {
	Register(userinfo Model.User) error
	GetUser(userId uint64) (Model.User, error)
	IsExist(username string) (Model.User, error)
	GetPostListByUserId(userId uint64) ([]Model.Post, error)
	UpdateProfileById(userId uint64, data *Model.User) error

	GetCommunityList() ([]Model.Community, error)

	StoreArticle(postInfo Model.Post) error
	GetList(latestTime int64) ([]Model.Post, error)
	GetPostById(postId int64) (Model.Post, error)
	GetUsernameById(AuthorId uint64) (author string, err error)
	GetCommunityNameById(CommunityId uint64) (CommunityName string, err error)
	DeletePostById(postId int64) error

	AddComment(comment *Model.Comment)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Utils.DbUser,
		Utils.DbPassWord, Utils.DbHost, Utils.DbPort, Utils.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			// use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}

	if err := db.AutoMigrate(&Model.Post{}); err != nil {
		log.Fatal(err)
	}
}
