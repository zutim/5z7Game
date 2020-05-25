package msg


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

type ChessDown struct {
	I int `json:"i"`
	J int `json:"j"`
}