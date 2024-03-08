package rabbitmq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/LuandersonFerreira/teste-golang-global/internal/models"
	"github.com/streadway/amqp"
)

func InitConsumer() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"UsersQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var user models.User

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
			err := json.NewDecoder(bytes.NewReader(d.Body)).Decode(&user)
			if err != nil {
				log.Printf("Erro ao fazer decode da mensagem da fila: %v", err)
				return
			}
			models.InsertUser(user)
		}
	}()

	fmt.Println("Successfully connected to our RabbitMQ instance")
	fmt.Println(" [*] - waiting for messages ")
	<-forever

	return nil
}
