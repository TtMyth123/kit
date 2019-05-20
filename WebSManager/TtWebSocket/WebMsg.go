package TtWebSocket

const (
	WebMsg_T_Ping  = 0  //ping 消息
	WebMsg_T_Pong  = 1  //pong 消息
	WebMsg_T_Jion  = 10 //加入消息
	WebMsg_T_leave = 11 //离开消息

	WebMsg_T_redirect = 302  //跳转
	WebMsg_T_Chat     = 1000 //聊天
)

type WebMsg struct {
	T   int
	C   string
	SID string

	Obj interface{}
}
