package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)


func consumer(){
	
	//* Create a new consumer instance and create configuration
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "KAFKA_BROKER",
		"group.id": "myGroup",
		"auto.offset.reset": "earliest",
	})


	if err != nil {
		panic(err)
	}

	//* If nothing error, subscribe to the topic
	err = c.SubscribeTopics([]string{"myTopic", "^Regex.*[Tt]opic"}, nil)
	
	if err != nil {
		panic(err)
	}

	//* A signal handler or similar could be used to set this to false to break the loop
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)

		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			//* You can do whatever you want with the message here, like writing to a file or processing it further.
		} else if !err.(kafka.Error).IsTimeout(){
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, err.(kafka.Error).Code())
		}
	}

	c.Close()


}