package game

import (
	"5z7Game/pkg/utils"
	"fmt"
	"github.com/ebar-go/ws"
	"5z7Game/msg"
	"sync"
)

type Game struct {
	r Rooms
}

func (this *Game) Handler(client *ws.Connection, reply *msg.Common) {
	s := reply.Op
	switch s {
		case "game_initClientStatus": //用户初始化连接
			room, _ := this.r.inRoom(client)  //进入房间
			Player.Add(client.ID, room)  //开始进入房间的玩家
			this.r.list[room].GameStart(client, reply) //等待开始
		case "game_chessDown":
			roomid := Player.Get(client.ID)
			if roomid == -1 {
				res,err := utils.JsonEncode(&msg.Common{
					Op:      "game_error",
					Args:    "",
					Msg:     "invalid roomid",
					MsgType: "string",
					FlagId:  0,
				})
				if  err != nil {
					fmt.Println(err)
				}
				client.Send([]byte(res))
				return
			}

			r, err := this.r.list[roomid].RoleIam(client.ID)
			if err!=nil{
				res,err := utils.JsonEncode(&msg.Common{
					Op:      "game_error",
					Args:    "",
					Msg:     "invalid role in this room",
					MsgType: "string",
					FlagId:  0,
				})
				if  err != nil {
					fmt.Println(err)
				}
				client.Send([]byte(res))
				return
			}

			if r == "white" {
				this.r.list[roomid].WhiteChessDown(client, reply)
			} else {
				this.r.list[roomid].BlackChessDown(client, reply)
			}
		case "game_disconnect":
			roomid := Player.Get(client.ID)
			if roomid == -1 {
				//manager.Send(uuid, &msg.Common{
				//	Op:      "game_error",
				//	Args:    "",
				//	Msg:     "invalid roomid",
				//	MsgType: "string",
				//	FlagId:  0,
				//})
				return
			}
			//out := this.r.list[roomid].ClientOut(client, client.ID)
			//for _, v := range out {
			//	Player.Del(v)
			//}
	}
}


type Rooms struct {
	list []*Room
	mu   sync.RWMutex
}

func (w *Rooms) init() {
	w.list = make([]*Room, 8)
	for i := range w.list {
		w.list[i] = &Room{
			id:      i,
			Clients: make(map[string]*ws.Connection, 2),
		}
	}
}

func (w *Rooms) GetById(id int) *Room {
	return w.list[id]
}

func (w *Rooms) addRooms() {

	w.list = append(w.list, &Room{id: len(w.list), Clients: make(map[string]*ws.Connection, 2)})

}

func (w *Rooms) WaitingRoom() int {
	w.mu.RLock()
	for _, v := range w.list {
		if v.status == 1 {
			w.mu.RUnlock()
			return v.id
		}
	}
	return -1
}

func (w *Rooms) FirstEmptyRoom() int {
	w.mu.RLock()
	for _, v := range w.list {
		if v.status == 0 {
			w.mu.RUnlock()
			return v.id
		}
	}
	return -1
}

func (w *Rooms) inRoom(c *ws.Connection) (roomid int, err error) {
	waitingRoom := w.WaitingRoom()
	if waitingRoom == -1 {
		//没有等待的房间
		emptyRoom := w.FirstEmptyRoom()
		if emptyRoom == -1 {
			//人居然满了我是不信的
			//添加房间
			w.addRooms()
			return w.inRoom(c)
		} else {
			_, err = w.list[emptyRoom].ClientIn(c)
		}
		roomid = emptyRoom
	} else {
		//开始对战
		_, err = w.list[waitingRoom].ClientIn(c)
		roomid = waitingRoom

	}
	return roomid, err
}



var Gm *Game
var Player Players

func init() {
	var House = Rooms{}
	House.init()
	Gm = &Game{
		r: House,
	}
	Player = Players{}
	Player.List = new(sync.Map)
}
