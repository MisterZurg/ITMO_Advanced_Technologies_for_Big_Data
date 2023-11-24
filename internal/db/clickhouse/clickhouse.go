package db

import (
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/jmoiron/sqlx"

	"ITMO_Advanced_Technologies_for_Big_Data/internal/config"
)

func makeAddr(cfg config.DatabaseConfig) string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

func ConnectToClickHouse(cfg config.DatabaseConfig) *sqlx.DB {
	addr := makeAddr(cfg)
	return sqlx.NewDb(clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{addr},
		Auth: clickhouse.Auth{
			Database: cfg.Name,
			Username: cfg.User,
			Password: cfg.Password,
		},
		//TLS: &tls.Config{
		//	InsecureSkipVerify: true,
		//},
		//Settings: clickhouse.Settings{
		//	"max_execution_time": 60,
		//},
		//DialTimeout: time.Second * 30,
		//Compression: &clickhouse.Compression{
		//	Method: clickhouse.CompressionLZ4,
		//},
		//Debug:                true,
		//BlockBufferSize:      10,
		//MaxCompressionBuffer: 10240,
		//ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
		//	Products: []struct {
		//		Name    string
		//		Version string
		//	}{
		//		{Name: "my-app", Version: "0.1"},
		//	},
		//},
	}), "clickhouse")
}
