package helper

import (
	"5z7Game/pkg/service/data"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/errors"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/gin-gonic/gin"
)

// GeLoginUserFromContext 通过Context获取登录的用户信息
func GeLoginUserFromContext(ctx *gin.Context) data.User {
	claims , _:= ctx.Get(app.Jwt().ClaimsKey)
	if claims == nil {
		secure.Panic(errors.Unauthorized("please login first"))
	}

	userClaims, ok := claims.(*data.UserClaims)
	if !ok {
		secure.Panic(errors.Unauthorized("please login first"))
	}

	return userClaims.User
}
