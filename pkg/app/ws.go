package app

import (
	"fmt"
	"github.com/ebar-go/ego"
)

var w ego.WsServer

func init()  {
	fmt.Println("wqewqewq")
	w = ego.WebsocketServer()
}

func Websocket() ego.WsServer {
	return w
}

