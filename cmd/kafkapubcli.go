package main

import (
	"fmt"
	"log"
	"os"

	"github.com/peterhoward42/mscomp-kafka-pub-cli/internal/kafka"
)

func main() {
	host := os.Getenv("host")
	port := os.Getenv("port")
	topic := os.Getenv("topic")

	producer, err := kafka.NewProducer(host, port, topic)
	if err != nil {
		log.Fatalf("NewProducer: %v", err)
	}

	fmt.Println("Hello World")
}
