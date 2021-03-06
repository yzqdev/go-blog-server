package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	_ "go-blog/routers"
	db "go-blog/service/databsae"
	"go-blog/utils"
	"go-blog/utils/sitemap"
)

func init() {
	conf, err := config.NewConfig("ini", "conf/app.conf")

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	database, _ := db.NewDataBase(conf.String("db::dbType"))
	orm.RegisterDriver(database.GetDriverName(), database.GetDriver())
	orm.RegisterDataBase(database.GetAliasName(), database.GetDriverName(), database.GetStr())

	beego.AddFuncMap("IndexForOne", utils.IndexForOne)
	beego.AddFuncMap("IndexAddOne", utils.IndexAddOne)
	beego.AddFuncMap("IndexDecrOne", utils.IndexDecrOne)
	beego.AddFuncMap("StringReplace", utils.StringReplace)
	beego.AddFuncMap("TimeStampToTime", utils.TimeStampToTime)

	// 每天0点定时更新站点地图
	go func() {
		c := cron.New()
		//*/1 0 * * *
		// 0 0 * * *
		c.AddFunc("*/1 * * * *", func() {
			url, _ := conf.String("url")
			sitemap.Sitemap("./", url)
		})
		c.Start()
	}()

}

func main() {
	//bee generate appcode -tables="cron" -driver=mysql -conn="root:root@tcp(127.0.0.1:3306)/blog" -level=3
	beego.Run()
}
