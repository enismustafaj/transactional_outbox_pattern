package database

import (
	"log"
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	client *redis.Client
	ctx context.Context
}

var redisConnection *RedisDB

func NewRedisDBConnection(options *redis.Options) *RedisDB {
	if redisConnection == nil {
		lock.Lock()
		defer lock.Unlock()

		return &RedisDB{
			client: redis.NewClient(options),
			ctx: context.Background(),
		}
	} 

	return redisConnection
}

func (redisDB *RedisDB) SetKey(key string, val int) {
	err:= redisDB.client.Set(redisDB.ctx, key, val, 0).Err()

	if err != nil {
		log.Fatal("Error setting the tokens")
	}
}

func (redisDB *RedisDB) GetVal(key string) string {
	val, err:= redisDB.client.Get(redisDB.ctx, key).Result()

	if err != nil {
		log.Fatal("Error setting the tokens")
	}

	return val
}	