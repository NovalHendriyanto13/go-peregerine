package redis

import(
	RedisRepository "peregerine/systems/services/redis/repository"
	RedisProvider "peregerine/systems/services/redis/provider"
)

type Container struct {
	RedisRepo *RedisRepository.RedisSettingRepo
}

func Build() (*Container, error) {
	rdb, ctx, err := RedisProvider.Build()
	if err != nil {
		return nil, err
	}

	repo := RedisRepository.Build(rdb, ctx)

	return &Container{ RedisRepo: repo }, nil
}