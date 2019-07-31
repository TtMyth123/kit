package cacheRedis

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"sync"
)
var
(
	onceClientredis sync.Once
	mClientredis *redis.Client
)

func init() {
	aDB, _ := beego.AppConfig.Int("Redis::RedisDB")
	aAddr := beego.AppConfig.String("Redis::RedisIP")
	aPassword := beego.AppConfig.String("Redis::RedisPassword")
	mClientredis = redis.NewClient(&redis.Options{
		Addr:     aAddr,
		Password: aPassword, // no password set
		DB:       aDB,                                    // use default DB
	})

	//	onceClientredis.Do(func() {
	//})

}

func GetClientRedisP() *redis.Client {
	return mClientredis
}
