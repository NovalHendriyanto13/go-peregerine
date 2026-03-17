package di

import (
	RedisService "peregerine/systems/services/redis"
	LoggerService "peregerine/systems/services/logger"
)

// DiHandler handles dependency injection for HTTP handlers.
type DiHandler struct {
	Redis *RedisService.Container
	Logger *LoggerService.Logger
}

func Build() (*DiHandler, error) {
	r, err := RedisService.Build()
	if err != nil {
		return nil, err
	}

	l, err := LoggerService.Build()
	if err != nil {
		return nil, err
	}

	return &DiHandler{
		Redis: r,
		Logger: l,
	}, nil
}