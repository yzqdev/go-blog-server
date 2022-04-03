package db

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/sirupsen/logrus"
)

type MysqlDataBase struct {
}

func NewMysqlDataBase() DataBase {
	database := MysqlDataBase{}
	return &database
}

func (database *MysqlDataBase) GetStr() string {
	conf, err := config.NewConfig("ini", "conf/app.conf")

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	dbUser, _ := conf.String("db::dbUser")
	dbPass, _ := conf.String("db::dbPass")
	dbHost, _ := conf.String("db::dbHost")
	dbPort, _ := conf.String("db::dbPort")
	dbName, _ := conf.String("db::dbName")

	return dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&loc=Asia%2FShanghai"

}

func (database *MysqlDataBase) GetDriverName() string {
	return "mysql"
}

func (database *MysqlDataBase) GetAliasName() string {
	return "default"
}

func (database *MysqlDataBase) GetDriver() orm.DriverType {
	return orm.DRMySQL
}

func init() {
	Register("mysql", NewMysqlDataBase)
}
