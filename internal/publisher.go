package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	writer          *kafka.Writer
	timeoutDuration time.Duration
}

func NewKafkaPublisher(connUrl, topic string, timeout string) *KafkaPublisher {
	timeoutDuration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalf("Cannot parse this timeout string: <%s>, with error: %v", timeout, err)
	}
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{connUrl},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &KafkaPublisher{
		writer:          writer,
		timeoutDuration: timeoutDuration,
	}
}

func (p *KafkaPublisher) Publish(msg []byte) error {
	kmsg := kafka.Message{
		Value: msg,
	}
	ctx, cancel := context.WithTimeout(context.Background(), p.timeoutDuration)
	defer cancel()
	err := p.writer.WriteMessages(ctx, kmsg)
	if err != nil {
		return fmt.Errorf("writeMessages: %v", err)
	}
	return nil
}
