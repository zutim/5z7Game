package app

import "github.com/ebar-go/ws"

var w ws.Manager

func init()  {
	w = ws.New()
}

func Websocket() ws.Manager {
	return w
}

