package di

import (
	Redis "peregerine/systems/services/redis"
	"github.com/redis/go-redis/v9"
)

type DiHandler struct {
	redis *Redis
}

func Build() (DiHandler, err) {
	r := Redis.Build()

	return &DiHandler{
		redis := r
	}, nil
} 