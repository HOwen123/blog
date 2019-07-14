package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.IndexController{})
}
