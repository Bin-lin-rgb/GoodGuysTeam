package Dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"server/Model"
)

type Manager interface {
	AddUser(user *Model.User)
	GetUser() ([]Model.User, error)

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
