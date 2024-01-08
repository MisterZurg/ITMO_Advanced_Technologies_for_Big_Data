package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
)

func PingKafka(conn *kafka.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		msg := kafka.Message{
			Value: []byte("Китайская скороговорка 四是四 4 это 4!"),
		}
		_, err := conn.WriteMessages(msg)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Pong"})
	}
}
