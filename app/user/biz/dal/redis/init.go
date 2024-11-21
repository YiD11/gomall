package redis

import (
	"context"

	"github.com/YiD11/gomall/app/user/conf"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	// RedisClient = redis.NewClient(&redis.Options{
	// 	Addr:     os.Getenv("REDIS_ADDRESS"),
	// 	Username: os.Getenv("REDIS_USERNAME"),
	// 	Password: os.Getenv("REDIS_PASSWORD"),
	// 	DB:       0,
	// })
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
