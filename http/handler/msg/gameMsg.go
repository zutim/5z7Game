package msg


type NewMsg struct {
	msg  *Common
	head bool
}

func NewGameMsg() NewMsg {
	return NewMsg{}
}

func (n *NewMsg) init() *NewMsg {
	n.msg = &Common{
		Op:      "",
		Args:    "",
		Msg:     "",
		MsgType: "",
		FlagId:  0,
	}
	n.head = false
	return n
}

func (n *NewMsg) CustomOp(key string) *NewMsg {
	n.msg.Op = "game_" + key
	n.head = true
	return n
}

func (n *NewMsg) CommonOp() *NewMsg {
	return n.CustomOp("common")
}
func (n *NewMsg) ErrorOp() *NewMsg {
	return n.CustomOp("error")
}
func (n *NewMsg) SetMsg(v string) *NewMsg {
	n.msg.Msg = v
	return n
}
func (n *NewMsg) hasOp() bool {
	if n.head {
		return true
	}
	return false
}
func (n *NewMsg) Done() *Common {
	if n.hasOp() {
		return n.msg
	}
	return nil
}
