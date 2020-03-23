package main

import (
	"log"
	"os"

	"github.com/peterhoward42/kafkapubcli/internal"
)

func main() {
	host := os.Getenv("host")
	port := os.Getenv("port")
	topic := os.Getenv("topic")

	publisher, err := internal.NewKafkaPublisher(host, port, topic)
	if err != nil {
		log.Fatalf("NewProducer: %v", err)
	}
	repl := internal.NewRepl(publisher)
	repl.RunForever()
}
