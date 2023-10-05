package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
)

type Config struct {
	DBUser     string `env:"POSTGRES_USER" envDefault:"postgres"`
	DBPassword string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	DBName     string `env:"POSTGRES_DB" envDefault:"db"`
	DBHost     string `env:"POSTGRES_HOST" envDefault:"localhost"` // "localhost"`
	DBPort     string `env:"POSTGRES_PORT" envDefault:"5432"`
	DBType     string `env:"DBTYPE" envDefault:"postgres"`
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
