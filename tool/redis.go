package tool

import (
	"github.com/go-redis/redis"
	"time"
)

var Rdb *redis.Client

func Redis() (err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     GetViper().GetString("redis.host") + ":" + GetViper().GetString("redis.port"),
		Password: GetViper().GetString("redis.password"),
		DB:       GetViper().GetInt("redis.db"),
	})
	Rdb = client
	return nil
}

func Set(keyUserId string, value interface{}, duration time.Duration) (err error) {
	err = Rdb.Set(keyUserId, value, duration).Err()
	return
}

func Get(keyUserId string) (value string, err error) {
	value, err = Rdb.Get(keyUserId).Result()
	return
}
