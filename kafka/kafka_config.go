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

func WriteKafka(input string) {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "quickstart",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte(input),
		},
	)
	if err != nil {
		fmt.Println("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		fmt.Println("failed to close writer:", err)
	}
}
