package kafkaservice

import (
	"context"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce_to_kafka(num_of_messages int) {
	conf := kafka.WriterConfig{
		Brokers: []string{"localhost:9094"},
		Topic:   "data-from-console",
	}

	producer := kafka.NewWriter(conf)

	for {

		if num_of_messages == 0 {
			break
		}

		producer.WriteMessages(context.Background(), kafka.Message{
			Value: []byte(strconv.Itoa(num_of_messages)),
		})

		num_of_messages -= 1
		time.Sleep(time.Millisecond * 500)
	}
}
