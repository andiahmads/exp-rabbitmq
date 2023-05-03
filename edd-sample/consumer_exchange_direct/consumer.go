package main

import (
	"edd-sample/amqpx"
)

func main() {
	conn := amqpx.NewConnection()

	ex_direct := "EXCHANGE_DIRECT"
	queueName := "order"
	route_key := "data"
	// DECLARE 1 DIRECT EXCHANGE
	conn.DeclareExchangeDirect(ex_direct)

	// BINDING QUUEUE
	conn.BindingExchangeToQueue(ex_direct, queueName, route_key)

	// consume
	conn.ReceiveMessage(queueName, route_key)
}
