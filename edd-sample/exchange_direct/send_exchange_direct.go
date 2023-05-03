package main

import "edd-sample/amqpx"

// mengirim queue berdasarkan route_key yang sama dengan consumer

func main() {
	conn := amqpx.NewConnection()

	ex_direct := "EXCHANGE_DIRECT"
	queueName := "order"
	route_key := "meta"
	// DECLARE 1 DIRECT EXCHANGE
	conn.DeclareExchangeDirect(ex_direct)

	//DECLARE QUEUE_NAME
	conn.DeclareQueue(queueName)

	// BINDING QUUEUE
	conn.BindingExchangeToQueue(ex_direct, queueName, route_key)

	// SEND MESSAGE
	conn.SendMessage(ex_direct, route_key, "@testinf")

}
