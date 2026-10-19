package RedisProvider

import (
	"context"
	"time"
	"github.com/redis/go-redis/v9"
	"peregerine/configs"
)

func Build() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: configs.RedisHost + ":" + configs.RedisPort,
		Password: configs.RedisPassword,
		DB: configs.RedisDB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		_ = rdb.Close()
		return nil, err
	}

	return rdb, nil
}
