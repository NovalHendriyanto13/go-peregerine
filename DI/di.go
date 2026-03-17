package di

import (
	RedisService "peregerine/systems/services/redis"
)

// DiHandler handles dependency injection for HTTP handlers.
type DiHandler struct {
	Redis *RedisService.Container
}

func Build() (*DiHandler, error) {
	r, err := RedisService.Build()
	if err != nil {
		return nil, err
	}

	return &DiHandler{
		Redis: r,
	}, nil
}