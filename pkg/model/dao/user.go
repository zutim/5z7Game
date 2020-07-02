package dao

import (
	"5z7Game/pkg/dto/request"
	"5z7Game/pkg/model/entity"
	"fmt"
	"github.com/jinzhu/gorm"
	"errors"
)

// UserDao
type UserDao struct {
	BaseDao
}

// User return UserDao pointer with db connection
func User(db *gorm.DB) *UserDao {
	dao := &UserDao{}
	dao.db = db
	return dao
}

// Auth 校验用户
func (dao *UserDao) Auth(req request.UserAuthRequest) error {
	var user entity.User
	dao.db.Where("username=?",req.Username).First(&user)
	if user.Username==req.Username{
		fmt.Println(user)
		return nil
	}
	return nil
}

// Register 注册用户
func (dao *UserDao) Register(req request.UserRegisterRequest) error {
	return nil
}

// Register 登录用户
func (dao *UserDao) Login(req request.UserAuthRequest) error {
	var user entity.User
	dao.db.Where("username=?",req.Username).First(&user)
	if user.Password==req.Password{
		return nil
	}
	return errors.New("登录失败")
}
