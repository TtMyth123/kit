// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/i18n"
	"net/http"
	"ttmyth123/kit/WebSManager"
	"ttmyth123/kit/WebSManager/TtWebSocket"
	"ttmyth123/kit/WebSManager/demo/DemoWebSocketIM/cacheWebsocet"
	"ttmyth123/kit/WebSManager/demo/DemoWebSocketIM/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	baseController
	//BaseWebSController
	//i18n.Locale
}

func (this *WebSocketController) GetRoom() *WebSManager.WebsocketManager {
	return cacheWebsocet.GetRoom()
}

func (this *WebSocketController) GetUID() string {
	return this.GetString("uname")
}

func (this *WebSocketController) GetSID() string {
	return this.GetString("sid")
}

func (this *WebSocketController) GetSessionSID() string {
	s := this.GetSession("SessionUser")
	if s != nil {
		userInfo := s.(models.BaseUserInfo)
		return userInfo.SID
	}
	return ""
}

func (this *WebSocketController) Prepare() {
	// Reset language option.
	this.Lang = "" // This field is from i18n.Locale.

	// 1. Get language information from 'Accept-Language'.
	al := this.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		if i18n.IsExist(al) {
			this.Lang = al
		}
	}

	// 2. Default language is English.
	if len(this.Lang) == 0 {
		this.Lang = "en-US"
	}

	// Set template level language option.
	this.Data["Lang"] = this.Lang
}

func (this *WebSocketController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

// Get method handles GET requests for WebSocketController.
func (this *WebSocketController) Get() {
	// Safe check.
	uname := this.GetString("uname")
	sid := this.GetString("sid")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}
	s := this.GetSession("SessionUser")
	if s != nil {
		userInfo := s.(models.BaseUserInfo)
		if sid != userInfo.SID {
			this.redirect("/")
		}

	} else {
		this.redirect("/")
	}

	cachetSID := cacheWebsocet.GetRoom().GetCacheConnSID(uname)
	if cachetSID != "" && cachetSID != sid {
		//如果 多重登录。给之前的连接发信息
		aWebMsg := TtWebSocket.WebMsg{T: 0, C: ""}
		cacheWebsocet.GetRoom().SendMsg(cachetSID, aWebMsg)
	}

	this.TplName = "websocket.html"

	this.Data["SID"] = sid
	this.Data["IsWebSocket"] = true
	this.Data["UserName"] = uname
}

// Join method handles WebSocket requests for WebSocketController.
func (this *WebSocketController) Join() {
	defer this.StopRun()
	uname := this.GetString("uname")
	sid := this.GetString("sid")
	if len(uname) == 0 {
		this.redirect("/")
		return
	}

	s := this.GetSession("SessionUser")
	if s != nil {
		userInfo := s.(models.BaseUserInfo)
		if sid != userInfo.SID {
			ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
			if _, ok := err.(websocket.HandshakeError); ok {
				http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
				return
			} else if err != nil {
				beego.Error("Cannot setup WebSocket connection:", err)
				return
			}

			cacheWebsocet.GetRoom().SendRedirectMsg(ws)
			return
		}

	} else {
		ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
		if _, ok := err.(websocket.HandshakeError); ok {
			http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
			return
		} else if err != nil {
			beego.Error("Cannot setup WebSocket connection:", err)
			return
		}

		cacheWebsocet.GetRoom().SendRedirectMsg(ws)
		return
	}

	//判断当前用户是不是已在线上。
	curSID := cacheWebsocet.GetRoom().GetCacheConnSID(uname)
	fmt.Println("curSID:", curSID, "SID:", sid)
	if curSID != sid {
		cacheWebsocet.GetRoom().CloseSID(curSID)
	}

	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	aWebsocket := cacheWebsocet.GetRoom().AddConn(uname, ws, sid, nil)

	strHint := fmt.Sprintf("[%s]加入房间", uname)
	aWebMsg := TtWebSocket.WebMsg{T: TtWebSocket.WebMsg_T_Jion, C: strHint}
	cacheWebsocet.GetRoom().BroadcastMsg(aWebMsg)

	aWebsocket.Run(onReadMessage)
}
func onReadMessage(conn *TtWebSocket.TtWebSocket, messageType int, p []byte) {

	aWebMsg := TtWebSocket.WebMsg{}
	err := json.Unmarshal(p, &aWebMsg)
	if err == nil {
		cacheWebsocet.GetRoom().BroadcastMsg(aWebMsg)
	}

	fmt.Println("onReadMessage", string(p))
}

//func (this *WebSocketController) Join() {
//	defer this.StopRun()
//	this.NewJoin(onReadMessage, nil)
//}
