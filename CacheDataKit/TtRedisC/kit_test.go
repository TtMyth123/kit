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

	testICache(mCache, t)
}

func testICache(mCache CacheDataKit.ICache, t *testing.T) {

	mCache.SetCache("a", 1, 0)

	a := 0
	b, e := mCache.GetCache("a", a)
	if e != nil {
		fmt.Println(e, b)
		t.Fail()
	}
	if b != 1 {
		t.Fail()
	}

	type Temp1 struct {
		A int
	}
	aa := Temp1{
		A: 100,
	}
	e = mCache.SetCache("aa", &aa, 0)
	if e != nil {
		fmt.Println(e, b)
		t.Fail()
	}

	bb, e := mCache.GetCache("aa", &aa)
	bb1 := bb.(*Temp1)
	if bb1.A != 100 {
		t.Fail()
	}

	e = mCache.SetCache("cc2", &aa, 0)
	if e != nil {
		fmt.Println(e, b)
		t.Fail()
	}

	cc, e := mCache.GetCache("cc2", &aa)
	cc1 := cc.(*Temp1)
	if cc1.A != 100 {
		t.Fail()
	}
}
