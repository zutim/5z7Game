package game

import (
	"github.com/ebar-go/ego"
	"sync"
)

type Rooms struct {
	list []*Room
	mu   sync.RWMutex
}

func (w *Rooms) init() {
	w.list = make([]*Room, 8)
	for i := range w.list {
		w.list[i] = &Room{
			id:      i,
			Clients: make(map[string]*ego.WebsocketConn, 2),
		}
	}
}

func (w *Rooms) GetById(id int) *Room {
	return w.list[id]
}

func (w *Rooms) addRooms() {

	w.list = append(w.list, &Room{id: len(w.list), Clients: make(map[string]*ego.WebsocketConn, 2)})

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

func (w *Rooms) inRoom(c *ego.WebsocketConn) (roomid int, err error) {
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

