package main

import (
	"flag"
	"gin-template/conf"
	"gin-template/services"
	"log"
)

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "c", "./conf/prod.yaml", "config path, eg: -c config.yaml")
	flag.Parse()

	//获取配置信息
	err := conf.NewConfig(configFilePath)
	if err != nil {
		log.Fatalf("config file init fail, err:%s\n", err)
	}
	//fmt.Println(conf.Get())
	services.Run(conf.Get())
}
