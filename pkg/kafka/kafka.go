package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func ProduceMessage(topic, message string) error {
	writer := kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(message),
		},
	)
	if err != nil {
		log.Printf("could not write message %s: %v", message, err)
		return err
	}

	log.Printf("message %s written to topic %s", message, topic)
	return nil
}
