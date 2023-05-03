package main

import "edd-sample/amqpx"

func main() {
	conn := amqpx.NewConnection()

	exchange := "EXCHANGE_FANOUT"
	queueName := "QUEUE_NAME_ONE"
	route_key := "ROUTE_KEY"
	// DECLARE 1 DIRECT EXCHANGE
	conn.DeclareExchangeFanout(exchange)

	//DECLARE QUEUE_NAME
	conn.DeclareQueue(queueName)

	// BINDING QUUEUE
	conn.BindingExchangeToQueue(exchange, queueName, "")

	// consume
	conn.ReceiveMessage(queueName, route_key)

}
