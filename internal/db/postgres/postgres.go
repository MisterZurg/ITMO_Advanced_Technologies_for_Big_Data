package db

import (
	"ITMO_Advanced_Technologies_for_Big_Data/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

func getPostgresConnectionString(DBType, DBUser, DBPassword, DBHost, DBPort, DBName string) string {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		DBType, DBUser, DBPassword, DBHost, DBPort, DBName)
	return dsn
}

func ConnectToPostgres(cfg config.DatabaseConfig) *sqlx.DB {
	pgDSN := getPostgresConnectionString("postgresql", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	pgDB, err := sqlx.Connect("pgx", pgDSN)
	if err != nil {
		fmt.Println(pgDSN)
		log.Fatal("Cannot connect to Postgres")
	}

	return pgDB
}
