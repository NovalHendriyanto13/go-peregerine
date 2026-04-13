package QueueType

import (
	"github.com/hibiken/asynq"
)

type QueueClient struct {
	Client *asynq.Client
}

type QueueServer struct {
	Server *asynq.Server
}

type QueueServeMux struct {
	Mux *asynq.ServeMux
}

type QueueMethod interface {
	SetQueue(taskName string, payload interface{}) error
}

type QueueExecuteMethod interface {
	ExecuteQueue(taskName string) error
}
