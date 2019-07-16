package controllers

import (
	"blog/syserror"
	"github.com/astaxie/beego/logs"
)

type ErrorController struct {
	BaseController
}

// ajax: {code:, msg:, reason:error}
func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
	logs.Info("这里是error404")
	if c.IsAjax() {
		c.jsonError(syserror.Error404{})
	}
}

 func (c *ErrorController) Error500() {
	c.TplName = "error/500.html"
	err, ok := c.Data["error"].(error)
	if !ok {
		err = syserror.New("未知错误", nil)
	}
	serr, ok := err.(syserror.Error)
	if !ok {
		serr = syserror.New(err.Error(), nil)
	}
	if serr.ReasonError() != nil {
		logs.Info(serr.Error(), serr.ReasonError())
	}
	if c.IsAjax() {
		c.jsonError(serr)
	}else {
		c.Data["content"] = serr.Error()
	}
}

func (c *ErrorController) jsonError(sysErr syserror.Error) {
	c.Ctx.Output.Status = 200
	c.Data["json"] = map[string]interface{}{
		"code": sysErr.Code(),
		"msg":  sysErr.Error(),
	}
	c.ServeJSON()
}