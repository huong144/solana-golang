package Kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func StartKafka() {
	conf := kafka.ReaderConfig{
		Brokers: []string{"localhost:9093"},
		Topic:   "myTopic",
		GroupID: "g1",
	}

	r := kafka.NewReader(conf)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln("Loi roi ", err)
		}
		fmt.Println("message is ", string(m.Value))
	}
}
