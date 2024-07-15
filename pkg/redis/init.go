package redis

import "github.com/redis/go-redis/v9"

func InitDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:16379",
		Password: "",
		DB:       0,
	})
}
