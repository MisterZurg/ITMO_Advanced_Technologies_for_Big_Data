package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaMessage struct {
	LogID       int32     `json:"log_id"`
	Timestamp   time.Time `json:"timestamp"`
	LevelOfMsg  uint8     `json:"level_of_msg"`
	Application string    `json:"application"`
	Hostname    string    `json:"hostname"`
	LoggerName  string    `json:"logger_name"`
	Source      string    `json:"source"`
	Pid         uint32    `json:"pid"`
	Line        string    `json:"line"`
	ErrorStatus uint32    `json:"error_status"`
	Message     string    `json:"message"`
}

func Ping(kafkaWriter *kafka.Writer, id string) func(http.ResponseWriter, *http.Request) {
	return func(wrt http.ResponseWriter, req *http.Request) {
		marshaled, _ := json.Marshal(KafkaMessage{
			LogID:       int32(rand.Int()),
			Timestamp:   time.Now(),
			LevelOfMsg:  1,
			Application: "producer",
			Hostname:    "big.data.com",
			LoggerName:  "big.data.logger",
			Source:      "producer/main.go",
			Pid:         uint32(os.Getpid()),
			Line:        "41",
			ErrorStatus: 200,
		})

		msg := kafka.Message{
			Key:   []byte("RNG:" + strconv.Itoa(rand.Int())),
			Value: []byte(marshaled),
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
