package main

import (
	"ITMO_Advanced_Technologies_for_Big_Data/internal/config"
	"ITMO_Advanced_Technologies_for_Big_Data/internal/repository"
	"fmt"
)

func main() {
	//fmt.Println("Здравствуйте!")
	cfg, err := config.New()
	if err != nil {
		fmt.Println("Cannot get Config")
		return
	}

	// TODO db connect
	pgDSN := config.GetPostgresConnectionString(cfg.DBType, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	_, err = repository.NewRepository(pgDSN)
	if err != nil {
		panic("error parsing config")
	} else {
		fmt.Println("Connected tho")
	}
}
