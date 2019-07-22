package controllers

import "blog/models"

type IndexController struct {
	BaseController
}

// @router / [get]
func (this *IndexController) GetIndex() {
	//每页显示10条数据
	limit := 3
	page, err := this.GetInt("page", 1)
	if err != nil || page <= 0 {
		page = 1
	}
	title := this.GetString("title", "")
	//根据当前页和每页显示的行数，得到文章列表数据集
	notes, err := models.QueryNotesByPage(title, page, limit)
	if err != nil {
		this.Abort500(err)
	}
	//将数据传到模板页面index.html，等待渲染
	this.Data["notes"] = notes
	//得到文章的总行数
	count, err := models.QueryNotesCount(title)
	if err != nil {
		this.Abort500(err)
	}
	totPage := count / limit
	if count%limit != 0 {
		totPage = totPage + 1
	}
	//将总页数 当前页 传到模板页面。等待渲染
	this.Data["totPage"] = totPage
	this.Data["page"] = page
	this.Data["title"] = title
	this.TplName = "index.html"
}

// @router /message [get]
func (this *IndexController) GetMessage() {
	this.TplName = "message.html"
}

// @router /about [get]
func (this *IndexController) GetAbout() {
	this.TplName = "about.html"
}

// @router /details [get]
func (this *IndexController) GetDetails() {
	this.TplName = "details.html"
}

// @router /comment [get]
func (this *IndexController) GetComment() {
	this.TplName = "comment.html"
}

// @router /error [get]
func (this *IndexController) GetError() {
	this.TplName = "error/404.html"
}

// @router /user [get]
func (this *IndexController) GetUser() {
	this.TplName = "user.html"
}

// @router /reg [get]
func (this *IndexController) GetReg() {
	this.TplName = "reg.html"
}
