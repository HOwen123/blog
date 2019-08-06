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

// @router /count [get]
func (this *MessageController) Count() {
	//查询 留言的总数量
	count, err := models.QueryMessagesCountByNoteKey("")
	if err != nil {
		this.Abort500(syserror.New("查询失败!", err))
	}
	// 将留言的总数量 返回给前台
	this.JsonOkH("查询成功", H{"count": count})
}

// @router /query [get]
func (this *MessageController) Query() {
	//获得第几页，默认第一页
	pageno, err := this.GetInt("pageno", 1)
	if err != nil || pageno < 1 {
		pageno = 1
	}
	//获得每页显示多少条数据，默认10条
	pagesize, err := this.GetInt("pagesize", 10)
	if err != nil {
		pagesize = 10
	}
	//调用数据库方法，查询出留言的数据集
	ms, err := models.QueryPageMessagesByNoteKey("", pageno, pagesize)
	if err != nil {
		this.Abort500(syserror.New("查询失败", err))
	}
	this.JsonOkH("查询成功", H{"data": ms})
}
