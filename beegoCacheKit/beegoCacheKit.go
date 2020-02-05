package beegoCacheKit

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
)

const Timeout8  = 3600*24*365*10
var
(
	mpBeegoCache map[string]*BeegoCache
)

func init() {
	mpBeegoCache = make(map[string]*BeegoCache)
}

type BeegoCache struct {
	cc   cache.Cache
	name string
}

func NewBeegoCache(name string) (*BeegoCache) {
	aBeegoCache := new(BeegoCache)
	aBeegoCache.name = name
	mpBeegoCache[name] = aBeegoCache
	return aBeegoCache
}
func GetBeegoCacheIns(name string) *BeegoCache {
	return mpBeegoCache[name]
}

func (this *BeegoCache) InitCache(host, pwd string, dbIndex int) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			this.cc = nil
		}
	}()
	config := fmt.Sprintf(`{"conn":"%s","dbNum":"%d","password":"%s"}`, host, dbIndex, pwd)
	this.cc, err = cache.NewCache("redis", config)
	return err
}

// SetCache
func (this *BeegoCache) SetCache(key string, value interface{}, timeout int) error {
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
	timeouts := time.Duration(timeout) * time.Second
	err = this.cc.Put(key, data, timeouts)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (this *BeegoCache) GetCache(key string, to interface{}) error {
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

// DelCache
func (this *BeegoCache) DelCache(key string) error {
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
