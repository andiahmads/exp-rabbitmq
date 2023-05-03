package main

import (
	"edd-sample/amqpx"
	"time"
)

// mengirim queue berdasarkan route_key yang sama dengan consumer

func main() {
	conn := amqpx.NewConnection()

	ex := "inquiry_data"
	// queueName := "order"
	route_key := "12345"

	conn.DeclareExchangeTopic(ex)

	for {
		timesNow := time.Now().String()

		messagex := timesNow
		time.Sleep(2 * time.Second)
		conn.SendMessage(ex, route_key, messagex)
	}
	// SEND MESSAGE

}
