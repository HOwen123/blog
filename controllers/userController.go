package controllers

import (
	"blog/models"
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
	user, err := models.QueryByEmailAndPwd(email, password)
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

// @router /reg [post]
func (userController *UserController) Reg() {
	name := userController.GetMustString("name", "昵称不能为空！")
	email := userController.GetMustString("email", "邮箱不能为空！")
	password := userController.GetMustString("password", "密码不能为空！")
	password2 := userController.GetMustString("password2", "确认密码不能为空！")

	if strings.Compare(password, password2) != 0 {
		userController.Abort500(errors.New("两次密码必须一致！"))
	}

	if u, err := models.QueryUserByName(name); err == nil && u.ID > 0 {
		userController.Abort500(errors.New("用户昵称已经存在！"))
	}

	if u, err := models.QueryUserByEmail(email); err == nil && u.ID > 0 {
		userController.Abort500(errors.New("邮箱已被注册"))
	}

	if err := models.SaveUser(&models.User{
		Name:   name,
		Email:  email,
		Avatar: "",
		Pwd:    password,
		Role:1, // 0 管理员 1正常用户
	}); err != nil {
		userController.Abort500(syserror.New("注册失败！",err))
	}
	userController.JsonOK("注册成功","/")
}

// @router /logout [get]
func (userController *UserController) Logout(){
	userController.MustLogin()
	userController.DelSession(SESSION_USER_KEY)
	userController.Redirect("/",302)
}
