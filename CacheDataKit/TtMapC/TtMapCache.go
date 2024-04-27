package TtMapC

import (
	"fmt"
	"github.com/TtMyth123/kit/CacheDataKit"
	"sync"
	"time"
)

type tmpT struct {
	T1 int
	T2 int
}

type TtMapCache struct {
	mpData sync.Map
	mpT    sync.Map
}

func NewTtMapCache() CacheDataKit.ICache {
	aTtRedisCache := TtMapCache{}
	return &aTtRedisCache
}
func (this *TtMapCache) tryDelCache() {
	this.mpT.Range(func(key, value any) bool {
		t := value.(*tmpT)
		if t.T1 > 1 {
			t.T1--
		}
		if t.T1 == 1 {
			k := key.(string)
			this.DelCache(k)
		}
		return true
	})
}
func (this *TtMapCache) run() {
	ticker1 := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker1.C:
			this.tryDelCache()
		}
	}
}

func (this *TtMapCache) StartAndGC(mpConfig map[string]interface{}) error {
	go this.run()
	return nil
}

func (this *TtMapCache) SetCache(key string, value interface{}, timeout int) error {
	this.mpData.Store(key, value)
	this.storeTime(key, timeout)
	return nil
}

func (this *TtMapCache) reStoreTime(key string) {
	if vt, ok2 := this.mpT.Load(key); ok2 {
		t := vt.(*tmpT)
		t.T1 = t.T2
		this.mpT.Store(key, t)
	} else {
		t := &tmpT{T2: -1, T1: -1}
		this.mpT.Store(key, t)
	}
}

func (this *TtMapCache) storeTime(key string, timeout int) {
	t := &tmpT{T2: timeout, T1: timeout}
	this.mpT.Store(key, t)
}
func (this *TtMapCache) GetCache(key string, to interface{}) (any, error) {
	if data, ok := this.mpData.Load(key); ok {
		//value := reflect.ValueOf(to)
		to = data
		this.reStoreTime(key)
		return to, nil
	}

	return to, fmt.Errorf("Cache不存在")
}

//func (this *TtMapCache) GetCacheData(key string) (any, error) {
//	if data, ok := this.mpData.Load(key); ok {
//		//value := reflect.ValueOf(to)
//		to := data.(any)
//		this.reStoreTime(key)
//		return to, nil
//	}
//
//	var to any
//	return to, fmt.Errorf("Cache不存在")
//}

func (this *TtMapCache) DelCache(key string) error {
	this.mpData.Delete(key)
	this.mpT.Delete(key)
	return nil
}
