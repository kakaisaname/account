package accounts

import (
	_ "account/testx"
	"database/sql"
	"github.com/kakaisaname/infra/base"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/dbx"
	"testing"
)

//测试是通过了的
func TestAccountDao_GetOne(t *testing.T) {
	//都在一个事务后面去执行 						****
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &AccountDao{
			runner: runner,
		}
		//Convey 函数，第一个参数是描述，第二个参数为t，第三个参数就是一个函数
		Convey("通过编号查询账户数据", t, func() {
			a := &Account{
				Balance:     decimal.NewFromFloat(100),
				Status:      1,
				AccountNo:   ksuid.New().Next().String(),
				AccountName: "测试资金账户",
				UserId:      ksuid.New().Next().String(),
				Username:    sql.NullString{String: "测试用户", Valid: true},
			}
			//插入account
			id, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)
			//获取一条数据
			na := dao.GetOne(a.AccountNo)
			So(na, ShouldNotBeNil)
			So(na.Balance.String(), ShouldEqual, a.Balance.String())
			So(na.CreatedAt, ShouldNotBeNil)
			So(na.UpdatedAt, ShouldNotBeNil)

		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}

}

func TestAccountDao_GetByUserId(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &AccountDao{
			runner: runner,
		}
		Convey("通过用户ID和账户类型查询账户数据", t, func() {
			a := &Account{
				Balance:     decimal.NewFromFloat(100),
				Status:      1,
				AccountNo:   ksuid.New().Next().String(),
				AccountName: "测试资金账户",
				UserId:      ksuid.New().Next().String(),
				Username:    sql.NullString{String: "测试用户", Valid: true},
				AccountType: 2,
			}
			id, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)

			na := dao.GetByUserId(a.UserId, a.AccountType)
			So(na, ShouldNotBeNil)
			So(na.Balance.String(), ShouldEqual, a.Balance.String())
			So(na.CreatedAt, ShouldNotBeNil)
			So(na.UpdatedAt, ShouldNotBeNil)

		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}

}

//最好一个Convey 下面最多再有一个Convey，不然的话，如果每多一个，就会再执行一遍最外层的Convey代码 					****
func TestAccountDao_UpdateBalance(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &AccountDao{
			runner: runner,
		}
		balance := decimal.NewFromFloat(100)
		Convey("更新账户余额", t, func() {
			a := &Account{
				Balance:     balance,
				Status:      1,
				AccountNo:   ksuid.New().Next().String(),
				AccountName: "测试资金账户",
				UserId:      ksuid.New().Next().String(),
				Username:    sql.NullString{String: "测试用户", Valid: true},
			}
			id, err := dao.Insert(a)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)

			//1.增加余额
			Convey("增加余额", func() {
				amount := decimal.NewFromFloat(10)
				rows, err := dao.UpdateBalance(a.AccountNo, amount)
				So(err, ShouldBeNil)
				So(rows, ShouldEqual, 1)
				na := dao.GetOne(a.AccountNo)
				newBalance := balance.Add(amount)
				So(na, ShouldNotBeNil)
				So(na.Balance.String(), ShouldEqual, newBalance.String())
			})

			//2.扣减余额，余额足够
			//把updateBalance余额扣减的第二个测试用例
			//作为作业留给同学们来编写

			//3.扣减余额，余额不够
			Convey("扣减余额，余额不够", func() {
				a1 := dao.GetOne(a.AccountNo)
				So(a1, ShouldNotBeNil)
				amount := decimal.NewFromFloat(-300)
				rows, err := dao.UpdateBalance(a.AccountNo, amount)
				So(err, ShouldBeNil)
				So(rows, ShouldEqual, 0)
				a2 := dao.GetOne(a.AccountNo)
				So(a2, ShouldNotBeNil)
				So(a1.Balance.String(), ShouldEqual, a2.Balance.String())
			})

		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}
