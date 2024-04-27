package TtMapC

import (
	"fmt"
	"github.com/TtMyth123/kit/CacheDataKit"
	"testing"
)

func TestTtRedisCache(t *testing.T) {
	CacheDataKit.Register(CacheDataKit.AdapterName_TtMapC, NewTtMapCache)
	mpConfig := make(map[string]interface{})

	mCache, e := CacheDataKit.NewCache(CacheDataKit.AdapterName_TtMapC, mpConfig)
	if e != nil {
		fmt.Println(e)
	}

	mCache.SetCache("a", 1, 0)

	a := 0
	b, e := mCache.GetCache("a", a)
	if e != nil {
		fmt.Println(e, b)
	}
}
