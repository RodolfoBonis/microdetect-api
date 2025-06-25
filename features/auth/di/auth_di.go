package di

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/RodolfoBonis/microdetect-api/core/config"
	"github.com/RodolfoBonis/microdetect-api/core/logger"
	"github.com/RodolfoBonis/microdetect-api/core/services"
	"github.com/RodolfoBonis/microdetect-api/features/auth/domain/usecases"
	"go.uber.org/fx"
)

// AuthModule provides the fx module for authentication dependencies.
var AuthModule = fx.Module("auth",
	fx.Provide(
		func(cfg *config.AppConfig) *gocloak.GoCloak {
			return services.NewAuthClient(cfg)
		},
		func(authClient *gocloak.GoCloak, logger logger.Logger) usecases.AuthUseCase {
			return usecases.NewAuthUseCase(authClient, config.EnvKeyCloak(), logger)
		},
	),
)
