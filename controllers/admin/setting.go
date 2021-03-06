package admin

import (
	"github.com/beego/beego/v2/client/orm"
	"go-blog/models/admin"
)

type SettingController struct {
	BaseController
}

func (c *SettingController) Add() {

	o := orm.NewOrm()
	var setting []*admin.Setting
	o.QueryTable(new(admin.Setting)).All(&setting)

	for _, v := range setting {
		c.Data[v.Name] = v.Value
	}

	c.TplName = "admin/setting.html"
}

func (c *SettingController) Save() {

	response := make(map[string]interface{})

	title := c.GetString("title")
	keyword := c.GetString("keyword")
	description := c.GetString("description")
	name := c.GetString("name")
	tag := c.GetString("tag")
	//remark := c.GetString("remark")
	remark_markdown_doc := c.GetString("remark-markdown-doc")
	remark_html_code := c.GetString("remark-markdown-doc")
	image := c.GetString("image")

	o := orm.NewOrm()
	_, _ = o.Begin()

	_, err := o.Delete(&admin.Setting{Name: "name"})
	_, err = o.Delete(&admin.Setting{Name: "title"})
	_, err = o.Delete(&admin.Setting{Name: "tag"})
	_, err = o.Delete(&admin.Setting{Name: "remark"})
	_, err = o.Delete(&admin.Setting{Name: "image"})
	_, err = o.Delete(&admin.Setting{Name: "keyword"})
	_, err = o.Delete(&admin.Setting{Name: "description"})
	_, err = o.Delete(&admin.Setting{Name: "remark_markdown_doc"})
	_, err = o.Delete(&admin.Setting{Name: "remark_html_code"})

	settings := []admin.Setting{
		{Name: "title", Value: title},
		{Name: "name", Value: name},
		{Name: "tag", Value: tag},
		{Name: "remark_markdown_doc", Value: remark_markdown_doc},
		{Name: "remark_html_code", Value: remark_html_code},
		{Name: "image", Value: image},
		{Name: "keyword", Value: keyword},
		{Name: "description", Value: description},
	}

	num, err := o.InsertMulti(7, settings)

	if err != nil {
		//err = o.Rollback()
	} else {
		//err = o.Commit()
	}

	if err != nil {
		response["msg"] = "操作失败！"
		response["code"] = 500
		response["err"] = err.Error()
	} else {
		response["msg"] = "操作成功！"
		response["code"] = 200
		response["id"] = num
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

func (c *SettingController) Notice() {

	o := orm.NewOrm()
	var setting admin.Setting
	o.QueryTable(new(admin.Setting)).Filter("name", "notice").One(&setting)
	c.Data["Notice"] = setting.Value
	c.TplName = "admin/notice.html"
}

func (c *SettingController) NoticeSave() {

	response := make(map[string]interface{})

	notice := c.GetString("notice")

	o := orm.NewOrm()
	_, _ = o.Begin()
	_, err := o.Delete(&admin.Setting{Name: "notice"})

	num, err := o.Insert(&admin.Setting{Name: "notice", Value: notice})

	if err != nil {
		//err = o.Rollback()
	} else {
		//err = o.Commit()
	}

	if err != nil {
		response["msg"] = "操作失败！"
		response["code"] = 500
		response["err"] = err.Error()
	} else {
		response["msg"] = "操作成功！"
		response["code"] = 200
		response["id"] = num
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}
