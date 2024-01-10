package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
)

//type msg struct {
//	log_id              UInt32
//	timestamp           DateTime64(9)
//	level_of_msg        UInt8
//	application         String
//	hostname            String
//	logger_name         String
//	source              String
//	pid                 UInt32
//	line                String
//	error_status        UInt32
//	message             String
//}

func Ping(kafkaWriter *kafka.Writer, id string) func(http.ResponseWriter, *http.Request) {
	return func(wrt http.ResponseWriter, req *http.Request) {
		//jsonMsg := json.Marshal()
		msg := kafka.Message{
			Key:   []byte("RNG:" + strconv.Itoa(rand.Int())),
			Value: []byte("here we json.Marshal" + id),
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
	thisName := strconv.Itoa(rand.Int())
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

	http.HandleFunc("/", Ping(kafkaWriter, thisName))

	fmt.Println("producer-api started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
