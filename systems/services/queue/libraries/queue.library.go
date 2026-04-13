package QueueLibrary

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"encoding/json"
)

func ConvertConfig(m fiber.Map) (asynq.Config, error) {
	var cfg asynq.Config

	data, err := json.Marshal(m)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	return cfg, err
}