package app

import (
	"github.com/RodolfoBonis/microdetect-api/core/config"
	"github.com/RodolfoBonis/microdetect-api/core/logger"
	"github.com/RodolfoBonis/microdetect-api/core/middlewares"
	"github.com/RodolfoBonis/microdetect-api/core/services"
	"github.com/RodolfoBonis/microdetect-api/features/auth/di"
	auth_uc "github.com/RodolfoBonis/microdetect-api/features/auth/domain/usecases"
	systemdi "github.com/RodolfoBonis/microdetect-api/features/system/di"
	system_uc "github.com/RodolfoBonis/microdetect-api/features/system/domain/usecases"
	"github.com/RodolfoBonis/microdetect-api/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// NewFxApp cria e retorna uma nova instância da aplicação Fx.
func NewFxApp() *fx.App {
	return fx.New(
		logger.Module,
		config.Module,
		middlewares.Module,
		di.AuthModule,
		systemdi.SystemModule,
		fx.Provide(
			gin.New,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, router *gin.Engine, systemUc system_uc.SystemUseCase, authUc auth_uc.AuthUseCase, monitoring *middlewares.MonitoringMiddleware, protectFactory func(handler gin.HandlerFunc, role string) gin.HandlerFunc, logger logger.Logger) {
				routes.InitializeRoutes(router, systemUc, authUc, monitoring, protectFactory, logger)
				RegisterHooks(lc, router, logger, monitoring)
				_ = services.OpenConnection(logger)
			},
		),
	)
}
