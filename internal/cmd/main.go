package main

import (
	"ITMO_Advanced_Technologies_for_Big_Data/internal/config"
	"ITMO_Advanced_Technologies_for_Big_Data/internal/handlers"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

var (
	err    error
	Cfg    *config.Config
	Db     *sql.DB
	Server *gin.Engine
)

func init() {
	Server = gin.Default()

	Cfg, err = config.New()
	if err != nil {
		fmt.Println("Cannot get Config")
		return
	}
	pgDSN := config.GetPostgresConnectionString(Cfg.DBType, Cfg.DBUser, Cfg.DBPassword, Cfg.DBHost, Cfg.DBPort, Cfg.DBName)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Cfg.DBHost, Cfg.DBPort, Cfg.DBUser, Cfg.DBPassword, Cfg.DBName)
	Db, err = sql.Open("postgres", psqlInfo)
	if nil != err {
		fmt.Println(pgDSN)
		log.Fatal("Cannot connect to Postgres")
	}
	handlers.Db = Db
}

func main() {
	router := Server.Group("/api")
	router.GET("/ping", handlers.Ping)
	log.Fatal(Server.Run(":8080"))
}
