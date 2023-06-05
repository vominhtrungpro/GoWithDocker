package main

import (
	"bufio"
	"fmt"
	"os"

	appKafka "github.com/vominhtrungpro/kafka"
)

func main() {
	// fmt.Println("Start!")
	// appKafka.StartKafka()
	// fmt.Println("Kafka started!")
	// time.Sleep(10 * time.Minute)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	appKafka.WriteKafka(text)
}
