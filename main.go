package main

import (
	"5z7Game/config"
	"5z7Game/http/route"
	"fmt"
	"github.com/ebar-go/ego"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/component/event"
	"github.com/ebar-go/ego/utils/secure"
)

func init()  {
	// 加载配置
	secure.Panic(config.ReadFromFile("app.yaml"))

	// 初始化数据库
	secure.Panic(app.InitDB())

	// 支持停止http服务时的回调
	event.Listen(event.BeforeHttpShutdown, func(ev event.Event) {
		// 关闭数据库
		fmt.Println("close database")
		_ = app.DB().Close()
	})

}

func main() {
	s := ego.HttpServer()

	route.Load(s.Router)

	secure.Panic(s.Start())
}
