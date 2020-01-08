package utils

import "os"

// GetEnv gets an ENV-variable or returns a default value
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
