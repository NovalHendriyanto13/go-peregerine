package configs

type WorkerConfig struct {
	Concurrency int
	Queues map[string]int
}