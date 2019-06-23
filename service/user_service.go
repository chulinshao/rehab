package service

type UserService interface {
	GetAll() string
}

func NewUserService() UserService {
	return &userService{}
}

type userService struct{}

func (s userService) GetAll() string {
	return "test"
}
