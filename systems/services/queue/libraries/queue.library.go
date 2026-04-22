package QueueLibrary

import (
	"log"
	"fmt"
	"net"
	"time"
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

func WaitForRedis(addr string) {
	for i := 0; i < 30; i++ {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			conn.Close()
			fmt.Println("✅ Redis ready")
			return
		}

		fmt.Println("⏳ Waiting Redis...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("❌ Redis not ready")
}