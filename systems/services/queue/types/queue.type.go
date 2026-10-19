package QueueType

import (
	"context"
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
	SetQueue(ctx context.Context, taskName string, payload interface{}, opts ...asynq.Option) error
}

type QueueExecuteMethod interface {
	ExecuteQueue(taskName string) error
}
