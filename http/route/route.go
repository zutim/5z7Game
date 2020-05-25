package route

import (
	"5z7Game/http/handler"
	"5z7Game/http/middleware"
	"github.com/ebar-go/ego/app"
	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine)  {
	router.Use(middleware.Recover)

	router.LoadHTMLGlob("./html/index.html")
	router.GET("/", handler.IndexHandler)

	// 定义需要token校验的路由
	auth := router.Group("v1/home").Use(middleware.JWT)
	{
		// 获取用户信息
		auth.GET("user/info", nil)

		// 获取用户历史棋局分页
		auth.GET("user/chess", nil)

		// 获取历史棋局详情
		auth.GET("user/chess/:id", nil)

		// 创建一局游戏
		auth.POST("user/chess", nil)

	}

	// websocket
	router.GET("v1/ws", handler.WebsocketHandler)

	go app.WebSocket().Start()

	// 不需要校验token的路由
	public := router.Group("v1/public")
	{
		// 用户登录
		public.POST("user/auth", handler.UserAuthHandler)

		// 用户注册
		public.POST("user/register", handler.UserRegisterHandler)
	}
}
