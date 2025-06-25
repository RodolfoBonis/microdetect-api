package app

import (
	"os"
	"strings"

	"github.com/RodolfoBonis/microdetect-api/core/config"
	"github.com/RodolfoBonis/microdetect-api/core/entities"
	"github.com/RodolfoBonis/microdetect-api/core/services"
	"github.com/RodolfoBonis/microdetect-api/docs"
)

func init() {
	config.LoadEnvVars()
	services.InitializeOAuthServer()

	// Use this for open connection with DataBase
	//appError := services.OpenConnection()
	//if appError != nil {
	//	logger.Log.Error(appError.Message, appError.ToMap())
	//	panic(appError)
	//}

	// Use this for Run Yours migrations
	// services.RunMigrations()

	// Use this for open connection with RabbitMQ
	// services.StartAmqpConnection()

	docs.SwaggerInfo.Title = "microdetect-api"
	docs.SwaggerInfo.Description = "API for YOLO training management"
	versionFileName := "version.txt"
	if config.EnvironmentConfig() == entities.Environment.Production {
		versionFileName = "/version.txt"
	}

	version := "unknown"
	if content, err := os.ReadFile(versionFileName); err == nil {
		version = strings.TrimSpace(string(content))
	}
	host := "localhost"

	if config.EnvironmentConfig() == entities.Environment.Production {
		host = "microdetect-api.rodolfodebonis.com.br"
	}

	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Version = version
	scheme := "http"

	if config.EnvironmentConfig() == entities.Environment.Production {
		scheme = "https"
	}

	docs.SwaggerInfo.Schemes = []string{scheme}
}
