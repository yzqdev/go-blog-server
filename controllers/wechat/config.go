package wechat

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/sirupsen/logrus"
)

var rHost, _ = beego.AppConfig.String("redis::rHost")
var rPort, _ = beego.AppConfig.String("redis::rPort")
var redisCache = cache.NewRedis(&cache.RedisOpts{

	Host: rHost + ":" + rPort,
})
var token, _ = beego.AppConfig.String("wechat::Token")
var appId, _ = beego.AppConfig.String("wechat::AppID")
var appSecret, _ = beego.AppConfig.String("wechat::AppSecret")
var encodingAes, _ = beego.AppConfig.String("wechat::EncodingAESKey")
var config = &wechat.Config{
	AppID:          appId,
	AppSecret:      appSecret,
	Token:          token,
	EncodingAESKey: encodingAes,
	Cache:          redisCache,
}

var log = logrus.New()
