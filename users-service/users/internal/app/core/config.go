package core

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	Database    string
	DatabaseURL string
	Broker      string
}

func NewConfig() *Config {
	port := "8000"
	DBUrl := "mongodb://localhost:27017"
	database := "userDatabase"
	broker := "amqp://admin:admin@localhost:5672/"

	if os.Getenv("USER_SERVICE_PORT") != "" {
		port = os.Getenv("USER_SERVICE_PORT")
	}

	if os.Getenv("USER_DATABASE_URL") != "" {
		DBUrl = os.Getenv("USER_DATABASE_URL")
	}

	if os.Getenv("USER_DATABASE") != "" {
		database = os.Getenv("USER_DATABASE")
	}

	if os.Getenv("RABBIT_USERNAME") != "" &&
		os.Getenv("RABBIT_PASSWORD") != "" &&
		os.Getenv("RABBIT_HOST") != "" {

		broker = fmt.Sprintf(
			"amqp://%s:%s@%s:5672/",
			os.Getenv("RABBIT_USERNAME"),
			os.Getenv("RABBIT_PASSWORD"),
			os.Getenv("RABBIT_HOST"))
	}

	return &Config{
		Port:        port,
		Database:    database,
		DatabaseURL: DBUrl,
		Broker:      broker,
	}
}
