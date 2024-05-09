package SyncMapCache

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	aTtMapCache := NewTtMapCache()
	type T1 struct {
		A1 int
	}
	type T2 struct {
		A1 int
	}
	a1 := T1{A1: 1000}
	a2 := T1{A1: 2000}
	aTtMapCache.SetCache("aa", a1, -1)
	aa1, e := aTtMapCache.GetCache("aa", T1{})

	a2 = aTtMapCache.GetCacheEx("aa", T2{}).(T1)

	fmt.Println(a2, e, aa1)

}
