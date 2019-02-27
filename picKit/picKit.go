package picKit

import "strings"

const PicServerIP = `http://120.78.212.181:7709/static/image/`

const (
	PT_Goods         = ""
	PT_WJ            = ""
	PT_LoginGift     = ""
	PT_MallGoods     = ""
	PT_TurnplateGift = ""
	PT_MailGift      = ""
)

func GetURL(picType string, url string) string {
	if len(url) > 4 && strings.ToLower(url[:4]) == "http" {
		return url
	} else {
		return PicServerIP + url
	}
}
