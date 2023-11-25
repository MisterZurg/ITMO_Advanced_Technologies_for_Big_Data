package main

import (
	"fmt"
	"log"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"

	"ITMO_Advanced_Technologies_for_Big_Data/internal/config"
	db "ITMO_Advanced_Technologies_for_Big_Data/internal/db/clickhouse"
	"ITMO_Advanced_Technologies_for_Big_Data/internal/handlers"
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
	chCFG := config.DatabaseConfig{
		Name:     "",
		Host:     "localhost",
		Port:     "18123",
		User:     "shittyshiter",
		Password: "suckass",
	}

	chDB := db.ConnectToClickHouse(chCFG)
	if err := chDB.Ping(); err != nil {
		log.Fatal("Cannot PING shit to ShitHouse")
	}

	handlers.DB = chDB
}

func main() {
	router := Server.Group("/api")
	router.GET("/ping", handlers.Ping)
	router.GET("/ping10", handlers.Ping10)
	log.Fatal(Server.Run(":8080"))
}
