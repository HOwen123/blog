package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.UserController{})
}
