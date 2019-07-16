package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string `gorm: "unique_index"`
	Email  string `gorm: "unique_index"`
	Pwd    string
	Avatar string
	Role   int `gorm: "default:1"` //0为管理员，1代表正常用户
}

func QueryByEmailAndPwd(email, password string) (user User, err error) {
	return user, db.Where("email = ? and Pwd = ?", email, password).Take(&user).Error
}

func QueryUserByName(name string) (user User, err error) {
	return user, db.Where("name = ?", name).Take(&user).Error
}

func QueryUserByEmail(email string) (user User, err error) {
	return user, db.Where("email = ?", email).Take(&user).Error
}

func SaveUser(user *User) (error){
	return db.Create(user).Error
}