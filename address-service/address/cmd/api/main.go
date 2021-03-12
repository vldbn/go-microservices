package main

import (
	"address/internal/app/core"
	"address/internal/app/server"
	"address/internal/app/server/broker"
	"address/internal/app/service/impl"
	"address/internal/app/store/mongodb"
	"context"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	config := core.NewConfig()

	conn, err := amqp.Dial(config.Broker)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Println(err)
			panic(err)
		}
	}()

	channel, err := conn.Channel()

	defer func() {
		if err := channel.Close(); err != nil {
			log.Println(err)
			panic(err)
		}
	}()

	err = channel.ExchangeDeclare(
		"users",
		"topic",
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("Declare Queue")
	if err := broker.QueueDeclare(channel, config.Queues); err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("Bind Queues")
	if err := broker.QueueBind(channel, config.Queues); err != nil {
		log.Println(err)
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DatabaseURL))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	log.Println("Try ping db")

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	store := mongodb.NewMongoStore(client, config.Database)
	serv := impl.NewServices(store)

	if err := server.StartServer(config, serv, channel); err != nil {
		panic(err)
	}
}
