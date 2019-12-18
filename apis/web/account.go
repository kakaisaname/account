package web

import (
	"github.com/kakaisaname/account/services"
	"github.com/kakaisaname/infra"
	"github.com/kakaisaname/infra/base"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
)

//定义web api的时候，对每一个子业务，定义统一的前缀 					***
//资金账户的根路径定义为:/account
//版本号：/v1/account

func init() {
	infra.RegisterApi(new(AccountApi))
}

type AccountApi struct {
	service services.AccountService
}

func (a *AccountApi) Init() {
	a.service = services.GetAccountService()
	groupRouter := base.Iris().Party("/v1/account")
	groupRouter.Post("/create", a.createHandler)                  //账户创建的接口
	groupRouter.Post("/transfer", a.transferHandler)              //转账的接口
	groupRouter.Get("/envelope/get", a.getEnvelopeAccountHandler) //查询红包账户的接口
	groupRouter.Get("/get", a.getAccountHandler)                  //查询账户信息的接口
}

//账户创建的接口: /v1/account/create
//POST body json
/*
{
	"UserId": "w123456",
	"Username": "测试用户1",
	"AccountName": "测试账户1",
	"AccountType": 0,
	"CurrencyCode": "CNY",
	"Amount": "100.11"
}

{
    "code": 1000,
    "message": "",
    "data": {
        "AccountNo": "1K1hrG0sQw7lDuF6KOQbMBe2o3n",
        "AccountName": "测试账户1",
        "AccountType": 0,
        "CurrencyCode": "CNY",
        "UserId": "w123456",
        "Username": "测试用户1",
        "Balance": "100.11",
        "Status": 1,
        "CreatedAt": "2019-04-18T13:26:51.895+08:00",
        "UpdatedAt": "2019-04-18T13:26:51.895+08:00"
    }
}
*/
func (a *AccountApi) createHandler(ctx iris.Context) {
	//获取请求参数，
	account := services.AccountCreatedDTO{}
	err := ctx.ReadJSON(&account) //读取request中的json
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		ctx.JSON(&r)
		logrus.Error(err)
		return
	}
	//执行创建账户的代码
	dto, err := a.service.CreateAccount(account)
	if err != nil {
		r.Code = base.ResCodeInnerServerError
		r.Message = err.Error()
		logrus.Error(err)
	}
	r.Data = dto
	ctx.JSON(&r)

}

//转账的接口 :/v1/account/transfer
/**
{
	"UserId": "w123456-1",
	"Username": "测试用户1",
	"AccountName": "测试账户1",
	"AccountType": 0,
	"CurrencyCode": "CNY",
	"Amount": "100.11"
}
{
	"UserId": "w123456-2",
	"Username": "测试用户2",
	"AccountName": "测试账户2",
	"AccountType": 0,
	"CurrencyCode": "CNY",
	"Amount": "100.11"
}
{
	"TradeNo": "trade123456",
	"TradeBody": {
		"AccountNo": "1K5YdR5Cng5FsBaF95fkcRJE08v",
		"UserId": "w123456-2",
		"Username": "测试用户2"
	},
	"TradeTarget": {
		"AccountNo": "1K5iy4IzhyywntWMeVlxKdxVn4G",
		"UserId": "w123456-1",
		"Username": "测试用户1"
	},
	"AmountStr": "1",

	"ChangeType": -1,
	"ChangeFlag": -1,
	"Decs": "转出"
}
*/
func (a *AccountApi) transferHandler(ctx iris.Context) {
	//获取请求参数，
	account := services.AccountTransferDTO{}
	err := ctx.ReadJSON(&account)
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		ctx.JSON(&r)
		logrus.Error(err)
		return
	}
	//执行转账逻辑
	status, err := a.service.Transfer(account)
	if err != nil {
		r.Code = base.ResCodeInnerServerError
		r.Message = err.Error()
		logrus.Error(err)
	}
	if status != services.TransferedStatusSuccess {
		r.Code = base.ResCodeBizError
		r.Message = err.Error()
	}
	r.Data = status
	ctx.JSON(&r)
}

//查询红包账户的web接口: /v1/account/envelope/get
func (a *AccountApi) getEnvelopeAccountHandler(ctx iris.Context) {
	userId := ctx.URLParam("userId")
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if userId == "" {
		r.Code = base.ResCodeRequestParamsError
		r.Message = "用户ID不能为空"
		ctx.JSON(&r)
		return
	}
	account := a.service.GetEnvelopeAccountByUserId(userId)
	r.Data = account
	ctx.JSON(&r)
}

//查询账户信息的web接口：/v1/account/get
func (a *AccountApi) getAccountHandler(ctx iris.Context) {
	accountNo := ctx.URLParam("accountNo")
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if accountNo == "" {
		r.Code = base.ResCodeRequestParamsError
		r.Message = "账户编号不能为空"
		ctx.JSON(&r)
		return
	}
	account := a.service.GetAccount(accountNo)
	r.Data = account
	ctx.JSON(&r)
}
