package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/segmentio/kafka-go"
)

func Ping(kafkaWriter *kafka.Writer) func(http.ResponseWriter, *http.Request) {
	return func(wrt http.ResponseWriter, req *http.Request) {
		msg := kafka.Message{
			Key:   []byte("сы шы сы"),
			Value: []byte("bu nihao"),
		}
		err := kafkaWriter.WriteMessages(req.Context(), msg)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func main() {
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")

	kafkaWriter := getKafkaWriter(kafkaURL, topic)

	defer func() {
		err := kafkaWriter.Close()
		if err != nil {
			fmt.Println("error closing producer: ", err)
			return
		}
		fmt.Println("producer closed")
	}()

	http.HandleFunc("/", Ping(kafkaWriter))

	fmt.Println("producer-api started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
