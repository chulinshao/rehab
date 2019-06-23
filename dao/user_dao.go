package dao

import "github.com/go-xorm/xorm"

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

func (dao UserDao) GetAll() string {
	return "test123"
}
