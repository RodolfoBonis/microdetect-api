package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RodolfoBonis/microdetect-api/core/config"
	"github.com/RodolfoBonis/microdetect-api/core/errors"
	"github.com/RodolfoBonis/microdetect-api/core/logger"
	"github.com/RodolfoBonis/microdetect-api/core/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// RegisterHooks registers application lifecycle hooks.
func RegisterHooks(lifecycle fx.Lifecycle, router *gin.Engine, logger logger.Logger, monitoring *middlewares.MonitoringMiddleware) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				err := router.SetTrustedProxies([]string{})
				if err != nil {
					appError := errors.RootError(err.Error(), nil)
					logger.LogError(ctx, "Erro ao configurar trusted proxies", appError)
					panic(err)
				}
				config.SentryConfig()
				router.Use(monitoring.SentryMiddleware())
				router.Use(monitoring.LogMiddleware)
				router.Use(gin.Logger())
				router.Use(gin.Recovery())
				router.Use(gin.ErrorLogger())
				runPort := fmt.Sprintf(":%s", config.EnvPort())
				go func() {
					err = router.Run(runPort)
					if err != nil && err != http.ErrServerClosed {
						appError := errors.RootError(err.Error(), nil)
						logger.LogError(ctx, "Erro ao subir servidor HTTP", appError)
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger.Info(ctx, "Stopping server.")
				return nil
			},
		},
	)
}
