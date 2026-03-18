// Package Configs is all main configuration of application
package configs

import (
	"os"
	"strconv"
)

var (
	// AppName for main application name
	AppName= getEnv("APP_NAME", "")
	// AppPort for main application port
	AppPort= getEnv("APP_PORT", "8000")
	// LLAMAHost for main AI Model Host
	LLAMAHost= getEnv("LLAMA_HOST", "")
	// RedisHost to define host of redis
	RedisHost= getEnv("REDIS_HOST", "")
	// RedisPort to define port of redis
	RedisPort= getEnv("REDIS_PORT", "6379")
	// RedisPassword to define password credential of redis
	RedisPassword= getEnv("REDIS_PASSWORD", "")
	// RedisDB to define host of redis
	RedisDB = getEnvInt("REDIS_DB", 0)
	JWTSecret = getEnvByte("JWT_SECRET", []byte{})
	JWTTTL = getEnvInt("JWT_TTL", 15)
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

func getEnvByte(key string, fallback []byte) []byte {
	return []byte(getEnv(key, string(fallback)))
}