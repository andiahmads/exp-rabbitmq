package main

import "edd-sample/amqpx"

func main() {
	conn := amqpx.NewConnection()

	exchange := "SEND_NOTIFICATION_PAYMENT"
	queueName := "notification"
	route_key := "12345"
	// DECLARE 1 DIRECT EXCHANGE
	conn.DeclareExchangeTopic(exchange)

	//DECLARE QUEUE_NAME
	conn.DeclareQueue(queueName)

	// BINDING QUUEUE
	conn.BindingExchangeToQueue(exchange, queueName, route_key)

	// consume
	conn.ReceiveMessage(queueName, route_key)

}
