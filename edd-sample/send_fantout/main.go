package main

import "edd-sample/amqpx"

// mengirim queue berdasarkan route_key yang sama dengan consumer

func main() {
	conn := amqpx.NewConnection()

	exchange := "EXCHANGE_FANOUT"

	// DECLARE 1 DIRECT EXCHANGE
	conn.DeclareExchangeFanout(exchange)

	// SEND MESSAGE
	conn.SendMessage(exchange, "", "fanout test")

}
