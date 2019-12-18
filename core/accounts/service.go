//具体的方法实现  ***
// 应用服务层 								***
package accounts

import (
	"fmt"
	"github.com/kakaisaname/account/services"
	"github.com/kakaisaname/infra/base"
	"github.com/kataras/iris/core/errors"
	"github.com/shopspring/decimal"
	"sync"
)

var _ services.AccountService = new(accountService)
var once sync.Once

// 在init函数中被实例化   services.IAccountService   once来只实例化一次
func init() {
	once.Do(func() {
		services.IAccountService = new(accountService)
	})
}

//禁止对外实例化 							***
type accountService struct {
}

func (a *accountService) CreateAccount(dto services.AccountCreatedDTO) (*services.AccountDTO, error) {
	domain := accountDomain{}
	//验证输入参数						***
	if err := base.ValidateStruct(&dto); err != nil {
		return nil, err
	}
	//验证账户是否存在和幂等性
	acc := domain.GetAccountByUserIdAndType(dto.UserId, services.AccountType(dto.AccountType))
	if acc != nil {
		return acc, errors.New(
			fmt.Sprintf("用户的该类型账户已经存在：username=%s[%s],账户类型=%d",
				dto.Username, dto.UserId, dto.AccountType))

	}
	//执行账户创建的业务逻辑
	amount, err := decimal.NewFromString(dto.Amount)
	if err != nil {
		return nil, err
	}
	account := services.AccountDTO{
		UserId:       dto.UserId,
		Username:     dto.Username,
		AccountType:  dto.AccountType,
		AccountName:  dto.AccountName,
		CurrencyCode: dto.CurrencyCode,
		Status:       1,
		Balance:      amount,
	}
	rdto, err := domain.Create(account)
	return rdto, err
}

func (a *accountService) Transfer(dto services.AccountTransferDTO) (services.TransferedStatus, error) {
	//验证参数
	domain := accountDomain{}
	//验证输入参数
	if err := base.ValidateStruct(&dto); err != nil {
		return services.TransferedStatusFailure, err
	}
	//执行转账逻辑
	amount, err := decimal.NewFromString(dto.AmountStr)
	if err != nil {
		return services.TransferedStatusFailure, err
	}
	dto.Amount = amount
	if dto.ChangeFlag == services.FlagTransferOut {
		if dto.ChangeType > 0 {
			return services.TransferedStatusFailure,
				errors.New("如果changeFlag为支出，那么changeType必须小于0")
		}
	} else {
		if dto.ChangeType < 0 {
			return services.TransferedStatusFailure,
				errors.New("如果changeFlag为收入,那么changeType必须大于0")
		}
	}

	status, err := domain.Transfer(dto)
	//转账成功，并且交易主体和交易目标不是同一个人，而且交易类型不是储值，则进行反向操作
	if status == services.TransferedStatusSuccess && dto.TradeBody.AccountNo != dto.TradeTarget.AccountNo && dto.ChangeType != services.AccountStoreValue {
		backwardDto := dto
		backwardDto.TradeBody = dto.TradeTarget
		backwardDto.TradeTarget = dto.TradeBody
		backwardDto.ChangeType = -dto.ChangeType
		backwardDto.ChangeFlag = -dto.ChangeFlag
		status, err := domain.Transfer(backwardDto)
		return status, err
	}
	return status, err
}

func (a *accountService) StoreValue(dto services.AccountTransferDTO) (services.TransferedStatus, error) {
	dto.TradeTarget = dto.TradeBody
	dto.ChangeFlag = services.FlagTransferIn
	dto.ChangeType = services.AccountStoreValue
	return a.Transfer(dto)
}

func (a *accountService) GetEnvelopeAccountByUserId(userId string) *services.AccountDTO {
	domain := accountDomain{}
	account := domain.GetEnvelopeAccountByUserId(userId)
	return account
}

func (a *accountService) GetAccount(accountNo string) *services.AccountDTO {
	domain := accountDomain{}
	return domain.GetAccount(accountNo)
}
