package main

import (
	"ITMO_Advanced_Technologies_for_Big_Data/internal/config"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	err error
	Cfg *config.Config
	// Db     *sqlx.DB
	Server *gin.Engine
)

//func init() {
//	Server = gin.Default()
//
//	Cfg, err = config.New()
//	if err != nil {
//		fmt.Println("Cannot get Config")
//		return
//	}
//
//	// set from envs ENVS
//	chCFG := config.DatabaseConfig{
//		Name:     "",
//		Host:     "localhost",
//		Port:     "18123",
//		User:     "shittyshiter",
//		Password: "suckass",
//	}
//
//	chDB := db.ConnectToClickHouse(chCFG)
//	if err := chDB.Ping(); err != nil {
//		log.Fatal("Cannot PING shit to ShitHouse")
//	}
//
//	handlers.DB = chDB
//}

func kafkaSetUp() {

}

//var responseChannels map[string]chan *sarama.ConsumerMessage
//var mu sync.Mutex

func main() {
	// TODO: MOVE HANDLER LOGIC INTO PRODUCER
	//router := Server.Group("/api")
	//router.GET("/ping", handlers.Ping)
	//router.GET("/ping10", handlers.Ping10)
	//log.Fatal(Server.Run(":8080"))

	// responseChannels - словарь для хранения каналов ответов, индексированных по ID запроса
	// mu - мьютекс для обеспечения синхронизации доступа к словарю responseChannels
	//producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	//if err != nil {
	//	log.Fatalf("Failed to create producer: %v", err)
	//}
	//defer producer.Close()
	//
	//// Создание консьюмера Kafka
	//consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, nil)
	//if err != nil {
	//	log.Fatalf("Failed to create consumer: %v", err)
	//}
	//defer consumer.Close()
	//
	//// Подписка на партицию "pong" в Kafka
	//partConsumer, err := consumer.ConsumePartition("pong", 0, sarama.OffsetNewest)
	//if err != nil {
	//	log.Fatalf("Failed to consume partition: %v", err)
	//}
	//defer partConsumer.Close()
	//
	//// Горутина для обработки входящих сообщений от Kafka
	//go func() {
	//	for {
	//		select {
	//		// Чтение сообщения из Kafka
	//		case msg, ok := <-partConsumer.Messages():
	//			if !ok {
	//				log.Println("Channel closed, exiting goroutine")
	//				return
	//			}
	//			responseID := string(msg.Key)
	//			mu.Lock()
	//			ch, exists := responseChannels[responseID]
	//			if exists {
	//				ch <- msg
	//				delete(responseChannels, responseID)
	//			}
	//			mu.Unlock()
	//		}
	//	}
	//}()
	//
	//message := "ЗЕЛИБОБА"
	//
	//// Преобразование сообщения в JSON что бы потом отправить через kafka
	//bytes, err := json.Marshal(message)
	//if err != nil {
	//	log.Fatal("Cannot marshal")
	//}
	//
	//requestID := uuid.New().String()
	//msg := &sarama.ProducerMessage{
	//	Topic: "ping",
	//	Key:   sarama.StringEncoder(requestID),
	//	Value: sarama.ByteEncoder(bytes),
	//}
	//
	//// отправка сообщения в Kafka
	//_, _, err = producer.SendMessage(msg)
	//if err != nil {
	//	log.Fatalf("Failed to send message to Kafka: %v", err)
	//}
	//conn, err := kafka.DialLeader(context.Background(),
	//	"tcp", "localhost:9092",
	//	"my-topic-1",
	//	1,
	//)
	//if err != nil {
	//	panic(err.Error())
	//}

	//conn, err := kafka.DialLeader(context.Background(),
	//	"tcp", "localhost:9092",
	//	"my-topic-2",
	//	2,
	//)
	//if err != nil {
	//	panic(err.Error())
	//}

	//conn, err := kafka.Dial("tcp", "localhost:9092")
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer conn.Close()
	//
	//partitions, err := conn.ReadPartitions()
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//m := map[string]struct{}{}
	//
	//for _, p := range partitions {
	//	m[p.Topic] = struct{}{}
	//}
	//for k := range m {
	//	fmt.Println(k)
	//}
	//fmt.Println("H1")
	//topic := "my-topic-1"
	//partition := 1
	//
	//conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	//if err != nil {
	//	log.Fatal("failed to dial leader:", err)
	//}
	//
	//fmt.Println("H2")
	//// conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
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
	//fmt.Println("HERE")
	//return

	//topic := "my-topic"
	//partition := 0
	//
	//conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	//if err != nil {
	//	log.Fatal("failed to dial leader:", err)
	//}
	//
	//conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	//batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	//
	//b := make([]byte, 10e3) // 10KB max per message
	//for {
	//	n, err := batch.Read(b)
	//	if err != nil {
	//		break
	//	}
	//	fmt.Println("DDDDD")
	//	fmt.Println(string(b[:n]))
	//}
	//
	//if err := batch.Close(); err != nil {
	//	log.Fatal("failed to close batch:", err)
	//}
	//
	//if err := conn.Close(); err != nil {
	//	log.Fatal("failed to close connection:", err)
	//}
}
