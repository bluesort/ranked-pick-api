package env

import (
	"os"
	"strconv"
)

func GetRequiredString(key string) string {
	envVal := os.Getenv(key)
	if envVal == "" {
		panic("missing required environment variable: " + key)
	}
	return envVal
}

func GetString(key string, fallback string) string {
	envVal := os.Getenv(key)
	if envVal != "" {
		return envVal
	}
	return fallback
}

func GetRequiredBool(key string) bool {
	envVal := os.Getenv(key)
	if envVal == "" {
		panic("missing required environment variable: " + key)
	}
	return envVal == "true"
}

func GetBool(key string, fallback bool) bool {
	envVal := os.Getenv(key)
	if envVal != "" {
		return envVal == "true"
	}
	return fallback
}

func GetRequiredInt(key string) int {
	envVal := os.Getenv(key)
	if envVal == "" {
		panic("missing required environment variable: " + key)
	}
	i, err := strconv.Atoi(envVal)
	if err != nil {
		panic("invalid environment variable: " + key)
	}
	return i
}

func GetInt(key string, fallback int) int {
	envVal := os.Getenv(key)
	if envVal != "" {
		i, err := strconv.Atoi(envVal)
		if err == nil {
			return i
		}
	}
	return fallback
}
