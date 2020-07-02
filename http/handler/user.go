package handler

import (
	"5z7Game/http/response"
	"5z7Game/pkg/dto/request"
	"5z7Game/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// UserAuthHandler 用户登录
func UserAuthHandler(ctx *gin.Context)  {
	req := request.UserAuthRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err!=nil{
		fmt.Println("错误了")
		fmt.Println(err)
	}

	service.User().Login(req)


}

// UserRegisterHandler 用户注册
func UserRegisterHandler(ctx *gin.Context)  {
	// TODO
	response.WrapContext(ctx).Success(nil)
}