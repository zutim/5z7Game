package app

import (
	"github.com/ebar-go/ego"
)

var ws ego.WsServer

func init()  {
	ws= ego.WebsocketServer()
}

func Websocket() ego.WsServer {
	return ws
}

