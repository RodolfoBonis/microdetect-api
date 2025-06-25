package services

import (
	"fmt"
	"github.com/RodolfoBonis/microdetect-api/core/errors"
	"github.com/RodolfoBonis/microdetect-api/features/system/domain/entities"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemService interface {
	GetCPUInfo() (entities.CPU, *errors.AppError)
	GetMemoryInfo() (entities.Memory, *errors.AppError)
	GetStorageInfo() (entities.Storage, *errors.AppError)
	GetHostInfo() (string, *errors.AppError)
}

type systemService struct{}

func NewSystemService() SystemService {
	return &systemService{}
}

func (s *systemService) GetCPUInfo() (entities.CPU, *errors.AppError) {
	infos, err := cpu.Info()

	if err != nil {
		return entities.CPU{}, errors.ServiceError(err.Error())
	}

	if len(infos) == 0 {
		return entities.CPU{}, errors.ServiceError("no CPU information available")
	}

	cpuInfo := infos[0]

	percent, _ := cpu.Percent(0, false)
	return entities.CPU{
		Model: cpuInfo.ModelName,
		Cores: cpuInfo.Cores,
		Usage: fmt.Sprintf("%.2f%%", percent[0]),
	}, nil
}

func (s *systemService) GetMemoryInfo() (entities.Memory, *errors.AppError) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return entities.Memory{}, errors.ServiceError(err.Error())
	}

	return entities.Memory{
		Total:      fmt.Sprintf("%.2f GB", float64(memInfo.Total)/(1024*1024*1024)),
		Available:  fmt.Sprintf("%.2f GB", float64(memInfo.Available)/(1024*1024*1024)),
		Used:       fmt.Sprintf("%.2f GB", float64(memInfo.Used)/(1024*1024*1024)),
		Percentage: fmt.Sprintf("%.2f%%", memInfo.UsedPercent),
	}, nil
}

func (s *systemService) GetStorageInfo() (entities.Storage, *errors.AppError) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return entities.Storage{}, errors.ServiceError(err.Error())
	}

	var totalUsed, totalTotal uint64
	partition := partitions[0]
	usage, err := disk.Usage(partition.Mountpoint)

	totalUsed += usage.Used
	totalTotal += usage.Total

	var usagePercentage float64
	if totalTotal > 0 {
		usagePercentage = float64(totalUsed) / float64(totalTotal) * 100
	}

	return entities.Storage{
		Used:       fmt.Sprintf("%v GB", totalUsed/(1024*1024*1024)),
		Total:      fmt.Sprintf("%v GB", totalTotal/(1024*1024*1024)),
		Percentage: fmt.Sprintf("%.2f%%", usagePercentage),
	}, nil
}

func (s *systemService) GetHostInfo() (string, *errors.AppError) {
	info, err := host.Info()

	if err != nil {
		return "", errors.ServiceError(err.Error())
	}

	return fmt.Sprintf("Platform: %s %s (%s)",
		info.Platform,
		info.PlatformVersion,
		info.PlatformFamily,
	), nil
}
