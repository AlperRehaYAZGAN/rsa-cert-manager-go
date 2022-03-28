package brokers

import (
	"github.com/streadway/amqp"
)

var mqApp *amqp.Connection
var mqConnErr error

func InitMQConnection(mqConnString string) {
	mqApp, mqConnErr = InitRabbitMQConnection(mqConnString)
	if mqConnErr != nil {
		panic(mqConnErr)
	}
}

func GetMqConn() *amqp.Connection {
	return mqApp
}
