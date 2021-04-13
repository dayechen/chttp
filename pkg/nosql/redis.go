package nosql

import (
	"cweb/global"

	"github.com/go-redis/redis"
)

func NewRDBEngine() (*redis.Client, error) {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Host,
		Password: global.RedisSetting.Password,
		DB:       global.RedisSetting.DBNumber,
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		return nil, err
	}
	return redisDB, nil
}
