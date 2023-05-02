package common

import (
	"github.com/streadway/amqp"
	"log"
)

var RabbitMq *amqp.Connection

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func CloseRabbitMq() {
	RabbitMq.Close()
}

func GetRabbitMq() *amqp.Connection {
	return RabbitMq
}

func InitRabbitMq() {
	var err error
	RabbitMq, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
}
