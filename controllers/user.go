package controllers

import (
	"blog/models"
	"blog/syserror"
)

type UserController struct {
	BaseController
}

//@router /login [post]
func (userController *UserController) Login(){
	email := userController.GetMustString("email","email不能为空！")
	password := userController.GetMustString("password","密码不能为空！")
	user, err := models.QueryByEmailAndPwd(email, password)
	if err != nil {
		userController.Abort500(syserror.New("登录失败！",err))
	}
	userController.SetSession(SESSION_USER_KEY,user)
	userController.Data["json"] = map[string]interface{}{
		"code":0,
		"action":"/",
	}
	userController.ServeJSON()
}
