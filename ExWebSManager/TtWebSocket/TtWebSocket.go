package TtWebSocket

import (
	"fmt"
	"encoding/json"
	"github.com/gorilla/websocket"
	"ttmyth123/kit/ExWebSManager/TtWebSocket/ErrorWeb"
)

type TtWebSocket struct {
	Conn  *websocket.Conn
	SID   string
	isRun bool

	UserInfo interface{}
}

func (this *TtWebSocket) Run(aReadMessage HandlerFunc_ReadMessage) {
	this.isRun = true
	this.run(aReadMessage)
}
func (this *TtWebSocket) run(aReadMessage HandlerFunc_ReadMessage) {
	for {
		messageType, p, err := this.Conn.ReadMessage()
		if err != nil {
			return
		}
		if !this.isRun {
			return
		}
		this.SendPongMsg()

		if aReadMessage != nil {
			aReadMessage(this, messageType, p)
		}
	}
}

func (this *TtWebSocket) SendPingMsg() *ErrorWeb.ErrorWebS {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()
	aWebMsg := WebMsg{T: WebMsg_T_Ping}
	data, _ := json.Marshal(aWebMsg)
	err := this.Conn.WriteMessage(websocket.TextMessage, data)
	if err == nil {
		return nil
	}
	return ErrorWeb.NewErrorWeb(ErrorWeb.EC_WriteMessage, err.Error())
}

func (this *TtWebSocket) SendPongMsg() *ErrorWeb.ErrorWebS {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()
	aWebMsg := WebMsg{T: WebMsg_T_Pong}
	data, _ := json.Marshal(aWebMsg)
	err := this.Conn.WriteMessage(websocket.TextMessage, data)
	if err == nil {
		return nil
	}
	return ErrorWeb.NewErrorWeb(ErrorWeb.EC_WriteMessage, err.Error())
}

func (this *TtWebSocket) SendMsg(aWebMsg WebMsg) *ErrorWeb.ErrorWebS {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()
	data, err := json.Marshal(aWebMsg)
	if err != nil {
		return ErrorWeb.NewErrorWeb(ErrorWeb.EC_Marshal, err.Error())
	}
	err = this.Conn.WriteMessage(websocket.TextMessage, data)
	if err == nil {
		return nil
	}
	return ErrorWeb.NewErrorWeb(ErrorWeb.EC_WriteMessage, err.Error())
}
