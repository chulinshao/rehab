package dao

import (
	"github.com/chulinshao/rehab/models"
	"github.com/go-xorm/xorm"
)

type DoctorDao struct {
	engine *xorm.Engine
}

func NewDoctorDao(engine *xorm.Engine) *DoctorDao {
	return &DoctorDao{
		engine:engine,
	}
}

func (dao DoctorDao) SelectByPrimaryKey(doctorCode string) models.Doctor {
	doctor := models.Doctor{}
	dao.engine.ID(doctorCode).Get(&doctor)
	return doctor
}
