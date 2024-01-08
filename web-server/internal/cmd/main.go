package main

import (
	"ITMO_Advanced_Technologies_for_Big_Data/web-server/internal/handlers"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}()

	server := gin.Default()
	router := server.Group("/api")
	router.GET("/ping", handlers.PingKafka(conn))
	log.Fatal(server.Run(":8080"))
}
