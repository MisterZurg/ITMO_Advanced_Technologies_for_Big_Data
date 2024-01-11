package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/ch-go/proto"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
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
		Host:     "ch-db",                                  // ХЗ так подключает к контейнеру
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

	conn, err := ch.Dial(ctx, ch.Options{
		Address:  "ch-db:9000",
		User:     "user",
		Password: "pass",
	})
	if err != nil {
		log.Fatal("dial fuckup: ", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal("ping fuckup", err)
	}

	if err := conn.Do(ctx, ch.Query{
		Body: `CREATE TABLE IF NOT EXISTS logs
		(
			log_id              UInt32,
			timestamp           DateTime64(9),
			level_of_msg        UInt8,
			application         String,
			hostname            String,
			logger_name         String,
			source              String,
			pid                 UInt32,
			line                String,
			error_status        UInt32,
			message             String
		) Engine = Memory`,
	}); err != nil {
		log.Fatal("TTTTTAAAAAAAAAAABBBBBBLLLLEEEE", err)
	}

	// Define all columns of table.
	var (
		log_id       proto.ColUInt32
		level_of_msg proto.ColUInt8
		application  proto.ColStr
		hostname     proto.ColStr
		logger_name  proto.ColStr
		source       proto.ColStr
		pid          proto.ColUInt32
		line         proto.ColStr
		error_status proto.ColUInt32
		message      proto.ColStr

		timestamp = new(proto.ColDateTime64).WithPrecision(proto.PrecisionNano)
	)

	fmt.Println("Start consuming!")
	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v\n%s = %s\n\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		var got KafkaMessage
		_ = json.Unmarshal(m.Value, &got)

		log_id.Append(uint32(got.LogID))
		level_of_msg.Append(got.LevelOfMsg)
		application.Append(string(m.Key))
		hostname.Append(got.Hostname)
		logger_name.Append(got.LoggerName)
		source.Append(got.Source)
		pid.Append(got.Pid)
		line.Append(got.Line)
		error_status.Append(got.ErrorStatus)
		message.Append(got.Message)
		timestamp.Append(got.Timestamp)
		//fmt.Printf("message at topic:%v partition:%v offset:%v\n%s = %s\n\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		// Insert single data block.
		input := proto.Input{
			{Name: "log_id", Data: &log_id},
			{Name: "level_of_msg", Data: &level_of_msg},
			{Name: "application", Data: &application},
			{Name: "hostname", Data: &hostname},
			{Name: "logger_name", Data: &logger_name},
			{Name: "source", Data: &source},
			{Name: "pid", Data: &pid},
			{Name: "line", Data: &line},
			{Name: "error_status", Data: &error_status},
			{Name: "message", Data: &message},
			{Name: "timestamp", Data: timestamp},
		}
		q := "INSERT INTO logs VALUES"
		if err := conn.Do(ctx, ch.Query{
			Body:  q,
			Input: input,
		}); err != nil {
			log.Fatal("Insert: ", err)
		}

		log_id.Reset()
		level_of_msg.Reset()
		application.Reset()
		hostname.Reset()
		logger_name.Reset()
		source.Reset()
		pid.Reset()
		line.Reset()
		error_status.Reset()
		message.Reset()
		timestamp.Reset()
	}
}
