package configs

import "os"

var (
	app_name= getEnv("APP_NAME", "")
	app_port= getEnv("APP_PORT", "8000")
)


func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}