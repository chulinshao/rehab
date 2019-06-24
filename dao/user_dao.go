package dao

import (
	"github.com/chulinshao/rehab/models"
	"github.com/go-xorm/xorm"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

func (dao UserDao) GetAll() []models.User {
	user := make([]models.User, 0)
	dao.engine.Find(&user)
	return user
}

func (dao UserDao) FindByID(id string) models.User {
	user := models.User{}
	dao.engine.Id(id).Get(&user)
	return user
}
