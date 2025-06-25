package di

import (
	"github.com/RodolfoBonis/microdetect-api/features/system/data/services"
	"github.com/RodolfoBonis/microdetect-api/features/system/domain/usecases"
)

func SystemInjection() usecases.SystemUseCase {
	service := services.NewSystemService()
	return usecases.NewSystemUseCase(service)
}
