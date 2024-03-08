package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/LuandersonFerreira/teste-golang-global/internal/models"
	"github.com/streadway/amqp"
)

func ConnectQueue() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"UsersQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(q)

	return nil
}

func SendToQueue(user models.User) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ch.Close()

	jsonBody, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Erro ao codificar JSON: %v", err)
		return err
	}

	err = ch.Publish(
		"",
		"UsersQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		},
	)
	if err != nil {
		fmt.Printf("Erro ao publicar mensagem: %v", err)
		return err
	}

	fmt.Println("Mensagem enviada com sucesso.")
	return nil
}
