package controllers

import (
	"blog/models"
	"blog/syserror"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"time"
)

type NoteController struct {
	BaseController
}

func (ctx *NoteController) NestPrepare() {
	ctx.MustLogin()
	if ctx.user.Role != 0 {
		ctx.Abort500(syserror.New("您没有权限修改文章", nil))
	}
}

// @router /new [get]
func (ctx *NoteController) NewPage() {
	ctx.Data["key"] = ctx.UUID()
	ctx.TplName = "note_new.html"
}

// @router /save/:key [post]
func (ctx *NoteController) Save() {
	//得到页面传过来的 key
	key := ctx.Ctx.Input.Param(":key")
	logs.Info(key)
	// 判空,为空就返回错误
	title := ctx.GetMustString("title", "标题不能为空！")
	content := ctx.GetMustString("content", "内容不能为空！")
	//获得文章摘要
	summary, _ := getSummary(content)
	// 根据key查询文章
	logs.Info(key, int(ctx.user.ID))
	note, err := models.QueryNoteByKeyAndUserId(key, int(ctx.user.ID))

	var n models.Note
	if err != nil {
		//存在错误不是查不到数据的错误，那就返回错误
		if err != gorm.ErrRecordNotFound {
			ctx.Abort500(syserror.New("保存失败！", err))
		}
		n = models.Note{
			Key:     key,
			Summary: summary,
			Title:   title,
			Content: content,
			UserID:  int(ctx.user.ID),
		}
	} else {
		//查询不报错，这文章存在，那就更新文章操作
		n = note
		n.Title = title
		n.Content = content
		n.Summary = summary
		n.UpdatedAt = time.Now()
	}

	//保存文章 saveNote 是根据id来判断是更新还是新增，id存在就更新，不存在就新增
	//上面更新操作是从数据库查出来的文章记录，修改数据,所以是存在id的
	if err := models.SaveNote(&n); err != nil {
		ctx.Abort500(syserror.New("保存失败！", err))
	}
	ctx.JsonOK("成功", "/details/"+key)
}

func getSummary(content string) (string, error) {
	var buf bytes.Buffer
	buf.Write([]byte(content))
	//用goquery来解析
	doc, err := goquery.NewDocumentFromReader(&buf)
	if err != nil {
		return "", err
	}
	// Text() 得到body元素下的文本内容（去掉html标签元素）
	str := doc.Find("body").Text()
	//截取字符串
	if len(str) > 600 {
		str = str[0:600] + "..."
	}
	return str, nil
}
