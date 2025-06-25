package gpu

import (
	"runtime"

	coreEntities "github.com/RodolfoBonis/microdetect-api/core/entities"
	"github.com/RodolfoBonis/microdetect-api/core/errors"
	"github.com/RodolfoBonis/microdetect-api/core/logger"
	"github.com/RodolfoBonis/microdetect-api/features/system/domain/entities"
)

// Service provides GPU service operations.
type Service interface {
	GetGPUInfo() (entities.GPU, *errors.AppError)
}

type service struct {
	detectors []Detector
}

// NewService creates a new Service instance.
func NewService(logger logger.Logger) Service {
	var detectors []Detector

	switch runtime.GOOS {
	case "darwin":
		detectors = append(detectors, NewMacOSDetector(logger))
	case "linux":
		detectors = append(detectors, NewNvidiaDetector(logger))
		detectors = append(detectors, NewLinuxDetector(logger))
	case "windows":
		detectors = append(detectors, NewNvidiaDetector(logger))
		detectors = append(detectors, NewWindowsDetector(logger))
	}

	return &service{detectors: detectors}
}

func (s *service) GetGPUInfo() (entities.GPU, *errors.AppError) {
	for _, detector := range s.detectors {
		gpuInfo, err := detector.GetGPUInfo()
		if err == nil {
			return gpuInfo, nil
		}
	}
	return entities.GPU{}, errors.NewAppError(coreEntities.ErrService, "No GPU detected", nil, nil)
}
