package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "quickstart",
		GroupID:  "g1",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(conf)

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error", err)
			continue
		}
		fmt.Println("Message: ", string(message.Value))
	}
}
