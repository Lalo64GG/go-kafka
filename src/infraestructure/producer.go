package infraestructure

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Producer(message string) error {
	// Crear un nuevo productor
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "34.207.197.203:9092",
	})

	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Manejar eventos de entrega de mensajes
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Entrega fallida: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Mensaje entregado a %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Enviar mensajes al tópico
	topic := "myTopic"

	err = p.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{
		Topic:     &topic,
		Partition: kafka.PartitionAny},
		Value: []byte(message)},
		nil,
	)

	if err != nil {
		fmt.Printf("Error al enviar mensaje: %v\n", err)
	}

	// Asegurarse de que los mensajes se envíen antes de finalizar
	p.Flush(15 * 1000)

	return nil
}
