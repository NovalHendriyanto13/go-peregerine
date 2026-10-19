package QueueProvider

import (
	"github.com/hibiken/asynq"
	"peregerine/configs"
	QueueType "peregerine/systems/services/queue/types"
)

func redisOpt() asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr: configs.RedisHost + ":" + configs.RedisPort,
		Password: configs.RedisPassword,
		DB: configs.RedisDB,
	}
}

func Build() (*QueueType.QueueClient, error) {
	client := &QueueType.QueueClient{
		Client: asynq.NewClient(redisOpt()),
	}

	return client, nil
}

func BuildServer(cfg configs.WorkerConfig) (*QueueType.QueueServer) {
	serv := &QueueType.QueueServer{
		Server: asynq.NewServer(
			redisOpt(),
			asynq.Config{
				Concurrency: cfg.Concurrency,
				Queues: cfg.Queues,
			},
		),
	}

	return serv
}
