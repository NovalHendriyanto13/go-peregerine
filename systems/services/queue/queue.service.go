package Queue

import (
	"peregerine/configs"
	QueueProvider "peregerine/systems/services/queue/provider"
	QueueClient "peregerine/systems/services/queue/types"
	QueueRepo "peregerine/systems/services/queue/repository"
)

type Container struct {
	QueueRepo QueueClient.QueueMethod
	client *QueueClient.QueueClient
}

type ContainerServer struct {
	Qs *QueueClient.QueueServer
}

func Build() (*Container, error) {
	q, err := QueueProvider.Build()
	if err != nil {
		return nil, err
	}
	repo := QueueRepo.Build(q)
	
	return &Container{ QueueRepo: repo, client: q }, nil
}

// Close releases the queue client connection
func (c *Container) Close() error {
	return c.client.Client.Close()
}

func BuildServer(cfg configs.WorkerConfig) (*ContainerServer) {
	qe := QueueProvider.BuildServer(cfg)
		
	return &ContainerServer{ Qs: qe }
}
