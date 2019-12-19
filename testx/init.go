package testx

import (
	"github.com/kakaisaname/infra"
	"github.com/kakaisaname/infra/base"
	"github.com/kakaisaname/props/ini"
	"github.com/kakaisaname/props/kvs"
)

func init() {
	//获取程序运行文件所在的路径
	file := kvs.GetCurrentFilePath("../brun/config.ini", 1)
	//加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)

	//

	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})

	app := infra.New(conf)
	app.Start()
}
