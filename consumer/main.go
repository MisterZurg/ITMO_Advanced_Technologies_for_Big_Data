package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	kafka "github.com/segmentio/kafka-go"
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
		Host:     "localhost",                              // ХЗ так подключает к контейнеру
		Port:     "18123",                                  // ХЗ так подключает к контейнеру
		User:     GetVar("CLICKHOUSE_USER", "LiBo"),        // LanLan
		Password: GetVar("CLICKHOUSE_PASSWORD", "YiSheng"), // XueSheng
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
		//Settings: clickhouse.Settings{
		//	"max_execution_time": 60,
		//},
		//Compression: &clickhouse.Compression{
		//	Method: clickhouse.CompressionLZ4,
		//},
		//BlockBufferSize:      10,
		//MaxCompressionBuffer: 10240,
		//TLS: &tls.Config{
		//	InsecureSkipVerify: true,
		//},
	}
	// fmt.Printf("%+v", opts)
	conn, err := clickhouse.Open(&opts)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx)
	if err != nil {
		log.Println("Cannot PING shit to ShitHouse")
		return nil, err
	}

	return &conn, nil
}

// connectToShitHouse карчое вот через SQL не яндексовский подключает
func connectToShitHouse(cfg *ClickConfig) (*sql.DB, error) {
	chDB := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: cfg.User,
			Password: cfg.Password,
		},
		Protocol: clickhouse.HTTP,
	})

	if err := chDB.Ping(); err != nil {
		log.Fatal("Cannot PING shit to ShitHouse")
		return nil, err
	}
	return chDB, nil
}

func main() {
	// ЗАКОМЕНТИЛ ДЛЯ ТЕСТА ТОЛЬКО БД
	// ctx := context.Background()

	//kafkaURL := os.Getenv("kafkaURL")
	//topic := os.Getenv("topic")
	//
	//reader := getKafkaReader(kafkaURL, topic)
	//
	//defer func() {
	//	err := reader.Close()
	//	if err != nil {
	//		log.Fatal("unable to close consumer connection to kafka: ", err)
	//	}
	//	fmt.Println("Consumer connection closed!")
	//}()

	cfg := NewClickConfig()
	fmt.Println(cfg)

	// DOES NOT WORK DO WE HAVE TO FIX THAT????
	//	_, err := connectToClickHouse(ctx, cfg)
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	// WORK
	_, err := connectToShitHouse(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// ==============================
	// МОЖНО РАСКОМЕНТИТЬ ДЛЯ ТОГО ЧТОБ УБЕДИТЬСЯ ЧТО СОЗДАЕТСЯ ТАБЛИЧКА ЛОГС
	//_, err = chDB.Exec(`
	//	CREATE TABLE IF NOT EXISTS logs
	//	(
	//		log_id              UInt32,
	//		timestamp           DateTime('Europe/Moscow'),
	//		level_of_msg        UInt8,
	//		application         String,
	//		hostname            String,
	//		logger_name         String,
	//		source              String,
	//		pid                 UInt32,
	//		line                String,
	//		error_status        UInt32,
	//		message             String
	//	) Engine = Memory`,
	//)
	//if err != nil {
	//	log.Fatal("Cannot Create Table")
	//}

	// ==============================
	// ЗАКОМЕНТИЛ ДЛЯ ТЕСТА ТОЛЬКО БД
	//fmt.Println("Start consuming!")
	//for {
	//	m, err := reader.ReadMessage(context.Background())
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	fmt.Printf("message at topic:%v partition:%v offset:%v\n%s = %s\n\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	//}
}
