package config

import (
	"os"
	"time"
)

type Config struct {
	HTTPAddr     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func Load() Config {
	return Config{
		HTTPAddr:     getEnv("HTTP_ADDR", ":8080"),
		ReadTimeout:  getDuration("READ_TIMEOUT", 5*time.Second),
		WriteTimeout: getDuration("WRITE_TIMEOUT", 10*time.Second),
		IdleTimeout:  getDuration("IDLE_TIMEOUT", 60*time.Second),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getDuration(key string, fallback time.Duration) time.Duration {
	if v := getEnv(key, ""); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return fallback
}
