package dao

import (
	"github.com/chulinshao/rehab/models"
	"github.com/go-xorm/xorm"
	"log"
)

type DoctorAccountLogDao struct {
	engine *xorm.Engine
}

func NewDoctorAccountLogDao(engine *xorm.Engine) *DoctorAccountLogDao {
	return &DoctorAccountLogDao{
		engine: engine,
	}
}

func (dao DoctorAccountLogDao) CountTotalCommentary(doctorCode string, accountLogType string) float64 {
	accountLog := models.DoctorAccountLog{
		DoctorCode: doctorCode,
		LogType:    accountLogType,
	}
	count, err := dao.engine.Sum(accountLog, "balance")
	if err != nil {
		log.Fatalln(err)
	}
	return count
}
