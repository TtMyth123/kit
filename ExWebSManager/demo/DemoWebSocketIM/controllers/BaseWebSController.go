package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"ttmyth123/kit/WebSManager"
	"ttmyth123/kit/WebSManager/TtWebSocket"
	"ttmyth123/kit/WebSManager/demo/DemoWebSocketIM/cacheWebsocet"
)

type BaseWebSI interface {
	GetRoom() *WebSManager.WebsocketManager
	GetUID() string
	GetSID() string
	GetSessionSID() string
}

type BaseWebSController struct {
	beego.Controller // Embed struct that has stub implementation of the interface.
	BaseWebSI
}

func (this *BaseWebSController) upgradeWebsocket(w http.ResponseWriter, r *http.Request, responseHeader http.Header, readBufSize, writeBufSize int) (*websocket.Conn, error) {
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return ws, err
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return ws, err
	}
	this.GetRoom().SendRedirectMsg(ws)
	return ws, err
}

func (this *BaseWebSController) NewJoin(aReadMessage TtWebSocket.HandlerFunc_ReadMessage, aBroadcastMsg *TtWebSocket.WebMsg) {
	uname := this.GetUID()
	sid := this.GetSID()
	if len(uname) == 0 {
		_, err := this.upgradeWebsocket(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
		if err != nil {
			return
		}
	}

	SessionSID := this.GetSessionSID()
	if SessionSID == "" || SessionSID != sid {
		_, err := this.upgradeWebsocket(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
		if err != nil {
			return
		}
	}

	//判断当前用户是不是已在线上。
	curSID := cacheWebsocet.GetRoom().GetCacheConnSID(uname)
	fmt.Println("curSID:", curSID, "SID:", sid)
	if curSID != sid {
		this.GetRoom().CloseSID(curSID)
	}

	ws, err := this.upgradeWebsocket(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		return
	}

	aWebsocket := this.GetRoom().AddConn(uname, ws, sid, nil)
	if aBroadcastMsg != nil {
		this.GetRoom().BroadcastMsg(*aBroadcastMsg)
	}
	aWebsocket.Run(aReadMessage)
}
