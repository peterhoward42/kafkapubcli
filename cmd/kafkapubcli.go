package main

import (
	"fmt"
	"log"
	"os"

	"github.com/peterhoward42/kafkapubcli/internal/pkg/kafka"
)

func main() {
	host := os.Getenv("host")
	port := os.Getenv("port")
	topic := os.Getenv("topic")

	_, err := kafka.NewProducer(host, port, topic)
	if err != nil {
		log.Fatalf("NewProducer: %v", err)
	}

	fmt.Println("Hello World")
}
