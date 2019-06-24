package service

import (
	"github.com/chulinshao/rehab/dao"
	"github.com/chulinshao/rehab/datasource"
)

type UserService interface {
	ListAll() string
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.GetEngine()),
	}
}

type userService struct {
	dao *dao.UserDao
}

func (s userService) ListAll() string {
	return s.dao.GetAll()
}
