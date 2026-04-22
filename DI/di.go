package di

import (
	"peregerine/configs"
	RedisService "peregerine/systems/services/redis"
	LoggerService "peregerine/systems/services/logger"
	QueueService "peregerine/systems/services/queue"
)

// DiHandler handles dependency injection for HTTP handlers.
type DiHandler struct {
	Redis *RedisService.Container
	Logger *LoggerService.Logger
	Queue *QueueService.Container
	QueueServer *QueueService.ContainerServer
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

	q, err := QueueService.Build()
	if err != nil {
		return nil, err
	}

	qe := QueueService.BuildServer(configs.WorkerConfig{
		Concurrency: 10,
		Queues: map[string]int{
			"default": 1,
		},
	})
	
	return &DiHandler{
		Redis: r,
		Logger: l,
		Queue: q,
		QueueServer: qe,
	}, nil
}