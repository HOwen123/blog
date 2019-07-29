package models

import (
	"fmt"
)

type Note struct {
	Model
	Key     string `gorm:"unique_index;not null;" json:"key"`
	UserID  int     `json:"userId"` //用户Id
	User    User   //用户
	Title   string //文章标题
	Summary string `gorm:"type:text"` //概要
	Content string `gorm:"type:text"` //文章内容
	Visit   int    `gorm:"default:0"` //浏览次数
	Praise  int    `gorm:"default:0"` //点赞次数
}

func QueryNoteByKeyAndUserId(key string, userId int) (note Note, err error) {
	return note, db.Model(&Note{}).Where("Key = ? and user_id = ?", key, userId).Take(&note).Error
}

func SaveNote(n *Note) error {
	return db.Save(n).Error
}

func QueryNotesByPage(title string, page, limit int) (note []*Note, err error) {
	//offset从第几行开始，limit:返回多少数据
	return note, db.Where("title like ?", fmt.Sprintf("%%%s%%", title)).
		Offset((page - 1) * limit).Limit(limit).Find(&note).Error
}

//查询文章的总数量
func QueryNotesCount(title string) (count int, err error) {
	return count, db.Model(&Note{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Count(&count).Error
}

func QueryNoteByKey(key string) (note Note,err error){
	return note, db.Model(&Note{}).Where("Key = ?",key).Take(&note).Error
}
