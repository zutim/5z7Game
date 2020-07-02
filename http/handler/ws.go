package handler

import (
	"5z7Game/game"
	"5z7Game/msg"
	"5z7Game/pkg/utils"
	"fmt"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/ebar-go/ws"
	"github.com/gin-gonic/gin"
	"net/http"
)



func WebsocketHandler(ctx *gin.Context)  {
	conn, err := ws.UpgradeConn(ctx.Writer, ctx.Request)
	if err != nil {
		secure.Panic(err)
	}

	ws.Register(conn, func(message []byte){
		if string(message) == "broadcast" {// 广播
			ws.Broadcast([]byte("hello,welcome"), nil)
			return
		}
		ws.Send(message, conn) // 单对单发送

	})


	conn , err := ws.GetUpgradeConnection(ctx.Writer, ctx.Request)
	if err != nil {
		http.NotFound(ctx.Writer, ctx.Request)
		return
	}

	client := ws.NewConnection(conn, func(ctx *ws.Context,c *ws.Connection) string {
		var message *msg.Common
		if err := utils.JsonDecode([]byte(ctx.GetMessage()),&message); err != nil {
			fmt.Println(err)
		}
		//在线广播
		Broad()
		//交由Gm管理
		game.Gm.Handler(c,message)
		//return ctx.GetMessage()
		return ""
	})

	app.WebSocket().Register(client)

	go client.Listen()
}

func Broad()  {
	online := app.WebSocket().GetOnLine()
	res,err := utils.JsonEncode(&msg.Common{
		Op:      "server_notice",
		Args:    "",
		Msg:     fmt.Sprintf("连接成功，当前在线人数：%v", online),
		MsgType: "string",
		FlagId:  0,
	})
	if  err != nil {
		fmt.Println(err)
	}
	app.WebSocket().Broadcast([]byte(res),nil)
}
