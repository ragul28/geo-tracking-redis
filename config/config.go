package config

import (
	"os"
)

type AppVar struct {
	Port  string
	Env   string
	Redis RedisVar
}

type RedisVar struct {
	RdHost     string
	RdPassword string
}

func GetEnv() *AppVar {
	var c AppVar

	// App
	c.Port = getEnv("PORT", "8081")
	c.Env = getEnv("ENV", "dev")

	// redis connection vars
	c.Redis.RdHost = getEnv("REDIS_HOST", "localhost:6379")
	c.Redis.RdPassword = getEnv("REDIS_PASSWORD", "")

	return &c
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
