package QueueRepository

import (
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

func (a *QueueRepo) SetQueue(taskName string, payload interface{}) error {
	data, _ := json.Marshal(payload)

	task := asynq.NewTask(taskName, data)

	_, err := a.client.Client.Enqueue(task)
	return err
}

func (a *QueueRepo) ExecuteQueue(taskName string) error {
	return nil
}