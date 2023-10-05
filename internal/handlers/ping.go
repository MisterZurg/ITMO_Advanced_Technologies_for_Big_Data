package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

var (
	Db *sqlx.DB
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

	_, err := Db.Exec(sqlStatement, req.LevelOfMsg, req.Pid, req.Source, req.ErrorStatus)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Pong"})
}
