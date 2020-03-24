package main

import (
	"log"
	"os"

	"github.com/peterhoward42/kafkapubcli/internal"
)

func main() {
	connUrl := os.Getenv("KAFKA_CONN_URL")
	topic := os.Getenv("KAFKA_TOPIC")

	publisher, err := internal.NewKafkaPublisher(connUrl, topic)
	if err != nil {
		log.Fatalf("NewKafkaPublisher: %v", err)
	}
	repl := internal.NewRepl(publisher)
	repl.RunForever()
}
