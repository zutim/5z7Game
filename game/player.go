package game

import (
	"log"
	"sync"
)

type Players struct {
	List *sync.Map
}

func (p *Players) Add(uuid string, v int) {
	p.List.Store(uuid, v)
}

func (p *Players) Get(uuid string) int {
	load, ok := p.List.Load(uuid)
	if !ok {
		log.Println("PlayerList load error,uuid:", uuid)
	}
	if load == nil {
		return -1
	}
	return load.(int)
}

func (p *Players) Del(uuid string) {
	p.List.Delete(uuid)
}
