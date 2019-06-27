package dao

import (
	"github.com/chulinshao/rehab/models"
	"github.com/go-xorm/xorm"
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
	account := models.DoctorAccount{Code: code}
	dao.engine.Id(code).Get(&account)
	return account
}
