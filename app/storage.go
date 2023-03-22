package app

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type CacheStorage struct {
	connection *redis.Client
}

func ConnectCacheStorage(options *redis.Options) *CacheStorage {
	connection := redis.NewClient(options)
	return &CacheStorage{connection: connection}
}

func (cache *CacheStorage) DisconnectCacheStorage() error {
	err := cache.connection.Close()
	if err != nil {
		return err
	}
	return nil
}

func (cache *CacheStorage) IsBlacklisted(key, token string) (bool, error) {
	return false, nil
}


func (cache *CacheStorage) Backlist(key, token string) error {
	return nil
}
