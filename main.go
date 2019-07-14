package main

import (
	"blog/controllers"
	_ "blog/models"
	_ "blog/routers"
	"github.com/astaxie/beego"
	"strings"
)

func main() {
	initSession()
	initTemplate()
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1, y1) == 0
	})
}

func initSession(){
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteBlog"

	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}