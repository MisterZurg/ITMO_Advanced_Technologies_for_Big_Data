package main

import (
	"ITMO_Advanced_Technologies_for_Big_Data/internal/config"
	"ITMO_Advanced_Technologies_for_Big_Data/internal/handlers"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	err    error
	Cfg    *config.Config
	Db     *sqlx.DB
	Server *gin.Engine
)

func init() {
	Server = gin.Default()

	Cfg, err = config.New()
	if err != nil {
		fmt.Println("Cannot get Config")
		return
	}
	//pgDSN := config.GetPostgresConnectionString(Cfg.DBType, Cfg.DBUser, Cfg.DBPassword, Cfg.DBHost, Cfg.DBPort, Cfg.DBName)
	//Db, err = sqlx.Connect("pgx", pgDSN)
	//if err != nil {
	//	fmt.Println(pgDSN)
	//	log.Fatal("Cannot connect to Postgres")
	//}

	shitHouseDSN := config.GetClickHouseConnectionString(Cfg.DBType, Cfg.DBUser, Cfg.DBPassword, Cfg.DBHost, Cfg.DBPort, Cfg.DBName)
	Db, err := clickhouse.Open(shitHouseDSN)
	if err != nil {
		//fmt.Println(pgDSN)
		log.Fatal("Cannot connect to Postgres")
	}
	handlers.Db = Db
}

func main() {
	router := Server.Group("/api")
	router.GET("/ping", handlers.Ping)
	router.GET("/ping10", handlers.Ping10)
	log.Fatal(Server.Run(":8080"))
}
