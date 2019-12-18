package main

import (
	"github.com/kakaisaname/infra"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

func main() {
	//获取程序运行文件所在的路径     获取配置文件
	file := kvs.GetCurrentFilePath("config.ini", 1)
	//加载和解析配置文件																**
	conf := ini.NewIniFileConfigSource(file)

	//base.InitLog(conf)															输出日志到日志文件

	//初始化  **
	//返回的是 BootApplication结构体
	app := infra.New(conf)
	app.Start()
	c := make(chan int)
	<-c
}
