package RedisDataKit

import "encoding/json"

func (this RedisData) Get(key string) string {
	return this.mClientredis.Get(key).Val()
}

func (this RedisData) Set(key string, data string) error {
	return this.mClientredis.Set(key, data, 0).Err()
}

func (this RedisData) Add2EnQueue(key string, data interface{}) error {
	strJson, _ := json.Marshal(data)
	return this.mClientredis.RPush(key, strJson).Err()
}
func (this RedisData) GetDeQueue(key string) string {
	return this.mClientredis.LPop(key).Val()
}

func (this RedisData) GetLRange(key string, index, count int64) ([]string, error) {
	return this.mClientredis.LRange(key, index, count).Result()
}

func (this RedisData) AddByFloat(key string, value float64) (float64, error) {
	return this.mClientredis.IncrByFloat(key, value).Result()
}

func (this RedisData) GetByFloat(key string) (float64, error) {
	return this.mClientredis.Get(key).Float64()
}
