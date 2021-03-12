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
	Queues      map[string]string
}

func NewConfig() *Config {
	port := "8001"
	DBUrl := "mongodb://localhost:27017"
	database := "addressDatabase"
	broker := "amqp://admin:admin@localhost:5672/"

	if os.Getenv("ADDRESS_SERVICE_PORT") != "" {
		port = os.Getenv("ADDRESS_SERVICE_PORT")
	}

	if os.Getenv("ADDRESS_DATABASE_URL") != "" {
		DBUrl = os.Getenv("ADDRESS_DATABASE_URL")
	}

	if os.Getenv("ADDRESS_DATABASE") != "" {
		database = os.Getenv("ADDRESS_DATABASE")
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

	q := map[string]string{"create": "users.create", "update": "users.update", "delete": "users.delete"}

	return &Config{
		Port:        port,
		Database:    database,
		DatabaseURL: DBUrl,
		Broker:      broker,
		Queues:      q,
	}
}
