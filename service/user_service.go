package service

import (
	"github.com/chulinshao/rehab/dao"
	"github.com/chulinshao/rehab/datasource"
	"github.com/chulinshao/rehab/models"
)

type UserService interface {
	ListAll() []models.User

	FindById(id string) models.User
}

func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.GetEngine()),
	}
}

type userService struct {
	dao *dao.UserDao
}

func (s userService) ListAll() []models.User {
	return s.dao.GetAll()
}

func (s userService) FindById(id string) models.User {
	return s.dao.FindByID(id)
}
