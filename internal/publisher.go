package internal

import (
	"context"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type Publisher interface {
	Publish(msg []byte)
}

type KafkaPublisher struct {
	writer *kafka.Writer
}

func NewKafkaPublisher(connUrl, topic string) (*KafkaPublisher, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{connUrl},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &KafkaPublisher{writer: writer}, nil
}

func (p *KafkaPublisher) Publish(msg []byte) {
	kmsg := kafka.Message{
		Value: msg,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	err := p.writer.WriteMessages(ctx, kmsg)
	if err != nil {
		log.Printf("WriteMessages: %v", err)
		return
	}
	log.Printf("Message written")
}
