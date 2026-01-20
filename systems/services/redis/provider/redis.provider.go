package redisprovider

import (
	"context"
	"github.com/redis/go-redis/v9"
	"peregerine/configs"
)

func RedisProvider() (*redis.Client, context.Context, error) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: configs.RedisHost + ":" + configs.RedisPort,
		DB: configs.RedisDB
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, nil, err
	}

	return rdb, ctx, nil
}