package brokers

import (
	"github.com/streadway/amqp"
)

func InitRabbitMQConnection(mqConnUrl string) (*amqp.Connection, error) {
	rabbitMQConn, mqErr := amqp.Dial(mqConnUrl)
	return rabbitMQConn, mqErr
}
