package dao

import (
	"github.com/chulinshao/rehab/models"
	"github.com/chulinshao/rehab/util"
	"github.com/go-xorm/xorm"
	"log"
)

type AccountDao struct {
	engine *xorm.Engine
}

func NewAccountDao(engine *xorm.Engine) *AccountDao {
	return &AccountDao{
		engine: engine,
	}
}

func (dao AccountDao) GetByCode(code string) models.DoctorAccount {
	account := models.DoctorAccount{}
	dao.engine.Id(code).Get(&account)
	return account
}

func (dao AccountDao) UpdateAlipayAccount(doctorCode string, alipayAccount string) int64 {
	rs, err := dao.engine.Exec("update dzy_doctor_account set alipay_account = ?, alipay_modifty_time = ? where doctor_code = ?", alipayAccount, util.CurrDateStr(), doctorCode)
	if err != nil {
		log.Fatal(err)
	} else {
		rs, err := rs.RowsAffected()
		if err != nil {
			log.Fatal(err)
		} else {
			return rs
		}
	}
	return 0
}
