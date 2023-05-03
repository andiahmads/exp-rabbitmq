package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnErorr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Transaction struct {
	ID     int
	Amount int
	Status bool
}

func main() {
	conn, err := amqp.Dial("amqp://root:endi@localhost:5672")
	failOnErorr(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnErorr(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when used
		false,   // exclusive
		false,   // no wait
		nil,     // arguments
	)
	failOnErorr(err, "failed to declare queue")

	// body := "jagung worldssss"

	for i := 0; i < 100000; i++ {
		transaction := Transaction{
			ID:     i,
			Amount: 20 + i,
			Status: false,
		}

		js, _ := json.Marshal(transaction)
		meong := string(js)
		meong += " " + fmt.Sprintf("%v", (i+1))
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(meong),
			})
		log.Printf(" [x] Sent %s", js)
	}

	failOnErorr(err, "Failed to publish a message")

}
