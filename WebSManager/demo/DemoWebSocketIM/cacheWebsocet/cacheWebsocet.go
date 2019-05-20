package cacheWebsocet

import "ttmyth123/kit/WebSManager"

var (
	room *WebSManager.WebsocketManager
)

func init() {
	room = WebSManager.NewWebsocketManager(10)
}

func GetRoom() (*WebSManager.WebsocketManager) {
	return room
}
