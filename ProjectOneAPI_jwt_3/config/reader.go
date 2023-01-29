package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func EnvReader(str string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error .env")
	}

	val := os.Getenv(str)
	return val
}
