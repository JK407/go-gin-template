package services

import (
	"fmt"
	"gin-template/cache"
	"gin-template/conf"
	"gin-template/controllers"
	"gin-template/utils/core"
	"gin-template/utils/logger"
	"gin-template/utils/mysql"

	"gin-template/utils/trace"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Run(cfg conf.Config) {
	//初始化日志
	logger.InitLog(cfg.App.LogFile)
	//初始化链路
	trace.InitTracer()

	//初始化数据库
	mysql.InitDB(cfg.MySql, cfg.App.RunMode)
	if cfg.App.RunMode == "prod" {
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
	}
	//初始化缓存
	redisClose := cache.InitRedis(&cfg)
	defer redisClose()

	// InitializeValidator
	core.InitializeValidator()

	//路由
	controllers.InitRouters(getListenHost())

}

func getListenHost() string {
	listenHost := ""
	if conf.Get().App.Host != "0.0.0.0" && conf.Get().App.Host != "127.0.0.1" {
		listenHost += conf.Get().App.Host
	}
	listenHost += ":" + strconv.Itoa(conf.Get().App.Port)
	fmt.Printf("listen server :%s\n", listenHost)
	return listenHost
}
