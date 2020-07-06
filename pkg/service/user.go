// service 服务模块
package service

import (
	"5z7Game/pkg/dto/request"
	"5z7Game/pkg/dto/response"
	"5z7Game/pkg/enum/statusCode"
	"5z7Game/pkg/service/dao"
	"5z7Game/pkg/service/data"
	"5z7Game/pkg/service/entity"
	"fmt"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/utils/date"
	"github.com/ebar-go/ego/utils/encrypt"
	"github.com/jinzhu/gorm"
)

type userService struct {
}

// User 用户服务
func User() *userService {
	return &userService{}
}

// Auth 校验用户登录
func (service *userService) Auth(req request.UserAuthRequest) (*response.UserAuthResponse, error) {
	// 调用dao的方法，根据邮箱获取用户信息
	user, err := dao.User(app.DB()).GetByUsername(req.Username)
	if err != nil {
		// 查询数据库失败时,返回自定义的error类型，用于panic拦截输出
		return nil, errors.New(statusCode.DataNotFound, fmt.Sprintf("获取用户信息失败:%v", err))
	}

	// 校验密码
	if encrypt.Md5(req.Password) != user.Password {
		return nil, errors.New(statusCode.PasswordWrong, "密码错误")
	}

	// 组装结构体
	res := new(response.UserAuthResponse)
	// 生成token
	userClaims := new(data.UserClaims)
	userClaims.ExpiresAt = date.GetTimeStamp() + 3600
	userClaims.User.Id = user.ID
	userClaims.User.Email = user.Username
	token, err := app.Jwt().GenerateToken(userClaims)

	if err != nil {
		return nil, errors.New(statusCode.TokenGenerateFailed, err.Error())
	}
	res.Token = token

	return res, nil
}

// Register 注册
func (service *userService) Register(req request.UserRegisterRequest) error {
	db := app.DB()
	db.LogMode(true)
	userDao := dao.User(db)
	// 根据邮箱获取用户信息
	user, err := userDao.GetByUsername(req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New(statusCode.DatabaseQueryFailed, fmt.Sprintf("获取用户信息失败:%v", err))
	}

	// 用户已存在
	if user != nil {
		return errors.New(statusCode.EmailRegistered, "该邮箱已被注册")
	}

	now := int(date.GetTime().Unix())

	user = new(entity.UserEntity)
	user.Username = req.Username
	user.Password = encrypt.Md5(req.Password)
	user.CreatedAt = now
	user.UpdatedAt = now

	if err := userDao.Create(user); err != nil {
		return errors.New(statusCode.DatabaseSaveFailed, fmt.Sprintf("保存数据失败:%v", err))
	}

	return nil
}
