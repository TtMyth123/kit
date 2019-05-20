package redisKit

import "github.com/go-redis/redis"

const gOnePageItemCount = 20
// map[string] 方式返回 HMGet 数据.
func GetMapHMGet(redisClient *redis.Client, key string, fields ...string) (map[string]interface{}, error) {
	mpResult := map[string]interface{}{}
	values, e := redisClient.HMGet(key, fields...).Result()
	for i, value := range values {
		mpResult[fields[i]] = value
	}

	return mpResult, e
}

func ExistZMember(redisClient *redis.Client, key, member string) bool {
	_, e := redisClient.ZRank(key, member).Result()
	return e == nil
}

func ExistKey(redisClient *redis.Client, key string) bool {
	_, e := redisClient.Exists(key).Result()
	return e == nil
}

func GetPageInfo(allItemCount int64, page int64) (pageCount, start, stop int64) {
	pageCount = allItemCount / + 1
	if allItemCount%gOnePageItemCount == 0 {
		pageCount = pageCount - 1
	}

	start = (page - 1) * gOnePageItemCount
	stop = page*gOnePageItemCount - 1
	return
}
