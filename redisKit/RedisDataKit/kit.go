package RedisDataKit

func (this RedisData) Get(key string) string {
	return this.mClientredis.Get(key).Val()
}

func (this RedisData) Set(key string, data string) error {
	return this.mClientredis.Set(key, data, 0).Err()
}
