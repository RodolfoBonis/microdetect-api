package routes

import (
	"github.com/RodolfoBonis/microdetect-api/core/health"
	"github.com/RodolfoBonis/microdetect-api/core/logger"
	"github.com/RodolfoBonis/microdetect-api/core/middlewares"
	"github.com/RodolfoBonis/microdetect-api/features/auth"
	auth_uc "github.com/RodolfoBonis/microdetect-api/features/auth/domain/usecases"
	"github.com/RodolfoBonis/microdetect-api/features/system"
	system_uc "github.com/RodolfoBonis/microdetect-api/features/system/domain/usecases"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(
	router *gin.Engine,
	systemUc system_uc.SystemUseCase,
	authUc auth_uc.AuthUseCase,
	monitoring *middlewares.MonitoringMiddleware,
	protectFactory func(handler gin.HandlerFunc, role string) gin.HandlerFunc,
	logger logger.Logger,
) {

	root := router.Group("/v1")

	root.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	health.HealthRoutes(root, logger)
	auth.AuthRoutes(root, authUc, protectFactory)
	system.SystemRoutes(root, systemUc)
}
