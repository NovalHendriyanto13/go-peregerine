// Configs application
package configs

import "os"

var (
	// AppName for main application name
	AppName= getEnv("APP_NAME", "")
	// AppPort for main application port
	AppPort= getEnv("APP_PORT", "8000")
	// LLAMAHost for main AI Model Host
	LLAMAHost= getEnv("LLAMA_HOST", "")
)


func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}