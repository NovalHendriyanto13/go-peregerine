package QueueProvider

import (
	"github.com/hibiken/asynq"
	"peregerine/configs"
	QueueType "peregerine/systems/services/queue/types"
)

func Build() (*QueueType.QueueClient, error) {
	redisAddr := configs.RedisHost + ":" + configs.RedisPort
	client := &QueueType.QueueClient{
		Client: asynq.NewClient(asynq.RedisClientOpt {
			Addr: redisAddr,
		}),
	}

	return client, nil
}

func BuildServer(cfg configs.WorkerConfig) (*QueueType.QueueServer) {
	redisAddr := configs.RedisHost + ":" + configs.RedisPort
	serv := &QueueType.QueueServer{
		Server: asynq.NewServer(
			asynq.RedisClientOpt { Addr: redisAddr },
			asynq.Config{
				Concurrency: cfg.Concurrency,
				Queues: cfg.Queues,
			},
		),
	}

	return serv
}