package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key, def string) string {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(".env file not found:", err)
	}

	val := os.Getenv(key)

	if len(val) > 0 {
		return val
	}

	return def
}