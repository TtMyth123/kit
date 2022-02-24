package TtRedisClient

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type TtRedisClient struct {
	mClientRedis *redis.Client
	prefix       string
}

func (this *TtRedisClient) GetNewKey(key string) string {
	return fmt.Sprintf("%s%s", this.prefix, key)
}

func (this *TtRedisClient) GetClientRedis() *redis.Client {
	return this.mClientRedis
}

func (this *TtRedisClient) CopyTo(mClientRedis *redis.Client) error {
	keys, e := this.GetClientRedis().Keys("*").Result()
	if e != nil {
		return e
	}

	for _, v := range keys {
		t := this.GetClientRedis().Type(v).Val()
		switch t {
		case "set":
			data := this.GetClientRedis().SMembers(v).Val()
			iidata := make([]interface{}, len(data))
			for ii, data1 := range data {
				iidata[ii] = data1
			}
			mClientRedis.SAdd(v, iidata...)
		case "list":
			i1 := int64(0)
			i2 := i1 + 100
			for {
				data := this.GetClientRedis().LRange(v, i1, i2).Val()
				iLen := len(data)
				iidata := make([]interface{}, iLen)
				for ii, data1 := range data {
					iidata[ii] = data1
				}
				mClientRedis.RPush(v, iidata...)
				if iLen != int(i2-i1+1) {
					break
				}
				i1 = i2 + 1
				i2 = i1 + 100
			}
		case "hash":
			data := this.GetClientRedis().HGetAll(v).Val()
			iidata := make(map[string]interface{})
			for ii, data1 := range data {
				iidata[ii] = data1
			}
			mClientRedis.HMSet(v, iidata)
		case "string":
			data := this.GetClientRedis().Get(v).Val()
			mClientRedis.Set(v, data, 0)
		}
	}

	return nil
}

func NewTtRedisClient(prefix, RedisIP, RedisPwd string, RedisDBIndex int) *TtRedisClient {
	aTtRedisClient := new(TtRedisClient)
	aTtRedisClient.mClientRedis = redis.NewClient(&redis.Options{
		Addr:     RedisIP,
		Password: RedisPwd,     // no password set
		DB:       RedisDBIndex, // use default DB
	})
	aTtRedisClient.prefix = prefix
	return aTtRedisClient
}

func (this *TtRedisClient) GetCache(key string, to interface{}) error {
	key = this.GetNewKey(key)
	data, e := this.mClientRedis.Get(key).Bytes()
	if e != nil {
		return e
	}
	err := Decode(data, to)
	return err
}

func (this *TtRedisClient) SetCache(key string, value interface{}, timeout int) error {
	key = this.GetNewKey(key)
	data, err := Encode(value)

	if err != nil {
		return err
	}

	timeouts := time.Duration(timeout) * time.Second
	return this.mClientRedis.Set(key, data, timeouts).Err()
}

func (this *TtRedisClient) DelCache(key string) error {
	key = this.GetNewKey(key)
	return this.mClientRedis.Del(key).Err()
}

func (this *TtRedisClient) Get(key string) string {
	key = this.GetNewKey(key)
	return this.mClientRedis.Get(key).Val()
}

func (this *TtRedisClient) Set(key string, data string) error {
	key = this.GetNewKey(key)
	return this.mClientRedis.Set(key, data, 0).Err()
}
func (this *TtRedisClient) SetByTime(key string, data string, expiration time.Duration) error {
	key = this.GetNewKey(key)
	return this.mClientRedis.Set(key, data, expiration).Err()
}

func (this *TtRedisClient) Add2EnQueue(key string, data interface{}) error {
	key = this.GetNewKey(key)
	strJson, _ := json.Marshal(data)
	return this.mClientRedis.RPush(key, strJson).Err()
}
func (this *TtRedisClient) GetDeQueue(key string) string {
	key = this.GetNewKey(key)
	return this.mClientRedis.LPop(key).Val()
}

func (this *TtRedisClient) GetLRange(key string, index, count int64) ([]string, error) {
	key = this.GetNewKey(key)
	return this.mClientRedis.LRange(key, index, count).Result()
}

func (this *TtRedisClient) AddByFloat(key string, value float64) (float64, error) {
	key = this.GetNewKey(key)
	return this.mClientRedis.IncrByFloat(key, value).Result()
}
func (this *TtRedisClient) IncrBy(key string, value int64) (int64, error) {
	key = this.GetNewKey(key)
	return this.mClientRedis.IncrBy(key, value).Result()
}

func (this *TtRedisClient) GetByFloat(key string) (float64, error) {
	key = this.GetNewKey(key)
	return this.mClientRedis.Get(key).Float64()
}

func (this *TtRedisClient) Exist(key string) bool {
	key = this.GetNewKey(key)
	v := this.mClientRedis.Exists(key).Val()
	return v > 0
}
func (this *TtRedisClient) Del(key string) bool {
	key = this.GetNewKey(key)
	v := this.mClientRedis.Del(key).Val()
	return v > 0
}

func (this *TtRedisClient) SAdd(key string, data ...string) int64 {
	key = this.GetNewKey(key)
	r := this.mClientRedis.SAdd(key, data).Val()
	return r
}

func (this *TtRedisClient) SIsMember(key string, data string) bool {
	key = this.GetNewKey(key)
	c := this.mClientRedis.SIsMember(key, data).Val()
	return c
}

func (this *TtRedisClient) SCard(key string) int64 {
	key = this.GetNewKey(key)
	data := this.mClientRedis.SCard(key).Val()
	return data
}

func (this *TtRedisClient) LPush(key string, data ...string) int64 {
	key = this.GetNewKey(key)
	c := this.mClientRedis.LPush(key, data).Val()
	return c
}

/**
取从位置0开始到位置2结束的3个元素。
lrange mykey 0 2
*/
func (this *TtRedisClient) LRange(key string, start, stop int64) []string {
	key = this.GetNewKey(key)
	data := this.mClientRedis.LRange(key, start, stop).Val()
	return data
}
func (this *TtRedisClient) LLen(key string) int64 {
	key = this.GetNewKey(key)
	data := this.mClientRedis.LLen(key).Val()
	return data
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
