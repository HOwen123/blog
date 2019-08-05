package controllers

import (
	"blog/models"
	"blog/syserror"
	"github.com/astaxie/beego/logs"
)

//评论处理的控制器
type MessageController struct {
	BaseController
}

//新增评论的控制器方法
// @router /new/?:key [post]
func (mc *MessageController) NewMessage() {
	//登录用户
	mc.MustLogin()
	// 获取文章的Key
	key := mc.Ctx.Input.Param(":key")
	logs.Info("content:" + mc.GetString("content"))
	// 获取评论的内容
	content := mc.GetMustString("content", "请输入内容！")
	logs.Info("key:" + key)

	// 新增一个评论的唯一key
	k := mc.UUID()
	// 定义评论结构体，保存评论
	m := &models.Message{
		Key:     k,
		NoteKey: key,
		User:    mc.user,
		UserId:  int(mc.user.ID),
		Content: content,
	}
	if err := models.SaveMessage(m); err != nil {
		//报错失败，提示页面错误
		mc.Abort500(syserror.New("保存失败", err))
	}
	mc.JsonOkH("保存成功！", H{"data": m})
}
