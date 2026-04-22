package RedisRepository

import (
	"context"
	"github.com/redis/go-redis/v9"
	RedisType "peregerine/systems/services/redis/types"
)

type RedisSettingRepo struct {
	rdb *redis.Client
	ctx context.Context
}

func Build(
	rdb *redis.Client,
	ctx context.Context,
) RedisType.SettingService {
	return &RedisSettingRepo{ rdb: rdb, ctx: ctx }
}

func (r *RedisSettingRepo) Set(group, key, value string) error {
	return r.rdb.HSet(r.ctx, group, key, value).Err()
}

func (r *RedisSettingRepo) Get(key, value string) (string, error) {
	return r.rdb.HGet(r.ctx, key, value).Result()
}

func (r *RedisSettingRepo) GetAll(group string) (map[string]string, error){
	return r.rdb.HGetAll(r.ctx, group).Result()
}

