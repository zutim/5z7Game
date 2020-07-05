package route

import (
	"5z7Game/http/handler"
	wshandler "5z7Game/http/handler/wsmanager"
	"5z7Game/http/middleware"
	"5z7Game/pkg/app"
	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine)  {
	router.Use(middleware.Recover)

	router.LoadHTMLFiles("./html/view/index.html", "./html/view/login.html")
	//加载静态资源，例如网页的css、js
	router.Static("html/static", "./html/static")

	//加载静态资源，一般是上传的资源，例如用户上传的图片
	//router.StaticFS("/upload", http.Dir("upload"))

	//加载单个静态文件
	//router.StaticFile("/favicon.ico", "./static/favicon.ico")

	//router.LoadHTMLGlob("./html/index.html")
	router.GET("/", handler.IndexHandler)

	//router.LoadHTMLGlob("./html/login.html")
	router.GET("/user",handler.UserHandler)

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
	ws :=app.Websocket()

	router.GET("/ws", wshandler.WebsocketHandler)

	go ws.Start()

	// 不需要校验token的路由
	public := router.Group("v1/public")
	{
		// 用户登录
		public.POST("user/auth", handler.UserAuthHandler)

		// 用户注册
		public.POST("user/register", handler.UserRegisterHandler)
	}
}
