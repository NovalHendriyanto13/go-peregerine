package QueueRepository

import (
	"time"
	"encoding/json"
	"github.com/hibiken/asynq"
	QueueType "peregerine/systems/services/queue/types"
)

type QueueRepo struct {
	client *QueueType.QueueClient
}

func Build(q *QueueType.QueueClient) QueueType.QueueMethod {
	return &QueueRepo{ client: q }
}

func (a *QueueRepo) SetQueue(taskName string, payload interface{}, opts ...asynq.Option) error {
	data, _ := json.Marshal(payload)

	defaultOption := []asynq.Option{
		asynq.Queue("default"),
		asynq.MaxRetry(3),
		asynq.Timeout(30 * time.Second),
		asynq.Retention(24 * time.Hour),
	}

	opts = append(defaultOption, opts...)

	task := asynq.NewTask(taskName, data)

	_, err := a.client.Client.Enqueue(task, opts...)
	return err
}

func (a *QueueRepo) ExecuteQueue(taskName string) error {
	return nil
}