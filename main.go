package main

import (
	"fmt"
	"time"

	appKafka "github.com/vominhtrungpro/kafka"
)

func main() {
	fmt.Println("Start!")
	appKafka.StartKafka()
	fmt.Println("Kafka started!")
	time.Sleep(10 * time.Minute)
}
