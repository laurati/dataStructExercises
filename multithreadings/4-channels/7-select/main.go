package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0
	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg // jogo mensagem no meu canal
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	// o loop infinito roda sempre simulando que ta sempre recebendo as mensagens do kafka e rabitMQ
	// hora recebe de um, hora recebe de outro
	for {
		// a informação chega de vários lugares e é processada
		select {
		case msg := <-c1: // rabbitmq
			fmt.Printf("Received from RabbitMQ: ID: %d - %s\n", msg.id, msg.Msg)

		case msg := <-c2: // kafka
			fmt.Printf("Received from Kafka: ID: %d - %s\n", msg.id, msg.Msg)

		case <-time.After(time.Second * 3):
			println("timeout")

		}
	}
}
