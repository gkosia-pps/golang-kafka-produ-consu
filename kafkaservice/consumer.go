package kafkaservice

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func Consume_from_kafka(num_of_messages int) {

	conf := kafka.ReaderConfig{
		Brokers:     []string{"localhost:9094"},
		Topic:       "data-from-console",
		GroupID:     "goLang-app",
		StartOffset: kafka.FirstOffset,
	}

	consumer := kafka.NewReader(conf)

	for {
		msg, err := consumer.ReadMessage(context.Background())

		if err != nil {
			fmt.Printf("Failed to read message: error %v", err)
			continue
		}

		fmt.Printf("Consumed messge: %v", string(msg.Value))

		num_of_messages -= 1
		if num_of_messages == 0 {
			break
		}

		time.Sleep(time.Millisecond * 500)
	}
}
