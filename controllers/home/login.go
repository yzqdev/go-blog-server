package home

import (
	beego "github.com/beego/beego/v2/server/web"
	"go-blog/models"
	"go-blog/utils"
)

type LoginController struct {
	beego.Controller
}

func (ctl *LoginController) Sign() {
	ctl.TplName = "login.html"
}
func (ctl *LoginController) Login() {

	username := ctl.GetString("username")
	password := ctl.GetString("password")

	password = utils.PasswordMD5(password, username)

	response := make(map[string]interface{})

	if customer, ok := models.CustomerLogin(username, password); ok {
		ctl.SetSession("Customer", *customer)
		response["code"] = 200
		response["msg"] = "登录成功！"
	} else {
		response["code"] = 500
		response["msg"] = "登录失败！"
	}
	ctl.Data["json"] = response
	ctl.ServeJSON()

}
