package qrcode

import (
	"ttmyth123/kit/wechatApi/mp/base"
	"ttmyth123/kit/wechatApi/mp/core"
)

// ShortURL 将一条长链接转成短链接.
func ShortURL(clt *core.Client, longURL string) (shortURL string, err error) {
	return base.ShortURL(clt, longURL)
}
