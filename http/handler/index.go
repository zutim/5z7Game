package handler

import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context)  {
	ctx.HTML(200,"index.html",nil)
	//response.WrapContext(ctx).Success("hello,5z7")
}
