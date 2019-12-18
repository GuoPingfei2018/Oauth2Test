/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  send
 * @Version: 1.0.0
 * @Date: 2019/12/12 10:49
 */

package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.Publish(
		"Test",  // exchange
		"Agent", // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			MessageId:   string(1),
			Type:        "AgentJob",
			Body:        []byte("Hello world"),
		})
	log.Println("send ok")
	failOnError(err, "Failed to publish a message")
}