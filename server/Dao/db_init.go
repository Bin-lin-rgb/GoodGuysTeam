package Dao

import (
	"log"
	"server/Model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Manager interface {
	Register(userinfo Model.User) error
	GetUser() ([]Model.User, error)
	IsExist(username string) (Model.User, error)

	AddComment(comment *Model.Comment)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:Admin666@tcp(127.0.0.1:3306)/blogserver?charset=utf8mb4&parseTime=True&loc=Local"
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
