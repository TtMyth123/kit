package TtRedisC

import (
	"fmt"
	"github.com/TtMyth123/kit/CacheDataKit"
	"testing"
)

func TestTtRedisCache(t *testing.T) {
	CacheDataKit.Register(CacheDataKit.AdapterName_Redis, NewTtRedisCache)
	mpConfig := make(map[string]interface{})

	mpConfig["host"] = "127.0.0.1:6379"
	mpConfig["index"] = 1
	mpConfig["password"] = "3.dirdir"
	mpConfig["name"] = "sTest"

	mCache, e := CacheDataKit.NewCache(CacheDataKit.AdapterName_Redis, mpConfig)
	if e != nil {
		fmt.Println(e)
	}

	mCache.SetCache("a", 1, 0)

	a := 0
	e = mCache.GetCache("a", &a)
	if e != nil {
		fmt.Println(e)
	}
}
