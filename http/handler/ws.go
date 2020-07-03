package handler

import (
	"5z7Game/game"
	"5z7Game/msg"
	"5z7Game/pkg/app"
	"5z7Game/pkg/utils"
	"fmt"
	"github.com/ebar-go/ego"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/gin-gonic/gin"
)

func WebsocketHandler(ctx *gin.Context)  {
	ws := app.Websocket()
	conn, err := ws.UpgradeConn(ctx.Writer, ctx.Request)
	if err != nil {
		secure.Panic(err)
	}

	ws.Register(conn, func(message []byte){
		var msgs *msg.Common
		if err := utils.JsonDecode(message,&msgs); err != nil {
			fmt.Println(err)
		}
		//在线广播
		Broad(ws)
		//交由Gm管理
		game.Gm.Handler(conn,msgs)
		//if string(message) == "broadcast" {// 广播
		//	ws.Broadcast([]byte("hello,welcome"), nil)
		//	return
		//}
		//ws.Send(message, conn) // 单对单发送

	})
}

func Broad(ws ego.WsServer)  {
	online := ws.GetOnline()
	res := msg.ResComm("server_notice","",fmt.Sprintf("连接成功，当前在线人数：%v", online),"string",0)
	ws.Broadcast([]byte(res),nil)
}
