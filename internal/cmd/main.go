package main

import (
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"

	"ITMO_Advanced_Technologies_for_Big_Data/internal/config"
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
	//pgDSN := config.GetPostgresConnectionString(Cfg.DBType, Cfg.DBUser, Cfg.DBPassword, Cfg.DBHost, Cfg.DBPort, Cfg.DBName)
	//Db, err = sqlx.Connect("pgx", pgDSN)
	//if err != nil {
	//	fmt.Println(pgDSN)
	//	log.Fatal("Cannot connect to Postgres")
	//}

	// shitHouseDSN := config.GetClickHouseConnectionString(Cfg.DBType, Cfg.DBUser, Cfg.DBPassword, Cfg.DBHost, Cfg.DBPort, Cfg.DBName)
	//chCFG := config.DatabaseConfig{
	//	Name:     "",
	//	Host:     "localhost",
	//	Port:     "18123",
	//	User:     "shittyshiter",
	//	Password: "suckass",
	//}

	// chDB := db.ConnectToClickHouse(chCFG)
	// connStr := fmt.Sprintf("clickhouse://%s:%s/?user=%s&password=%s", chCFG.Host, chCFG.Port, chCFG.User, chCFG.Password)
	// connStr := "clickhouse://localhost:18123/?user=shittyshiter&password=suckass"
	//chDB, err := sqlx.Connect("clickhouse", connStr)
	//if err != nil {
	//	fmt.Println(connStr)
	//	log.Fatal("Cannot connect to ShitHouse")
	//}
	//if err := chDB.Ping(); err != nil {
	//	log.Fatal("Cannot PING shit to ShitHouse")
	//}
	chDB := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"localhost:18123"},
		Auth: clickhouse.Auth{
			Database: "",
			Username: "shittyshiter",
			Password: "suckass",
		},
		Protocol: clickhouse.HTTP,
	})
	if err := chDB.Ping(); err != nil {
		log.Fatal("Cannot PING shit to ShitHouse")
	}
	fmt.Println("ДРОЧИТЬ")
	/*
		docker run --rm -e CLICKHOUSE_DB=my_database \
		-e CLICKHOUSE_USER=username \
		-e CLICKHOUSE_DEFAULT_ACCESS_MANAGEMENT=1 \
		-e CLICKHOUSE_PASSWORD=password \
		-p 9000:9000/tcp clickhouse/clickhouse-server
	*/

	//_, err = chDB.Exec(`
	//	CREATE TABLE IF NOT EXISTS example (
	//		  Col1 UInt8
	//		, Col2 String
	//		, Col3 FixedString(3)
	//		, Col4 UUID
	//		, Col5 Map(String, UInt8)
	//		, Col6 Array(String)
	//		, Col7 Tuple(String, UInt8, Array(Map(String, String)))
	//		, Col8 DateTime
	//	) Engine = Memory
	//`)
	//if err != nil {
	//	log.Fatal("Cannot Shit in ShitHouse")
	//}

	//dsn := "tcp://127.0.0.1:18123?debug=true"
	//db, err := sql.Open("clickhouse", dsn)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//// Create the "users" table
	//_, err = db.Exec(`
	//	CREATE TABLE IF NOT EXISTS users (
	//		id UInt64,
	//		name String,
	//		age UInt8
	//	) ENGINE = Memory
	//`)
	//if err != nil {
	//	panic(err)
	//}

	// handlers.DB = chDB
}

func main() {
	router := Server.Group("/api")
	router.GET("/ping", handlers.Ping)
	router.GET("/ping10", handlers.Ping10)
	log.Fatal(Server.Run(":8080"))
}
