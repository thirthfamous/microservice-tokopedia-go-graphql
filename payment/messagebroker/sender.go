package messagebroker

import (
	"encoding/json"
	"log"
	"thirthfamous/tokopedia-clone-go-graphql/model/message"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func SendToMessageQueue(method string, orderId int) {
	conn, err := amqp.Dial("amqp://myuser:mypassword@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	body, _ := json.Marshal(message.PaymentMessage{
		Method:  method,
		OrderId: orderId,
	})
	err = ch.Publish(
		"logs_direct", // exchange
		"payment",     // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
