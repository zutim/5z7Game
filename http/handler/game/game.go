package game

import (
	"5z7Game/http/handler/msg"
	"5z7Game/pkg/app"
	"github.com/ebar-go/ego"
	"sync"
)

var Gm *Game
var Player Players

type Game struct {
	r Rooms
}

func init() {
	var House = Rooms{}
	House.init()
	Gm = &Game{
		r: House,
	}
	Player = Players{}
	Player.List = new(sync.Map)
}

func (this *Game) Handler(client *ego.WebsocketConn, reply *msg.Common) {
	s := reply.Op
	switch s {
	case "game_initClientStatus": //用户初始化连接
		room, _ := this.r.inRoom(client)  //进入房间
		Player.Add(client.ID, room)  //开始进入房间的玩家
		this.r.list[room].GameStart(client, reply) //等待开始
	case "game_chessDown":
		//验证房间
		roomid := Player.Get(client.ID)
		if roomid == -1 {
			res := msg.ResComm("game_error","","找不到房间","string",0)
			app.Websocket().Send([]byte(res),client)
			return
		}

		//验证角色
		r, err := this.r.list[roomid].RoleIam(client.ID)
		if err!=nil{
			res := msg.ResComm("game_error","","房间里找不到角色","string",0)
			app.Websocket().Send([]byte(res),client)
			return
		}

		//执行下棋
		if r == "white" {
			this.r.list[roomid].WhiteChessDown(client, reply)
		} else {
			this.r.list[roomid].BlackChessDown(client, reply)
		}
	case "game_disconnect": //下棋结束，退出连接
		roomid := Player.Get(client.ID)
		if roomid == -1 {
			res := msg.ResComm("game_error","","找不到房间","string",0)
			app.Websocket().Send([]byte(res),client)
			return
		}
		out := this.r.list[roomid].ClientOut(client, client.ID)
		for _, v := range out {
			Player.Del(v)
		}
	}
}







