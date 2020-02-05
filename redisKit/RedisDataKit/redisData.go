package RedisDataKit

import (
	"github.com/go-redis/redis"
	"sync"
)

var (
	mRedisData      *RedisData
	mpRedisData sync.Map
)

type RedisData struct {
	mClientredis *redis.Client
}

/**
	初始化创建一个 RedisData
 */
func IniNew(name, RedisIP, RedisPwd string, RedisDBIndex int) {
	mRedisData = new(RedisData)
	mRedisData.mClientredis = redis.NewClient(&redis.Options{
		Addr:    RedisIP,
		Password: RedisPwd, // no password set
		DB:       RedisDBIndex,                                    // use default DB
	}) 
}

/**
删除 RedisData
 */
func DelRedisData(name string) {
	mpRedisData.Delete(name)
}

//获取 Redis 实例
func GetRedisData(name string) *RedisData {
	return mRedisData

	if redisData, ok := mpRedisData.Load(name); ok {
		aRedisData := redisData.(*RedisData)
		return aRedisData
	}
	return nil
}

//获取 Redis 实例
func GetClientRedisP(name string) (*redis.Client) {
	return mRedisData.mClientredis

	if redisData, ok := mpRedisData.Load(name); ok {
		aRedisData := redisData.(*RedisData)
		return aRedisData.mClientredis
	}
	return nil
}


