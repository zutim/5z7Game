package game

import (
	"5z7Game/msg"
	"errors"
	"log"
	"sync"
	"time"
)

type Battle struct {
	Board       [16][16]int
	Black       string //uuid
	White       string //uuid
	NowRole     int    //default 1 现在轮到谁落子
	NowRoleName string
	Timer       <-chan time.Time
	TimerEvt    func()
	TimerEvtIng bool
	TurnTime    int
	End         chan int
	mu          *sync.Mutex
}

func (v *Battle) NewBoard() {
	for i := 0; i < 16; i++ {
		v.Board[i] = [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	}
}
func (v *Battle) CheckWin(m *msg.ChessDown) bool {
	i, j := m.I, m.J
	win := 0
	//当前是否赢
	var role int
	if v.NowRole == 1 {
		role = 2
	} else {
		role = 1
	}

	//竖向检测
	for _, x := range v.Board {

		if x[j] == role {
			win++
		} else {
			win = 0
		}
		if win >= 5 {
			log.Println("win by |")
			return true
		}
	}
	win = 0
	//横向检测
	for _, x := range v.Board[i] {
		if x == role {
			win++
		} else {
			win = 0
		}
		if win >= 5 {
			log.Println("win by -")
			return true
		}
	}
	win = 0
	//\向检测
	var tmp [2]int
	start := 0
	if i > j {
		tmp = [2]int{i - j, 0}
	}
	if i <= j {
		tmp = [2]int{0, j - i}
	}
	for ; start < 15; start++ {
		tmpI := tmp[0] + start
		tmpJ := tmp[1] + start
		if tmpI > 15 || tmpJ > 15 {
			break
		}
		if v.Board[tmpI][tmpJ] == role {
			win++
		} else {
			win = 0
		}
		if win > 4 {
			log.Println("win by \\")
			return true
		}
	}
	win = 0
	// /向检测
	start = 0
	if i >= 15-j {
		tmp = [2]int{15, j + i - 15} //i 14,j 2
	}
	if i < 15-j {
		tmp = [2]int{i + j, 0} //i 10,j 1
	}
	for ; start < 15; start++ {
		tmpI := tmp[0] - start
		tmpJ := tmp[1] + start
		if tmpJ > 15 || tmpI < 0 {
			break
		}
		if v.Board[tmpI][tmpJ] == role {
			win++
		} else {
			win = 0
		}
		if win > 4 {
			log.Println("win by /")
			return true
		}
	}
	return false
}

func (v *Battle) chessDown(in *msg.ChessDown, role int) error {
	v.mu.Lock()
	if v.Board[in.I][in.J] == 0 && v.NowRole == role {
		v.Board[in.I][in.J] = role
		v.nextRole()
		v.resetTimer()
	} else {
		v.mu.Unlock()
		return errors.New("not your turn or invalid position")
	}
	v.mu.Unlock()
	return nil
}

func (v *Battle) BlackChessDown(in *msg.ChessDown) error {
	return v.chessDown(in, 1)
}
func (v *Battle) WhiteChessDown(in *msg.ChessDown) error {
	return v.chessDown(in, 2)
}
func (v *Battle) nextRole() {
	if v.NowRole == 1 {
		v.NowRole = 2
		v.NowRoleName = "white"
	} else {
		v.NowRole = 1
		v.NowRoleName = "black"
	}
}

func (v *Battle) Turn() int {
	return v.NowRole
}
func (v *Battle) resetTimer() {
	v.TurnTime = 15
}

func (v *Battle) GameOver() {
	v.End = make(chan int, 1)
	v.End <- 1

}

func (v *Battle) WhoAmI(uuid string) string {
	if v.Black == uuid {
		return "black"
	}
	if v.White == uuid {
		return "white"
	}
	return ""
}

func (v *Battle) SetTimerFunc(xx func(in bool)) {
	if !v.TimerEvtIng {
		v.TimerEvt = func() {
			go func() {
				var ticker = time.NewTicker(time.Second * 1)
				for {
					select {
					case <-ticker.C:
						v.TurnTime--
						if v.TurnTime < 0 {
							xx(v.NowRole == 1)
							return
						}
					case <-v.End:
						return
					}
				}
			}()
		}
		v.TimerEvtIng = true
		v.TimerEvt()
	}
}

func NewBattle(timerFunc func(in bool)) *Battle {
	b := &Battle{
		Board:       [16][16]int{},
		NowRole:     1,
		NowRoleName: "black",
		mu:          new(sync.Mutex),
	}
	b.NewBoard()
	b.resetTimer()
	b.SetTimerFunc(timerFunc)
	return b
}

func init() {

}
