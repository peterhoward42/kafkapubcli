package internal

import (
	"context"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	writer *kafka.Writer
}

func NewKafkaPublisher(connUrl, topic string) *KafkaPublisher {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{connUrl},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &KafkaPublisher{writer: writer}
}

func (p *KafkaPublisher) Publish(msg []byte) error {
	kmsg := kafka.Message{
		Value: msg,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	err := p.writer.WriteMessages(ctx, kmsg)
	if err != nil {
		return fmt.Errorf("writeMessages: %v", err)
	}
	return nil
}
