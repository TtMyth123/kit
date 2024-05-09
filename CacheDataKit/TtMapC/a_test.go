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

	testICache(mCache, t)
}

func testICache(mCache CacheDataKit.ICache, t *testing.T) {

	mCache.SetCache("a", 1, 0)

	a := 0
	e := mCache.GetCache("a", &a)
	if e != nil {
		fmt.Println(e)
		t.Fail()
	}
	if a != 1 {
		t.Fail()
	}

	type Temp1 struct {
		A int
	}
	aa := Temp1{
		A: 100,
	}
	aa1 := Temp1{
		A: 1000,
	}
	e = mCache.SetCache("aa", &aa, 0)
	if e != nil {
		t.Fail()
	}
	e = mCache.GetCache("aa", &aa1)
	if e != nil {
		t.Fail()
	}
	if aa1.A != 100 {
		t.Fail()
	}

	e = mCache.SetCache("cc2", &aa, 0)
	if e != nil {
		t.Fail()
	}

	aa.A = 2000
	aa1.A = 1000
	e = mCache.GetCache("cc2", &aa1)
	if e != nil {
		t.Fail()
	}

	if aa1.A != 100 {
		t.Fail()
	}
}
