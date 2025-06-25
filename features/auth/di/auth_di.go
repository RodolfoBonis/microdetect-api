package di

import (
	"github.com/RodolfoBonis/microdetect-api/core/config"
	"github.com/RodolfoBonis/microdetect-api/core/services"
	"github.com/RodolfoBonis/microdetect-api/features/auth/domain/usecases"
)

func AuthInjection() usecases.AuthUseCase {
	return usecases.AuthUseCase{
		KeycloakClient:     services.AuthClient,
		KeycloakAccessData: config.EnvKeyCloak(),
	}
}
