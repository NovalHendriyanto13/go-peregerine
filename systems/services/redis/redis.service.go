package redis

import(
	GoRedis "github.com/redis/go-redis/v9"
	RedisRepository "peregerine/systems/services/redis/repository"
	RedisProvider "peregerine/systems/services/redis/provider"
	RedisType "peregerine/systems/services/redis/types"
)

type Container struct {
	RedisRepo RedisType.SettingService
	client *GoRedis.Client
}

func Build() (*Container, error) {
	rdb, err := RedisProvider.Build()
	if err != nil {
		return nil, err
	}
	repo := RedisRepository.Build(rdb)

	return &Container{ RedisRepo: repo, client: rdb }, nil
}

// Close releases the redis connection pool
func (c *Container) Close() error {
	return c.client.Close()
}
