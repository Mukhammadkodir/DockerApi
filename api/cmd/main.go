package main

import (
	"github/DockerApi/api/api"
	"github/DockerApi/api/config"
	"github/DockerApi/api/pkg/logger"
)

func main() {

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api")

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
