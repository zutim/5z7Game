package game

import (
	"5z7Game/http/handler/msg"
	"5z7Game/pkg/app"
	"5z7Game/pkg/utils"
	"errors"
	"fmt"
	"github.com/ebar-go/ego"
	"log"
	"sync"
)

type Room struct {
	id      int
	Clients map[string]*ego.WebsocketConn
	timer   int //每步20秒
	status  int //0空，1，有1个人，2，有2个人
	battle  *Battle
	mu      sync.Mutex
}



func (r *Room) RoleIam(uuid string) (string, error) {
	if r.battle.Black == uuid {
		return "black", nil
	} else if r.battle.White == uuid {
		return "white", nil
	}
	return "", errors.New("invalid player")
}

//func (r *Room) SendToWhite(c *ws.Connection, m *msg.Common) {
//	c.Send(r.battle.White, m)
//}
//func (r *Room) SendToBlack(c *ws.Connection, m *msg.Common) {
//	c.Send(r.battle.Black, m)
//
//}

func (r *Room) determineRole(c *ego.WebsocketConn, m *msg.Common) {
	r.mu.Lock()
	for i := range r.Clients {
		if r.battle.Black == "" {
			r.battle.Black = i
			res :=msg.ResComm("game_role","black","1","",0)
			app.Websocket().Send([]byte(res),r.Clients[i])
		} else {
			r.battle.White = i
			res :=msg.ResComm("game_role","white","2","",0)
			app.Websocket().Send([]byte(res),r.Clients[i])
		}
	}
	r.mu.Unlock()
}

func (r *Room) UpdateStatus() {
	r.mu.Lock()
	r.status = len(r.Clients)
	r.mu.Unlock()
}

func (r *Room) GameStart(c *ego.WebsocketConn, m *msg.Common) {
	if r.status == 2 {
		r.battle = NewBattle(func(in bool) {
			var Msg string
			if in {
				Msg = "white"
			} else {
				Msg = "black"
			}
			res := msg.ResComm("game_over","winner",Msg,"timeout",0)
			app.Websocket().Broadcast([]byte(res),nil)
		})
		r.determineRole(c, m)
		res := msg.ResComm("game_start","black","black","",0)
		app.Websocket().Broadcast([]byte(res),nil)
	} else {
		res := msg.ResComm("game_waiting","",fmt.Sprintf("进入房间,id:%d成功，等待对手中", r.id),"",0)
		app.Websocket().Broadcast([]byte(res),nil)
	}
}

func (r *Room) GameOver() {
	r.battle.GameOver()
}

func (r *Room) BlackChessDown(c *ego.WebsocketConn, m *msg.Common) {
	var msgChess *msg.ChessDown
	if err := utils.JsonDecode([]byte(m.Msg),&msgChess); err != nil {
		fmt.Println(err)
	}
	err := r.battle.BlackChessDown(msgChess)
	if err != nil {
		log.Println("BlackChessDown_err:", err)
	} else {
		res,err := utils.JsonEncode(&m)
		if  err != nil {
			fmt.Println(err)
		}
		//r.Clients[r.battle.White].Send([]byte(res))
		app.Websocket().Send([]byte(res),r.Clients[r.battle.White])
		win := r.battle.CheckWin(msgChess)
		if win {
			log.Println("game_over winner:black")

			res := msg.ResComm("game_over","winner","black","",0)
			app.Websocket().Broadcast([]byte(res),nil)

			r.GameOver()
			r.Clear()

		}
	}
}

func (r *Room) WhiteChessDown(c *ego.WebsocketConn, m *msg.Common) {
	var msgChess *msg.ChessDown
	if err := utils.JsonDecode([]byte(m.Msg),&msgChess); err != nil {
		fmt.Println(err)
	}

	err := r.battle.WhiteChessDown(msgChess)
	if err != nil {
		log.Println("WhiteChessDown_err:", err)
	} else {
		res,err := utils.JsonEncode(&m)
		if  err != nil {
			fmt.Println(err)
		}
		//r.Clients[r.battle.Black].Send([]byte(res))
		app.Websocket().Send([]byte(res),r.Clients[r.battle.Black])
		win := r.battle.CheckWin(msgChess)
		if win {
			log.Println("game_over winner:white")
			res := msg.ResComm("game_over","winner","white","",0)
			app.Websocket().Broadcast([]byte(res),nil)
			r.GameOver()
			r.Clear()
		}
	}
}

//用户进入
func (r *Room) ClientIn(c *ego.WebsocketConn) (bool, error) {
	ok := false
	if r.status > 2 {
		return false, errors.New("this room is full")
	}

	r.mu.Lock()
	r.Clients[c.ID] = c
	r.mu.Unlock()
	r.UpdateStatus()
	if r.status == 2 {
		ok = true
	}
	return ok, nil
}

func (r *Room) initClient() {
	r.mu.Lock()
	r.Clients = make(map[string]*ego.WebsocketConn, 2)
	r.mu.Unlock()
}

func (r *Room) ClientOut(c *ego.WebsocketConn, uuid string) []string {
	r.mu.Lock()
	if r.battle != nil {
		//正在对战，对方获胜
		role, _ := r.RoleIam(uuid)
		r.GameOver()
		if role == "white" {
			res := msg.ResComm("game_over","winner","black","disconnect",0)
			app.Websocket().Send([]byte(res),c)
		}
		if role == "black" {
			res := msg.ResComm("game_over","winner","white","disconnect",0)
			app.Websocket().Send([]byte(res),c)
		}
	}
	var res []string
	for v := range r.Clients {
		res = append(res, v)
	}
	r.mu.Unlock()
	r.Clear()
	return res
}

func (r *Room) Clear() {
	r.initClient()
	r.UpdateStatus()
}
