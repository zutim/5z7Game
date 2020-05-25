package game

import (
	"5z7Game/msg"
	"5z7Game/pkg/utils"
	"errors"
	"fmt"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ws"
	"log"
	"sync"
)

type Room struct {
	id      int
	Clients map[string]*ws.Connection
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

func (r *Room) determineRole(c *ws.Connection, m *msg.Common) {
	r.mu.Lock()
	for i := range r.Clients {
		if r.battle.Black == "" {
			r.battle.Black = i

			res,err := utils.JsonEncode(&msg.Common{
				Op:      "game_role",
				Args:    "black",
				Msg:     "1",
				MsgType: "",
				FlagId:  0,
			})
			if  err != nil {
				fmt.Println(err)
			}
			r.Clients[i].Send([]byte(res))
			//r.SendToBlack(c, &msg.Common{
			//	Op:      "game_role",
			//	Args:    "black",
			//	Msg:     "1",
			//	MsgType: "",
			//	FlagId:  0,
			//})
		} else {
			r.battle.White = i
			res,err := utils.JsonEncode(&msg.Common{
				Op:      "game_role",
				Args:    "white",
				Msg:     "2",
				MsgType: "",
				FlagId:  0,
			})
			if  err != nil {
				fmt.Println(err)
			}
			r.Clients[i].Send([]byte(res))
			//r.SendToWhite(c, &msg.Common{
			//	Op:      "game_role",
			//	Args:    "white",
			//	Msg:     "2",
			//	MsgType: "",
			//	FlagId:  0,
			//})
		}
	}
	r.mu.Unlock()
}

func (r *Room) UpdateStatus() {
	r.mu.Lock()
	r.status = len(r.Clients)
	r.mu.Unlock()
}

func (r *Room) GameStart(c *ws.Connection, m *msg.Common) {
	if r.status == 2 {
		r.battle = NewBattle(func(in bool) {
			var Msg string
			if in {
				Msg = "white"
			} else {
				Msg = "black"
			}

			res,err := utils.JsonEncode(&msg.Common{
				Op:      "game_over",
				Args:    "winner",
				Msg:     Msg,
				MsgType: "timeout",
				FlagId:  0,
			})
			if  err != nil {
				fmt.Println(err)
			}
			app.WebSocket().Broadcast([]byte(res),nil)
		})
		r.determineRole(c, m)
		res,err := utils.JsonEncode(&msg.Common{
			Op:      "game_start",
			Args:    "black",
			Msg:     "black",
			MsgType: "",
			FlagId:  0,
		})
		if  err != nil {
			fmt.Println(err)
		}
		app.WebSocket().Broadcast([]byte(res),nil)
	} else {
		res,err := utils.JsonEncode(&msg.Common{
			Op:      "game_waiting",
			Args:    "",
			Msg:     fmt.Sprintf("进入房间,id:%d成功，等待对手中", r.id),
			MsgType: "",
			FlagId:  0,
		})
		if  err != nil {
			fmt.Println(err)
		}
		app.WebSocket().Broadcast([]byte(res),nil)
	}
}

func (r *Room) GameOver() {
	r.battle.GameOver()
}

func (r *Room) BlackChessDown(c *ws.Connection, m *msg.Common) {
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
		r.Clients[r.battle.White].Send([]byte(res))
		win := r.battle.CheckWin(msgChess)
		if win {
			log.Println("game_over winner:black")

			res,err := utils.JsonEncode(&msg.Common{
				Op:      "game_over",
				Args:    "winner",
				Msg:     "black",
				MsgType: "",
				FlagId:  0,
			})
			if  err != nil {
				fmt.Println(err)
			}
			app.WebSocket().Broadcast([]byte(res),nil)

			r.GameOver()
			r.Clear()

		}
	}
}

func (r *Room) WhiteChessDown(c *ws.Connection, m *msg.Common) {
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
		r.Clients[r.battle.Black].Send([]byte(res))

		win := r.battle.CheckWin(msgChess)
		if win {
			log.Println("game_over winner:white")

			res,err := utils.JsonEncode(&msg.Common{
				Op:      "game_over",
				Args:    "winner",
				Msg:     "white",
				MsgType: "",
				FlagId:  0,
			})
			if  err != nil {
				fmt.Println(err)
			}
			app.WebSocket().Broadcast([]byte(res),nil)
			r.GameOver()
			r.Clear()
		}
	}
}

//用户进入
func (r *Room) ClientIn(c *ws.Connection) (bool, error) {
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
	r.Clients = make(map[string]*ws.Connection, 2)
	r.mu.Unlock()
}

//func (r *Room) ClientOut(c *ws.Connection, uuid string) []string {
//	r.mu.Lock()
//	if r.battle != nil {
//		//正在对战，对方获胜
//		role, _ := r.RoleIam(uuid)
//		r.GameOver()
//		if role == "white" {
//			r.SendToBlack(c, &msg.Common{
//				Op:      "game_over",
//				Args:    "winner",
//				Msg:     "black",
//				MsgType: "disconnect",
//				FlagId:  0,
//			})
//		}
//		if role == "black" {
//			r.SendToWhite(c, &msg.Common{
//				Op:      "game_over",
//				Args:    "winner",
//				Msg:     "white",
//				MsgType: "disconnect",
//				FlagId:  0,
//			})
//		}
//	}
//	var res []string
//	for v := range r.Clients {
//		res = append(res, v)
//	}
//	r.mu.Unlock()
//	r.Clear()
//	return res
//}

func (r *Room) Clear() {
	r.initClient()
	r.UpdateStatus()
}
