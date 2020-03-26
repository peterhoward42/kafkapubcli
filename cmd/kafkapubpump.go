package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/peterhoward42/kafkapubcli/internal"
)

func main() {
	log.Println("Application starting")

	connURL := os.Getenv("KAFKA_CONN_URL")
	topic := os.Getenv("KAFKA_TOPIC")
	publisher := internal.NewKafkaPublisher(connURL, topic)

	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		msg := []byte(fmt.Sprintf("msg at: %v", time.Now()))
		err := publisher.Publish(msg)
		if err != nil {
			log.Printf("Publish: %v", err)
		}
	}

	log.Println("Application terminating")
}
