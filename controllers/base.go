package controllers

import (
	"blog/models"
	"errors"
	"github.com/astaxie/beego"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	user    models.User
	IsLogin bool
}

func (ctx *BaseController) Prepare() {
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	u, ok := ctx.GetSession(SESSION_USER_KEY).(models.User)
	ctx.IsLogin = false
	if ok {
		ctx.user = u
		ctx.IsLogin = true
		ctx.Data["User"] = ctx.user
	}
}

func (ctx *BaseController) Abort500(err error) {
	ctx.Data["error"] = err
	ctx.Abort("500")
}

func (ctx *BaseController) GetMustString(key,message string) string{
	value := ctx.GetString(key,"")
	if len(value) == 0 {
		ctx.Abort500(errors.New(message))
	}
	return value
}
