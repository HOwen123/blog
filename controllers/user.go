package controllers

import (
	"blog/syserror"
	"errors"
	"strings"
)

type UserController struct {
	BaseController
}

//@router /login [post]
func (userController *UserController) Login() {
	email := userController.GetMustString("email", "email不能为空！")
	password := userController.GetMustString("password", "密码不能为空！")
	user, err := userController.Dao.QueryByEmailAndPwd(email, password)
	if err != nil {
		userController.Abort500(syserror.New("登录失败！", err))
	}
	userController.SetSession(SESSION_USER_KEY, user)
	userController.Data["json"] = map[string]interface{}{
		"code":   0,
		"action": "/",
	}
	userController.ServeJSON()
}

// router /reg [post]
func (userController *UserController) Reg(){
	name := userController.GetMustString("name", "昵称不能为空！")
	email := userController.GetMustString("email", "邮箱不能为空！")
	password := userController.GetMustString("password", "密码不能为空！")
	password2 := userController.GetMustString("password2", "确认密码不能为空！")

	if strings.Compare(password,password2)!=0 {
		userController.Abort500(errors.New("两次密码必须一致！"))
	}

	if u,err := userController.Dao.QueryUserByName(name); err == nil && u.ID != 0{
		userController.Abort500(errors.New("用户昵称已经存在！"))
	}

}
