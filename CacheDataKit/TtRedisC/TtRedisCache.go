package TtRedisC

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/TtMyth123/kit/CacheDataKit"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
)

type TtRedisCache struct {
	cc   cache.Cache
	name string
}

func NewTtRedisCache() CacheDataKit.ICache {
	aTtRedisCache := TtRedisCache{}
	return &aTtRedisCache
}

func (this *TtRedisCache) StartAndGC(mpConfig map[string]interface{}) error {
	host := mpConfig[ConfigKey_host]
	index := mpConfig[ConfigKey_index]
	password := mpConfig[ConfigKey_password]
	name := fmt.Sprint(mpConfig[ConfigKey_name])

	config := fmt.Sprintf(`{"conn":"%s","dbNum":"%v","password":"%s"}`, host, index, password)
	cc, err := cache.NewCache(CacheDataKit.AdapterName_Redis, config)
	if err != nil {
		return err
	}
	this.cc = cc
	this.name = name
	return nil
}
func (this *TtRedisCache) SetCache(key string, value interface{}, timeout int) error {
	key = this.name + key
	data, err := Encode(value)
	if err != nil {
		return err
	}
	if this.cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			this.cc = nil
		}
	}()
	if timeout == 0 {
		timeout = 20000
	}
	timeouts := time.Duration(timeout) * time.Second

	err = this.cc.Put(key, data, timeouts)
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (this *TtRedisCache) GetCache(key string, to interface{}) error {
	key = this.name + key

	if this.cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			this.cc = nil
		}
	}()

	data := this.cc.Get(key)
	if data == nil {
		return errors.New("Cache不存在")
	}

	err := Decode(data.([]byte), to)
	if err != nil {

	}

	return err
}
func (this *TtRedisCache) DelCache(key string) error {
	key = this.name + key

	if this.cc == nil {
		return errors.New("cc is nil")
	}
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			this.cc = nil
		}
	}()
	err := this.cc.Delete(key)
	if err != nil {
		return errors.New("Cache删除失败")
	} else {
		return nil
	}
}
func (this *TtRedisCache) GetCacheData(key string) (any, error) {
	key = this.name + key
	var to any
	if this.cc == nil {
		return to, errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			this.cc = nil
		}
	}()

	data := this.cc.Get(key)
	if data == nil {
		return to, errors.New("Cache不存在")
	}

	err := Decode(data.([]byte), to)
	if err != nil {

	}

	return to, err
}

// Encode
// 用gob进行数据编码
//
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode
// 用gob进行数据解码
//
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
