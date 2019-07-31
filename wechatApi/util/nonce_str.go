package util

import "ttmyth123/other/rand"

func NonceStr() string {
	return string(rand.NewHex())
}
