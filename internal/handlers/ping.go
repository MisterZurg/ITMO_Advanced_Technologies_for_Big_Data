package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
	//  *sql.Conn
)

type CreateLogRequest struct {
	LevelOfMsg  int
	Pid         int
	Source      string
	ErrorStatus int
}

func Ping10(ctx *gin.Context) {
	for i := 0; i < 10; i++ {
		Ping(ctx)
	}
}

func Ping(ctx *gin.Context) {
	req := CreateLogRequest{
		LevelOfMsg:  1,
		Pid:         42069,
		Source:      "Gin Server",
		ErrorStatus: http.StatusOK,
	}

	sqlStatement := `
			INSERT INTO logs (timestamp, level_of_msg, pid, source, error_status)
			VALUES (CURRENT_TIMESTAMP, $1, $2, $3, $4)
		`

	_, err := DB.Exec(sqlStatement, req.LevelOfMsg, req.Pid, req.Source, req.ErrorStatus)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Pong"})
}

func PingKafka(kafkaWriter *kafka.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		//body, err := ioutil.ReadAll(req.Body)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		msg := kafka.Message{
			Value: []byte("Китайская скороговорка 四是四 4 это 4!"),
		}
		err := kafkaWriter.WriteMessages(c, msg)
		if err != nil {
			// wrt.Write([]byte(err.Error()))
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Pong"})
	}
}
