package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"index.html",nil)
	//response.WrapContext(ctx).Success("hello,5z7")
}

func UserHandler(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"login.html",nil)
}
