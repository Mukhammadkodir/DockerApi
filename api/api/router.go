package api

import (
	v1 "github/DockerApi/api/api/handler"
	"github/DockerApi/api/config"
	"github/DockerApi/api/pkg/db"
	"github/DockerApi/api/pkg/logger"
	"github/DockerApi/api/storage"
	"log"

	"github/DockerApi/api/api/docs"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
)

// @title 			Test api
// @version			1.0
// @description		This is Test User Api
// @termsOfService	http://swagger.io/terms/

// @securityDefinitions.apikey BearerAuth
// @in  header
// @name Authorization

// @contact.name	Api Support
// @contact.url		http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url		http://www.apache.org/license/LICENSE-2.0.html

// @host	localhost:8080
// @BasePath /v1

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	StorageManager storage.IStorage
}

func New(option Option) *gin.Engine {
	router := gin.New()
	cfg := config.Load()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/v1"
	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		StorageManager: option.StorageManager,
		Cfg:            option.Conf,
	}, connDB)

	api := router.Group("/v1")

	api.POST("/users", handlerV1.CreateUser)
	api.GET("/user/:id", handlerV1.Get)
	api.PUT("/user", handlerV1.UpdateUser)
	api.DELETE("/user/:id", handlerV1.DeleteUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	return router
}
