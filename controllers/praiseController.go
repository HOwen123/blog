package controllers

import (
	"blog/models"
	"blog/syserror"
	"errors"
)

type PraiseController struct {
	BaseController
}

func (this *PraiseController) NextPrepare() {
	this.MustLogin()
}

//@router /:type/:key [post]
func (this *PraiseController) Praise() {
	// 通过type 辨别点赞类型
	ttype := this.Ctx.Input.Param(":type")
	// 获取key
	key := this.Ctx.Input.Param(":key")

	//定义table变量
	table := "notes"
	switch ttype {
	case "message":
		table = "messages"
	case "note":
		table = "notes"
	default:
		this.Abort500(errors.New("未知类型"))
	}
	//更新点赞的方法
	pcnt, err := models.UpdatePraise(table, key, int(this.user.ID))
	if err != nil {
		if e2, ok := err.(syserror.HasPraiseError); ok {
			this.Abort500(e2)
		}
		this.Abort500(syserror.New("点赞失败", err))
	}
	//点赞成功，返回点赞数量
	this.JsonOkH("点击成功", H{"praise": pcnt})
}
