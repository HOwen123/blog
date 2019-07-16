package controllers

import (
	"blog/models"
	"errors"
	"github.com/astaxie/beego"
	"liteblog/syserrors"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	user    models.User
	IsLogin bool
	Dao     *models.DB
}

func (ctx *BaseController) Prepare() {
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	u, ok := ctx.GetSession(SESSION_USER_KEY).(models.User)
	ctx.IsLogin = false
	if ok {
		ctx.user = u
		ctx.IsLogin = true
		ctx.Data["User"] = ctx.user
		ctx.Data["isLogin"] = ctx.IsLogin
	}
}

func (ctx *BaseController) JsonOK(msg string, action string) {
	//var action string
	//if len(actions)>0 {
	//	action = actions[0]
	//}
	ctx.Data["json"] = map[string]interface{}{
		"code":   0,
		"msg":    msg,
		"action": action,
	}
	ctx.ServeJSON()
}

func (ctx *BaseController) Abort500(err error) {
	ctx.Data["error"] = err
	ctx.Abort("500")
}

func (ctx *BaseController) GetMustString(key, message string) string {
	value := ctx.GetString(key, "")
	if len(value) == 0 {
		ctx.Abort500(errors.New(message))
	}
	return value
}


func (ctx *BaseController) MustLogin() {
	if !ctx.IsLogin {
		ctx.Abort500(syserrors.NoUserError{})
	}
}
