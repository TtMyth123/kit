package TtWebSocket


type HandlerFunc_ReadMessage func(this *TtWebSocket, messageType int, p []byte)

type HandlerFunc_CloseMessage func(id string)