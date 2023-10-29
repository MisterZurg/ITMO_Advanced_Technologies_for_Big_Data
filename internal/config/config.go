package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
)

type Config struct {
	DBUser     string `env:"POSTGRES_USER" envDefault:"postgres"`
	DBPassword string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	DBName     string `env:"POSTGRES_DB" envDefault:"db"`
	DBHost     string `env:"POSTGRES_HOST" envDefault:"localhost"` // "localhost"` 192.168.0.108
	DBPort     string `env:"POSTGRES_PORT" envDefault:"5432"`
	DBType     string `env:"DBTYPE" envDefault:"postgresql"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func GetPostgresConnectionString(DBType, DBUser, DBPassword, DBHost, DBPort, DBName string) string {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		DBType, DBUser, DBPassword, DBHost, DBPort, DBName)
	return dsn
}

func GetClickHouseConnectionString(DBType, DBUser, DBPassword, DBHost, DBPort, DBName string) string {
	// dial_timeout - a duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix such as "300ms", "1s". Valid time units are "ms", "s", "m". (default 30s)
	// connection_open_strategy - round_robin/in_order
	// clickhouse://username:password@host1:9000,host2:9000/database?dial_timeout=200ms&max_execution_time=60
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?dial_timeout=200ms&max_execution_time=60",
		DBType, DBUser, DBPassword, DBHost, DBPort,
	)
	return dsn
}
