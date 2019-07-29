package models

type User struct {
	Model
	Name   string `gorm: "unique_index" json:"name"`
	Email  string `gorm: "unique_index" json:"email"`
	Pwd    string `json:"-"`//“-”
	Avatar string `json:avatar`
	Role   int    `gorm: "default:0" json:"role"` //0为管理员，1代表正常用户
}

func QueryByEmailAndPwd(email, password string) (user User, err error) {
	return user, db.Model(&User{}).Where("email = ? and Pwd = ?", email, password).Take(&user).Error
}

func QueryUserByName(name string) (user User, err error) {
	return user, db.Model(&User{}).Where("name = ?", name).Take(&user).Error
}

func QueryUserByEmail(email string) (user User, err error) {
	return user, db.Model(&User{}).Where("email = ?", email).Take(&user).Error
}

func SaveUser(user *User) error {
	return db.Model(&User{}).Create(user).Error
}
