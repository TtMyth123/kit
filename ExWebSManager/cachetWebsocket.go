package ExWebSManager

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
	"time"
	"ttmyth123/kit"
	"ttmyth123/kit/ExWebSManager/TtWebSocket"
	"ttmyth123/kit/ExWebSManager/TtWebSocket/ErrorWeb"
)

type WebsocketManager struct {
	//mapConn    map[string]*TtWebSocket.TtWebSocket
	//mapSIDConn map[string]*TtWebSocket.TtWebSocket
	//游客连接
	mapConn sync.Map
	//注册用户连接
	mapUserConn sync.Map
	//全部连接
	mapSIDConn sync.Map

	onCloseMessage TtWebSocket.HandlerFunc_CloseMessage
	CCount         int

	userCount int

	addConnLock sync.Mutex
}

func NewWebsocketManager(pingTime int) (*WebsocketManager) {
	aWebsocketManager := new(WebsocketManager)
	aWebsocketManager.init()
	aWebsocketManager.lifeDec(pingTime)
	aWebsocketManager.CCount = 0;
	aWebsocketManager.userCount = 0;
	return aWebsocketManager
}
func (this *WebsocketManager) init() {
	//this.mapConn = make(map[string]*TtWebSocket.TtWebSocket)
	//this.mapSIDConn = make(map[string]*TtWebSocket.TtWebSocket)
}
func (this *WebsocketManager) CCountAdd(c int) {
	this.CCount = this.CCount + c
}

func (this *WebsocketManager) UserCountAdd(c int) {
	this.CCount = this.CCount + c
	this.userCount = this.userCount + c
}
func (this *WebsocketManager) GetCCount() int {
	return this.CCount
}
func (this *WebsocketManager) AddConn(id string, conn *websocket.Conn, userInfo interface{}) (*TtWebSocket.TtWebSocket) {
	this.addConnLock.Lock()
	defer this.addConnLock.Unlock()

	SID := kit.GetGuid()
	v, ok := this.mapConn.Load(id)
	if ok {
		aCacheWebsocket := v.(*TtWebSocket.TtWebSocket)
		this.mapConn.Delete(id)
		this.CloseSID(aCacheWebsocket.SID)
	} else {
		this.CCountAdd(1)
	}

	aCacheWebsocket := &TtWebSocket.TtWebSocket{Conn: conn, SID: SID, UserInfo: userInfo}
	this.mapConn.Store(id, aCacheWebsocket)
	this.mapSIDConn.Store(SID, aCacheWebsocket)

	return aCacheWebsocket
}

func (this *WebsocketManager) AddUserConn(id, userid string, conn *websocket.Conn, userInfo interface{}) (*TtWebSocket.TtWebSocket) {
	this.addConnLock.Lock()
	defer this.addConnLock.Unlock()
	SID := kit.GetGuid()

	v, ok := this.mapUserConn.Load(userid)
	if ok {
		aCacheWebsocket := v.(*TtWebSocket.TtWebSocket)
		this.mapUserConn.Delete(userid)
		this.CloseSID(aCacheWebsocket.SID)
	} else {
		this.UserCountAdd(1)
	}

	aCacheWebsocket := &TtWebSocket.TtWebSocket{Conn: conn, SID: SID, UserInfo: userInfo}
	this.mapUserConn.Store(userid, aCacheWebsocket)
	this.mapSIDConn.Store(SID, aCacheWebsocket)

	return aCacheWebsocket
}

func (this *WebsocketManager) GetUserInfoInt(id int) interface{} {
	return this.GetUserInfo(strconv.Itoa(id))
}

func (this *WebsocketManager) GetUserInfo(id string) interface{} {
	if value, ok := this.mapConn.Load(id); ok {
		v, ok := value.(*TtWebSocket.TtWebSocket)
		if ok {
			return v.UserInfo
		}
	}
	return nil
}
func (this *WebsocketManager) lifeDec(pingTime int) {
	if pingTime < 2 {
		pingTime = 2;
	}
	ticker1 := time.NewTicker(time.Duration(pingTime) * time.Second)
	go func(t *time.Ticker) {
		for {
			<-t.C
			this.ping()
		}
	}(ticker1)
}

/**
发送ping判断是否断线
 */
func (this *WebsocketManager) ping() {
	this.mapConn.Range(func(key, value interface{}) bool {
		id, ok := key.(string)
		if !ok {
			return true
		}
		v, ok := value.(*TtWebSocket.TtWebSocket)
		if !ok {
			return true
		}
		e := v.SendPingMsg()
		if e != nil {
			if this.onCloseMessage != nil {
				this.CCountAdd(-1)
				this.mapConn.Delete(id)
				this.onCloseMessage(id)
			}
		}
		return true
	})
}

func (this *WebsocketManager) GetUserConnSID(id string) string {
	if value, ok := this.mapUserConn.Load(id); ok {
		v, ok := value.(*TtWebSocket.TtWebSocket)
		if ok {
			return v.SID
		}
	}

	return ""
}

func (this *WebsocketManager) CloseSID(sid string) {
	if value, ok := this.mapSIDConn.Load(sid); ok {
		if c, ok := value.(*TtWebSocket.TtWebSocket); ok {
			this.SendRedirectMsg(c.Conn)
		}
		this.mapSIDConn.Delete(sid)
	}
}

func (this *WebsocketManager) SendRedirectMsg(conn *websocket.Conn) *ErrorWeb.ErrorWebS {
	aWebMsg := TtWebSocket.WebMsg{T: TtWebSocket.WebMsg_T_redirect}
	data, err := json.Marshal(aWebMsg)
	if err != nil {
		return ErrorWeb.NewErrorWeb(ErrorWeb.EC_Marshal, err.Error())
	}
	err = conn.WriteMessage(websocket.TextMessage, data)
	return nil
}

func (this *WebsocketManager) SendMsg(id string, aWebMsg TtWebSocket.WebMsg) *ErrorWeb.ErrorWebS {
	if value, ok := this.mapConn.Load(id); ok {
		conn, ok := value.(*TtWebSocket.TtWebSocket)
		if ok {
			return conn.SendMsg(aWebMsg)
		} else {
			return ErrorWeb.NewErrorWeb(ErrorWeb.EC_ConnNil, "无连接")
		}
	} else {
		return ErrorWeb.NewErrorWeb(ErrorWeb.EC_ConnNil, "无连接")
	}
}

func (this *WebsocketManager) Send2OnlineUserMsg(userId string, aWebMsg TtWebSocket.WebMsg) *ErrorWeb.ErrorWebS {
	if value, ok := this.mapUserConn.Load(userId); ok {
		conn, ok := value.(*TtWebSocket.TtWebSocket)
		if ok {
			return conn.SendMsg(aWebMsg)
		} else {
			return ErrorWeb.NewErrorWeb(ErrorWeb.EC_ConnNil, "无连接")
		}
	} else {
		return ErrorWeb.NewErrorWeb(ErrorWeb.EC_ConnNil, "无连接")
	}

}


func (this *WebsocketManager) BroadcastMsg(aWebMsgs ...TtWebSocket.WebMsg) *ErrorWeb.ErrorWebS {
	this.mapSIDConn.Range(func(key, value interface{}) bool {
		c, ok := value.(*TtWebSocket.TtWebSocket)
		if ok {
			fmt.Println("Key:", key, "C", c.SID)
			for _, v := range aWebMsgs {
				c.SendMsg(v)
			}
		} else {
			fmt.Println("Key:", key, "C___ERR")
		}
		return true
	})
	return nil
}
