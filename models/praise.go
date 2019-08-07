package models

import (
	"blog/syserror"
	"github.com/jinzhu/gorm"
)

type PraiseLog struct {
	Model
	UserId int    //点赞用户id
	Key    string //文章或评论的key
	Table  string //点赞的类型（评论为messages,文章为notes）
	Flag   bool   // 是否点赞
}

type TempPraise struct {
	Praise int
}

//核心方法
func UpdatePraise(table, key string, userid int) (pcnt int, err error) {
	//开启事务
	d := db.Begin()
	defer func() {
		if err != nil {
			d.Rollback()
		} else {
			d.Commit()
		}
	}()

	var p PraiseLog
	err = d.Model(&PraiseLog{}).Where("`key`= ? and `table` = ? and user_id = ?", key, table, userid).Take(&p).Error

	if err == gorm.ErrRecordNotFound {
		p = PraiseLog{
			Key:    key,
			Table:  table,
			UserId: userid,
			Flag:   false,
		}
	} else if err != nil {
		return 0, err
	}

	if p.Flag {
		return 0, syserror.HasPraiseError{}
	}

	// 更新点赞为true
	p.Flag = true
	if err = d.Save(&p).Error; err != nil {
		return 0, err
	}

	//更新文章或留言表的点赞数量
	var ppp TempPraise
	err = d.Table(table).Where("key = ?", key).Select("praise").Scan(&ppp).Error
	if err != nil {
		return 0, err
	}
	pcnt = ppp.Praise + 1
	if err = d.Table(table).Where("key = ?", key).Update("praise", pcnt).Error; err != nil {
		return 0, err
	}
	return pcnt, nil
}
