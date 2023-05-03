package main

import (
	"design-pattern/rabbitmq_workerpool/config"
	"design-pattern/rabbitmq_workerpool/pool"
	"log"

	"github.com/streadway/amqp"
)

func failOnErorr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	cfg := config.Get()
	conn, err := amqp.Dial("amqp://root:endi@localhost:5672/")
	failOnErorr(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnErorr(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		cfg.QueueName, // name
		false,         // durable
		false,         //delete when used
		false,         // exclusive
		false,         // no-wait
		nil,           //arguments
	)
	failOnErorr(err, "Failed to declare queue")

	msgs, err := ch.Consume(
		"hello", // queue name
		q.Name,  // consumer
		true,    // auto-ack,
		false,   // exclusive
		false,   // no-local
		false,   // no wait
		nil,     // arg
	)

	failOnErorr(err, "Failed to register a consumer")

	forever := make(chan bool)

	p := pool.NewPool(int32(cfg.Numworkers))

	p.Run()

	go func() {
		var counter int32 = 1

		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
		counter++
	}()

	log.Printf(" [*] Waiting for message. To Exit press CMD+C")

	<-forever

}
