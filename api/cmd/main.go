package main

import (
	"github/DockerApi/api/api"
	"github/DockerApi/api/config"
	"github/DockerApi/api/pkg/logger"
)

func main() {

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api-gateway")

	// serviceManager, err := services.NewServiceManager(&cfg)
	// if err != nil {
	// 	log.Error("gRPC dial error", logger.Error(err))
	// }

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
