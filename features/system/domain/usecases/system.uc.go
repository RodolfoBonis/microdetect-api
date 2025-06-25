package usecases

import (
	"github.com/RodolfoBonis/microdetect-api/features/system/data/services"
	"github.com/RodolfoBonis/microdetect-api/features/system/domain/entities"
	"github.com/gin-gonic/gin"
)

type SystemUseCase interface {
	GetStorage(c *gin.Context)
	GetMemory(c *gin.Context)
	GetSystemStatus(c *gin.Context)
	GetCPU(c *gin.Context)
}

type SystemUseCaseImpl struct {
	Service services.SystemService
}

func NewSystemUseCase(service services.SystemService) SystemUseCase {
	return &SystemUseCaseImpl{
		Service: service,
	}
}

func (uc *SystemUseCaseImpl) GetStorage(c *gin.Context) {
	storage, appError := uc.Service.GetStorageInfo()
	if appError != nil {
		httpError := appError.ToHttpError()
		c.JSON(httpError.StatusCode, httpError.ToMap())
		return
	}
	c.JSON(200, storage)
}

func (uc *SystemUseCaseImpl) GetMemory(c *gin.Context) {
	memory, appError := uc.Service.GetMemoryInfo()
	if appError != nil {
		httpError := appError.ToHttpError()
		c.JSON(httpError.StatusCode, httpError.ToMap())
		return
	}
	c.JSON(200, memory)
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

func (uc *SystemUseCaseImpl) GetCPU(c *gin.Context) {
	cpu, appError := uc.Service.GetCPUInfo()
	if appError != nil {
		httpError := appError.ToHttpError()
		c.JSON(httpError.StatusCode, httpError.ToMap())
		return
	}
	c.JSON(200, cpu)
}
