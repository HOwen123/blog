package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	user    models.User
	IsLogin bool
}

func (this *BaseController) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI
	u, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	this.IsLogin = false
	if ok {
		this.user = u
		this.IsLogin = true
		this.Data["User"] = this.user
	}
}

func (this *BaseController) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}
