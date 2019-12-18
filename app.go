package main

import (
	_ "account/apis/web"
	_ "account/core/accounts"
	"github.com/kakaisaname/infra"
	"github.com/kakaisaname/infra/base"
)

//注册我们的starter  启动器
func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	//infra.Register(&accounts.AccountClientStarter{})
	infra.Register(&base.HookStarter{})
}
