package common

import (
	. "TikTokLite/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func DBInit() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	creatTable()
}

func creatTable() {
	DB.AutoMigrate(&User{}, &Video{}, &Favorite{}, &Comment{}, &Relation{}, &Message{})
}

func GetDB() *gorm.DB {
	return DB
}
func CloseDataBase() {
	DB.Close()
}
