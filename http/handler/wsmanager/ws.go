package wsmanager

import (
	"5z7Game/http/handler/game"
	"5z7Game/http/handler/msg"
	"5z7Game/pkg/app"
	"5z7Game/pkg/utils"
	"fmt"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/gin-gonic/gin"
)

func WebsocketHandler(ctx *gin.Context)  {
	conn, err := app.Websocket().UpgradeConn(ctx.Writer, ctx.Request)
	if err != nil {
		secure.Panic(err)
	}

	app.Websocket().Register(conn, func(message []byte){
		var msgs *msg.Common
		if err := utils.JsonDecode(message,&msgs); err != nil {
			fmt.Println(err)
		}
		//在线广播
		Broad()
		//交由Gm管理
		game.Gm.Handler(conn,msgs)
	})
}

func Broad()  {
	online := app.Websocket().GetOnline()
	res := msg.ResComm("server_notice","",fmt.Sprintf("连接成功，当前在线人数：%v", online),"string",0)
	app.Websocket().Broadcast([]byte(res),nil)
}
