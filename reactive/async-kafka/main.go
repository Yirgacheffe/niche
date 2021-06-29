package main

import (
	"fmt"
	"log"
	"net/http"

	sarama "gopkg.in/Shopify/sarama.v1"
)

func ProcessResponse(producer sarama.AsyncProducer) {
	for {
		select {
		case result := <-producer.Successes():
			log.Printf("> message: \"%s\" sent to partition %d at offset %d\n", result.Value, result.Partition, result.Offset)
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
		}
	}
}

type KafkaController struct {
	producer sarama.AsyncProducer
}

func (c *KafkaController) Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg := r.FormValue("msg")
	if msg == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("msg must be set"))
		return
	}

	c.producer.Input() <- &sarama.ProducerMessage{Topic: "example", Key: nil, Value: sarama.StringEncoder(msg)}

	w.WriteHeader(http.StatusOK)
}

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Success = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	defer producer.AsyncClose()

	go ProcessResponse(producer)

	c := KafkaController{
		producer,
	}
	http.HandleFunc("/", c.Handler)
	fmt.Println("Listening on port: 3333")
	panic(http.ListenAndServe(":3333", nil))
}
