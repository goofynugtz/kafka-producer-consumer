package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	KAFKA_TOPIC string
	MONGODB_URI string
	DB_NAME     string
}

func config() env {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	var e env
	e.KAFKA_TOPIC = os.Getenv("KAFKA_TOPIC")
	e.MONGODB_URI = os.Getenv("MONGODB_URI")
	e.DB_NAME = os.Getenv("DB_NAME")
	return e
}

var Env env = config()
