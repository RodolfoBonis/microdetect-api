package di

import (
	"github.com/RodolfoBonis/microdetect-api/core/config"
	"github.com/RodolfoBonis/microdetect-api/core/logger"
	"github.com/RodolfoBonis/microdetect-api/features/system/data/services"
	gpuService "github.com/RodolfoBonis/microdetect-api/features/system/data/services/gpu"
	"github.com/RodolfoBonis/microdetect-api/features/system/domain/usecases"
	"go.uber.org/fx"
)

// SystemModule provides the fx module for system dependencies.
var SystemModule = fx.Module("system",
	fx.Provide(
		func(logger logger.Logger, cfg *config.AppConfig) usecases.SystemUseCase {
			gpu := gpuService.NewService(logger)
			service := services.NewSystemService(gpu)
			return usecases.NewSystemUseCase(service, logger)
		},
	),
)
