package internal

import (
	"context"
	"tokay/configs"

	"github.com/redis/go-redis/v9"
)

var store = ConnectCacheStorage(&configs.RedisOptions)
var ctx = context.Background()

type CacheStorage struct {
	connection *redis.Client
}

func ConnectCacheStorage(options *redis.Options) *CacheStorage {
	connection := redis.NewClient(options)
	return &CacheStorage{connection: connection}
}

func (cache *CacheStorage) DisconnectCacheStorage() error {
	return cache.connection.Close()
}

func (cache *CacheStorage) IsBlacklisted(key, token string) (bool, error) {
	ismember, err := cache.connection.SIsMember(ctx, key, token).Result()
	if !ismember || err != nil {
		return false, err
	}
	return true, nil
}

func (cache *CacheStorage) Blacklist(key, token string) error {
	return cache.connection.SAdd(ctx, key, token).Err()
}
