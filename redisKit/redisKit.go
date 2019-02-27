package redisKit

import "github.com/go-redis/redis"

// map[string] 方式返回 HMGet 数据.
func GetMapHMGet(redisClient *redis.Client, key string, fields ...string) (map[string]interface{}, error) {
	mpResult := map[string]interface{}{}
	values, e := redisClient.HMGet(key, fields...).Result()
	for i, value := range values {
		mpResult[fields[i]] = value
	}

	return mpResult, e
}
