package kafka

//func WriteMassege(ctx *gin.Context) {
//	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
//
//	_, err := conn.WriteMessages(
//		kafka.Message{Value: []byte("one!")},
//		kafka.Message{Value: []byte("two!")},
//		kafka.Message{Value: []byte("three!")},
//	)
//	if err != nil {
//		log.Fatal("failed to write messages:", err)
//	}
//
//	if err := conn.Close(); err != nil {
//		log.Fatal("failed to close writer:", err)
//	}
//}
//
//func Write10Masseges(conn *kafka.Conn) {
//	for i := 0; i < 10; i++ {
//		WriteMassege(conn)
//	}
//}
