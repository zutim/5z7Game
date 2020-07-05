package msg

import (
	"5z7Game/pkg/utils"
	"fmt"
)

type Common struct {
	Op      string `json:"op"`
	Args    string `json:"args"`
	Msg     string `json:"msg"`
	MsgType string `json:"msgType"`
	FlagId  int    `json:"flagId"`
}

var comm *Common

func init() {
	comm =  &Common{
		Op:      "",
		Args:    "",
		Msg:     "",
		MsgType: "",
		FlagId:  0,
	}
}

func ResComm(op string,arg string,msg string,msgtype string,flag int)(string){
	res,err:= utils.JsonEncode(&Common{
		Op:      op,
		Args:    arg,
		Msg:     msg,
		MsgType: msgtype,
		FlagId:  flag,
	})
	if err !=nil{
		fmt.Print(err)
	}
	return res
}

type ChessDown struct {
	I int `json:"i"`
	J int `json:"j"`
}