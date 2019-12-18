package services

import (
	"github.com/kakaisaname/infra/base"
	"github.com/shopspring/decimal"
	"time"
)

//在这里 IAccountService 还没有被实例化 ，在 core下的accounts/service下的init函数中被实例化了    ****
var IAccountService AccountService

//用于对外暴露账户应用服务，唯一的暴露点
//check函数的作用是检查传入的参数是否为nil，如果是nil，会报错
func GetAccountService() AccountService {
	base.Check(IAccountService)
	return IAccountService
}

//账户创建对象																		****
//怎么没有账户编号  account_no  ???
type AccountCreatedDTO struct {
	UserId       string `validate:"required"` //用户id
	Username     string `validate:"required"` //用户名
	AccountName  string `validate:"required"` //账户名
	AccountType  int    //账户类型，用来区分不同类型的账户：积分账户、会员卡账户、钱包账户、红包账户
	CurrencyCode string //货币类型编码：CNY人民币，EUR欧元，USD美元 。。。
	Amount       string `validate:"numeric"`
}

//账户信息数据        																 ****
type AccountDTO struct {
	AccountNo    string          //账户编号,账户唯一标识
	AccountName  string          //账户名称,用来说明账户的简短描述,账户对应的名称或者命名，比如xxx积分、xxx零钱
	AccountType  int             //账户类型，用来区分不同类型的账户：积分账户、会员卡账户、钱包账户、红包账户
	CurrencyCode string          //货币类型编码：CNY人民币，EUR欧元，USD美元 。。。
	UserId       string          //用户编号, 账户所属用户
	Username     string          //用户名称
	Balance      decimal.Decimal //账户可用余额
	Status       int             //账户状态，账户状态：0账户初始化，1启用，2停用
	CreatedAt    time.Time       //创建时间
	UpdatedAt    time.Time       //更新时间
}

//账户的交易参与者 					***
type TradeParticipator struct {
	AccountNo string `validate:"required"` //账户编号 账户ID
	UserId    string `validate:"required"` //用户编号
	Username  string `validate:"required"` //用户名称
}

//转账对象 	交易者
type AccountTransferDTO struct {
	TradeNo     string            `validate:"required"`         //交易单号 全局不重复字符或数字，唯一性标识
	TradeBody   TradeParticipator `validate:"required"`         //交易主体
	TradeTarget TradeParticipator `validate:"required"`         //交易对象
	AmountStr   string            `validate:"required,numeric"` //交易金额,该交易涉及的金额
	Amount      decimal.Decimal   ``                            //交易金额,该交易涉及的金额
	ChangeType  ChangeType        `validate:"required,numeric"` //流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义
	ChangeFlag  ChangeFlag        `validate:"required,numeric"` //交易变化标识：-1 出账 1为进账，枚举
	Decs        string            ``                            //交易描述
}

//定义账户需要实现的接口															***
type AccountService interface {
	CreateAccount(dto AccountCreatedDTO) (*AccountDTO, error)  //创建账户接口的方法
	Transfer(dto AccountTransferDTO) (TransferedStatus, error) //转账接口方法
	StoreValue(dto AccountTransferDTO) (TransferedStatus, error)
	GetEnvelopeAccountByUserId(userId string) *AccountDTO
	GetAccount(accountNo string) *AccountDTO
}

func (this *AccountDTO) CopyTo(target *AccountDTO) {
	target.AccountNo = this.AccountNo
	target.AccountName = this.AccountName
	target.AccountType = this.AccountType
	target.CurrencyCode = this.CurrencyCode
	target.UserId = this.UserId
	target.Username = this.Username
	target.Balance = this.Balance
	target.Status = this.Status
	target.CreatedAt = this.CreatedAt
	target.UpdatedAt = this.UpdatedAt
}

//账户流水
type AccountLogDTO struct {
	LogNo           string          //流水编号 全局不重复字符或数字，唯一性标识
	TradeNo         string          //交易单号 全局不重复字符或数字，唯一性标识
	AccountNo       string          //账户编号 账户ID
	TargetAccountNo string          //账户编号 账户ID
	UserId          string          //用户编号
	Username        string          //用户名称
	TargetUserId    string          //目标用户编号
	TargetUsername  string          //目标用户名称
	Amount          decimal.Decimal //交易金额,该交易涉及的金额
	Balance         decimal.Decimal //交易后余额,该交易后的余额
	ChangeType      ChangeType      //流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义
	ChangeFlag      ChangeFlag      //交易变化标识：-1 出账 1为进账，枚举
	Status          int             //交易状态：
	Decs            string          //交易描述
	CreatedAt       time.Time       //创建时间
}
