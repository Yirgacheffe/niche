package main

import (
	"log"

	sarama "gopkg.in/Shopify/sarama.v1"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	partitionConsmer, err := consumer.ConsumePartition("example", 0, consumer.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer partitionConsmer.Close()

	for {
		msg := <-partitionConsmer.Message()
		log.Printf("Consumed message: \"%s\" at offset: %d\n", msg.Value, msg.Offset)
	}

}
