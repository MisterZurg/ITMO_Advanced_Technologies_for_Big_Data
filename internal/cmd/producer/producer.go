package main

import (
	"ITMO_Advanced_Technologies_for_Big_Data/internal/config"
	"ITMO_Advanced_Technologies_for_Big_Data/internal/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"log"
)

var (
	err error
	Cfg *config.Config
	// Db     *sqlx.DB
	Server *gin.Engine
)

func init() {
	Server = gin.Default()

	Cfg, err = config.New()
	if err != nil {
		fmt.Println("Cannot get Config")
		return
	}

	// set from envs ENVS
	//chCFG := config.DatabaseConfig{
	//	Name:     "",
	//	Host:     "localhost",
	//	Port:     "18123",
	//	User:     "shittyshiter",
	//	Password: "suckass",
	//}

	//chDB := db.ConnectToClickHouse(chCFG)
	//if err := chDB.Ping(); err != nil {
	//	log.Fatal("Cannot PING shit to ShitHouse")
	//}
	//
	//handlers.DB = chDB
}

// https://github.com/yusufsyaifudin/go-kafka-example/blob/accadecfca65c956bf03a6199ff1f3944bc6d7dc/cmd/api/main.go
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func main() {
	topic := "my-topic"
	//partition := 0
	kw := getKafkaWriter("localhost:9092", topic)
	//
	//conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	//if err != nil {
	//	log.Fatal("failed to dial leader:", err)
	//}

	//conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	//_, err = conn.WriteMessages(
	//	kafka.Message{Value: []byte("one!")},
	//	kafka.Message{Value: []byte("two!")},
	//	kafka.Message{Value: []byte("three!")},
	//)
	//if err != nil {
	//	log.Fatal("failed to write messages:", err)
	//}
	//
	//if err := conn.Close(); err != nil {
	//	log.Fatal("failed to close writer:", err)
	//}
	//
	router := Server.Group("/api")
	router.GET("/ping", handlers.PingKafka(kw))
	//router.GET("/ping10", handlers.Ping10)
	log.Fatal(Server.Run(":8080"))
}
