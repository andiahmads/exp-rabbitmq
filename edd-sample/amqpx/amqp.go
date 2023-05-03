package amqpx

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ interface {
	DeclareExchangeDirect(ExchangeName string)
	DeclareExchangeFanout(ExchangeName string)
	DeclareExchangeTopic(ExchangeName string)
	DeclareQueue(queueName string)
	BindingExchangeToQueue(ExchangeName string, queueName string, route string)
	SendMessage(exchangeName string, route string, message string)
	ReceiveMessage(queueName string, route string)
}

type rabbitMQ struct {
	connection *amqp.Connection
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func NewConnection() *rabbitMQ {
	uri := "amqp://root:endi@localhost:5672"
	fmt.Println(uri)

	if uri == "" {
		log.Fatal("RabbitMQ URI is invalid")
	}

	connection, err := amqp.Dial(uri)
	failOnError(err, "Failed to connect to RabbitMQ")

	log.Println("connected to RabbitMQ")

	return &rabbitMQ{
		connection,
	}
}

func (rabbitmq *rabbitMQ) DeclareExchangeDirect(exchangeName string) {
	ch, err := rabbitmq.connection.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		exchangeName,
		amqp.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare exchange")

	log.Printf("RabbitMQ: %s exchange created", exchangeName)
}

func (rabbitmq *rabbitMQ) DeclareExchangeFanout(exchangeName string) {
	ch, err := rabbitmq.connection.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		exchangeName,
		amqp.ExchangeFanout,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare exchange")

	log.Printf("RabbitMQ: %s exchange created", exchangeName)
}

func (rabbitmq *rabbitMQ) DeclareExchangeTopic(exchangeName string) {
	ch, err := rabbitmq.connection.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		exchangeName,
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare exchange")
	log.Printf("RabbitMQ: %s exchange created", exchangeName)
}

func (rabbitmq *rabbitMQ) DeclareQueue(queueName string) {
	ch, err := rabbitmq.connection.Channel()
	failOnError(err, "Failed to open a channel")

	_, err = ch.QueueDeclare(
		queueName, true, false, false, false, nil,
	)
	failOnError(err, "Failed to declare a queue")
}

func (rabbitmq *rabbitMQ) BindingExchangeToQueue(exchangeName string, queueName string, route string) {
	ch, err := rabbitmq.connection.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.QueueBind(
		queueName,
		route,
		exchangeName,
		false,
		nil,
	)

	failOnError(err, "Failed to binding queue with exchange")

	log.Printf("RabbitMQ: Binding %s to %s successful", exchangeName, queueName)

}

func (rabbitmq *rabbitMQ) SendMessage(exchangeName string, route string, message string) {
	ch, err := rabbitmq.connection.Channel()
	failOnError(err, "Failed to open a channel")

	content := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}

	err = ch.Publish(
		exchangeName,
		route,
		false,
		false,
		content,
	)

	failOnError(err, "Failed to publish content")

	log.Printf("RabbitMQ: published message to %s", exchangeName)
}

func (rabbitmq *rabbitMQ) ReceiveMessage(queueName string, route string) {
	ch, err := rabbitmq.connection.Channel()

	rabbitmq.DeclareQueue(queueName)

	failOnError(err, "Failed to open a channel")
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register consumer")
	log.Printf("RabbitMQ: Consumer created - %s -> %s", queueName, route)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("RabbitMQ: received message %s - %s : %s", queueName, route, d.Body)

		}

	}()

	<-forever
}

// func Cron(times string) {
// 	// set scheduler berdasarkan zona waktu sesuai kebutuhan

// 	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
// 	scheduler := cron.New(cron.WithLocation(jakartaTime))

// 	scheduler.AddFunc(times, func() { SendAutomail("New Year") })

// 	go scheduler.Start()

// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
// 	<-sig
// }

// func SendAutomail(automailType string) {

// 	// ... instruksi untuk mengirim automail berdasarkan automailType
// 	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " SendAutomail " + automailType + " telah dijalankan.\n")

// }
