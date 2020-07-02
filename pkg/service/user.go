package service

import (
	"5z7Game/pkg/app"
	"5z7Game/pkg/dto/request"
	"5z7Game/pkg/model/dao"
)

type UserService struct {}

// User 用户的Service
func User() *UserService {
	return &UserService{}
}

// Auth 校验用户
func (service *UserService) Auth(req request.UserAuthRequest) {
	userDao := dao.User(app.DB())

	if err := userDao.Auth(req); err != nil {
		// TODO
	}

	// TODO
}


func (service *UserService) Login(req request.UserAuthRequest){
	userDao := dao.User(app.DB())

	if err := userDao.Login(req); err != nil {
		//auth.Jwt()

	}
}
