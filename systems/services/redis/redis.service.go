package redis

import(
	RedisRepository "peregerine/systems/services/redis/repository"
	RedisProvider "peregerine/systems/services/redis/provider"
	RedisType "peregerine/systems/services/redis/types"
)

type Container struct {
	RedisRepo RedisType.SettingService
}

func Build() (*Container, error) {
	rdb, ctx, err := RedisProvider.Build()
	if err != nil {
		return nil, err
	}
	repo := RedisRepository.Build(rdb, ctx)

	return &Container{ RedisRepo: repo }, nil
}