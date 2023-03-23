package configs

import "github.com/redis/go-redis/v9"

var RedisOptions = redis.Options{
	Addr:     Env.REDIS_ADDRESS,
	Password: "",
	DB:       0,
}
