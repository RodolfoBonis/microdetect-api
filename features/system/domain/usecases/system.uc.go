package usecases

import (
	"github.com/RodolfoBonis/microdetect-api/features/system/data/services"
	"github.com/RodolfoBonis/microdetect-api/features/system/domain/entities"
	"github.com/gin-gonic/gin"
)

type SystemUseCase interface {
	GetSystemStatus(c *gin.Context)
}

type SystemUseCaseImpl struct {
	Service services.SystemService
}

func NewSystemUseCase(service services.SystemService) SystemUseCase {
	return &SystemUseCaseImpl{
		Service: service,
	}
}

func (uc *SystemUseCaseImpl) GetSystemStatus(c *gin.Context) {
	systemStatus := entities.SystemStatus{}

	cpu, appError := uc.Service.GetCPUInfo()
	if appError != nil {
		httpError := appError.ToHttpError()
		c.JSON(httpError.StatusCode, httpError.ToMap())
		return
	}

	memory, appError := uc.Service.GetMemoryInfo()

	if appError != nil {
		httpError := appError.ToHttpError()
		c.JSON(httpError.StatusCode, httpError.ToMap())
		return
	}

	storage, appError := uc.Service.GetStorageInfo()

	if appError != nil {
		httpError := appError.ToHttpError()
		c.JSON(httpError.StatusCode, httpError.ToMap())
		return
	}

	hostInfo, appError := uc.Service.GetHostInfo()

	if appError != nil {
		httpError := appError.ToHttpError()
		c.JSON(httpError.StatusCode, httpError.ToMap())
		return
	}

	systemStatus.OS = hostInfo
	systemStatus.CPU = cpu
	systemStatus.Memory = memory
	systemStatus.Storage = storage

	c.JSON(200, systemStatus)
}
