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
	timeout := os.Getenv("KAFKA_TIMEOUT")
	pumpInterval := os.Getenv("PUMP_INTERVAL")

	timeoutDuration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalf("Could not parse this timeout: <%s>, with error: %v", timeout, err)
	}
	tick, err := time.ParseDuration(pumpInterval)
	if err != nil {
		log.Fatalf("Could not parse this pumpInterval: <%s>, with error: %v", pumpInterval, err)
	}

	publisher := internal.NewKafkaPublisher(connURL, topic, timeoutDuration)

	ticker := time.NewTicker(tick)
	for range ticker.C {
		msg := []byte(fmt.Sprintf("msg at: %v", time.Now()))
		err := publisher.Publish(msg)
		if err != nil {
			log.Printf("Publish: %v", err)
			continue
		}
		log.Print("Message published")
	}

	log.Println("Application terminating")
}
