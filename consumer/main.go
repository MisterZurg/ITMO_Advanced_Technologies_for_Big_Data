package main

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"os"
	"strings"
)

func GetVar(key, def string) string {
	val := os.Getenv(key)
	if len(val) > 0 {
		return val
	}
	return def
}

type ClickConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

func NewClickConfig() *ClickConfig {
	cfg := ClickConfig{
		Host:     "ch-db",
		Port:     "8123",
		User:     GetVar("CLICKHOUSE_USER", "LanLan"),
		Password: GetVar("CLICKHOUSE_PASSWORD", "XueSheng"),
	}
	return &cfg
}

func getKafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Partition: 0,
		Topic:     topic,
		MinBytes:  1,
		MaxBytes:  10e6, // 10MB
	})
}

func connectToClickHouse(ctx context.Context, cfg *ClickConfig) (*driver.Conn, error) {
	opts := clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: cfg.User,
			Password: cfg.Password,
		},
		//Debug: true,
		//Debugf: func(format string, v ...any) {
		//	fmt.Printf(format+"\n", v...)
		//},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
	}
	fmt.Printf("%+v", opts)
	conn, err := clickhouse.Open(&opts)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &conn, nil
}

func main() {
	ctx := context.Background()

	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")

	reader := getKafkaReader(kafkaURL, topic)

	defer func() {
		err := reader.Close()
		if err != nil {
			log.Fatal("unable to close consumer connection to kafka: ", err)
		}
		fmt.Println("Consumer connection closed!")
	}()

	cfg := NewClickConfig()
	_, err := connectToClickHouse(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start consuming!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v\n%s = %s\n\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
