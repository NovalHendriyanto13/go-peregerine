package RedisRepository

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	RedisType "peregerine/systems/services/redis/types"
)

type RedisSettingRepo struct {
	rdb *redis.Client
}

func Build(rdb *redis.Client) RedisType.SettingService {
	return &RedisSettingRepo{ rdb: rdb }
}

func (r *RedisSettingRepo) Set(ctx context.Context, group, key, value string) error {
	return r.rdb.HSet(ctx, group, key, value).Err()
}

func (r *RedisSettingRepo) Get(ctx context.Context, key, value string) (string, error) {
	val, err := r.rdb.HGet(ctx, key, value).Result()
	if errors.Is(err, redis.Nil) {
		return "", RedisType.ErrNotFound
	}
	return val, err
}

func (r *RedisSettingRepo) GetAll(ctx context.Context, group string) (map[string]string, error){
	return r.rdb.HGetAll(ctx, group).Result()
}
