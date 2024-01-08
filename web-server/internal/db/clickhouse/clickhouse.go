package db

import (
	"ITMO_Advanced_Technologies_for_Big_Data/web-server/internal/config"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/jmoiron/sqlx"
)

func makeAddr(cfg config.DatabaseConfig) string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

func ConnectToClickHouse(cfg config.DatabaseConfig) *sqlx.DB {
	chAddr := makeAddr(cfg)
	chDB := sqlx.NewDb(clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{chAddr},
		Auth: clickhouse.Auth{
			Database: "",
			Username: cfg.User,
			Password: cfg.Password,
		},
		Protocol: clickhouse.HTTP,
	}), "clickhouse")
	if err := chDB.Ping(); err != nil {
		log.Fatal("Cannot PING shit to ShitHouse")
	}
	return chDB
}
