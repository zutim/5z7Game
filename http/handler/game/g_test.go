package game

import (
	"fmt"
	"math/rand"
	"testing"
)

func randInt() int {
	return rand.Intn(16)
}

func randInt2() int {
	return rand.Intn(2)
}

func TestXxx(t *testing.T) {

	x := Rooms{}
	x.init()
	fmt.Println(len(x.list))
	x.addList()
	fmt.Println(len(x.list))

}
