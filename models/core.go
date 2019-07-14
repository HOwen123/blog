package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	_, err = os.Stat("data")
	if err != nil {
		if err = os.Mkdir("data", 0777); err != nil {
			panic("failed :" + err.Error())
		}
	}
	db, err = gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	//如果数据库里面没用用户数据，我们新增一条admin记录
	var count int
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&User{
			Name:   "admin",
			Email:  "admin@qq.com",
			Pwd:    "123",
			Avatar: "/static/img/info-img.png",
			Role:   0,
		})
	}
}