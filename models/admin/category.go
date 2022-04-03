package admin

import (
	"github.com/beego/beego/v2/client/orm"
)

type Category struct {
	Id     int
	Name   string
	Pid    int
	Sort   int
	Status int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Category))
}
