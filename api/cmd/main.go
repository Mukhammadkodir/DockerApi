package main

import (
	"fmt"
	"github/DockerApi/api/api"
	"github/DockerApi/api/config"
	"github/DockerApi/api/pkg/logger"
	"os"

	"github.com/jmoiron/sqlx"
)

func main() {
	addr := os.Getenv("DB")
	fmt.Println("Postgres addr: " + addr)

	_, err := sqlx.Connect("postgres", addr)

	if err != nil {
		fmt.Println("Could not connect...")
	} else {
		fmt.Println("Connecting successful")
	}

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api")

	server := api.New(api.Option{
		Conf:   cfg,
		Logger: log,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
