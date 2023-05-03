package main

import "edd-sample/amqpx"

func main() {
	conn := amqpx.NewConnection()

	exchange := "inquiry_data"
	queueName := "get-data"
	route_key := "12345"
	// DECLARE 1 DIRECT EXCHANG
	conn.DeclareExchangeTopic(exchange)

	//DECLARE QUEUE_NAME
	conn.DeclareQueue(queueName)

	// BINDING QUUEUE
	conn.BindingExchangeToQueue(exchange, queueName, route_key)

	// consume
	conn.ReceiveMessage(queueName, route_key)

}
