package util

import "ttmyth123/kit/other/rand"

func NonceStr() string {
	return string(rand.NewHex())
}
