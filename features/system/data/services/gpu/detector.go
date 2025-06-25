package gpu

import (
	"github.com/RodolfoBonis/microdetect-api/core/errors"
	"github.com/RodolfoBonis/microdetect-api/features/system/domain/entities"
)

// Detector provides GPU detection capabilities.
type Detector interface {
	GetGPUInfo() (entities.GPU, *errors.AppError)
}
