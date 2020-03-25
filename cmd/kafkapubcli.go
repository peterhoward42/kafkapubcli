package main

import (
	"log"
	"os"

	"github.com/peterhoward42/kafkapubcli/internal"
)

func main() {
	connUrl := os.Getenv("KAFKA_CONN_URL")
	topic := os.Getenv("KAFKA_TOPIC")

	log.Printf("Creating publisher\n")
	publisher, err := internal.NewKafkaPublisher(connUrl, topic)
	if err != nil {
		log.Fatalf("NewKafkaPublisher: %v\n", err)
	}
	log.Printf("Creating repl\n")
	repl := internal.NewRepl(publisher)
	log.Printf("Launching run forever\n")

	repl.RunForever()
	log.Printf("Run forever finished, so main will exit\n")
}
