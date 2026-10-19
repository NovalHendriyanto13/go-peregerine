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
}

// DiWorker handles dependency injection for background workers.
type DiWorker struct {
	Redis *RedisService.Container
	Logger *LoggerService.Logger
	QueueServer *QueueService.ContainerServer
}

// Build creates the dependencies needed by the HTTP server
func Build() (*DiHandler, error) {
	l, err := LoggerService.Build()
	if err != nil {
		return nil, err
	}

	r, err := RedisService.Build()
	if err != nil {
		return nil, err
	}

	q, err := QueueService.Build()
	if err != nil {
		_ = r.Close()
		return nil, err
	}

	return &DiHandler{
		Redis: r,
		Logger: l,
		Queue: q,
	}, nil
}

// Close releases all resources held by the HTTP server dependencies
func (d *DiHandler) Close() {
	_ = d.Queue.Close()
	_ = d.Redis.Close()
	_ = d.Logger.Log.Sync()
}

// BuildWorker creates the dependencies needed by the queue worker
func BuildWorker() (*DiWorker, error) {
	l, err := LoggerService.Build()
	if err != nil {
		return nil, err
	}

	r, err := RedisService.Build()
	if err != nil {
		return nil, err
	}

	qe := QueueService.BuildServer(configs.WorkerConfig{
		Concurrency: 10,
		Queues: map[string]int{
			"default": 1,
		},
	})

	return &DiWorker{
		Redis: r,
		Logger: l,
		QueueServer: qe,
	}, nil
}

// Close releases all resources held by the worker dependencies
func (d *DiWorker) Close() {
	_ = d.Redis.Close()
	_ = d.Logger.Log.Sync()
}
