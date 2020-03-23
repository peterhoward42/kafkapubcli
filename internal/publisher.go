package internal

type Publisher interface{}

type KafkaPublisher struct {
}

func NewKafkaPublisher(host, port, topic string) (*KafkaPublisher, error) {
	_ = []string{host, port, topic}
	return &KafkaPublisher{}, nil
}
