package src

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Consumer() {
	// Crear un nuevo consumidor
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "KAFKA_BROKER",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Suscribirse al tópico
	err = c.SubscribeTopics([]string{"myTopic"}, nil)
	if err != nil {
		panic(err)
	}

	run := true
	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Mensaje recibido en %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Error en el consumidor: %v (%v)\n", err, err.(kafka.Error).Code())
		}
	}
}
