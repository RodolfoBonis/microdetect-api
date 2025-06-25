package gpu

import (
	"github.com/RodolfoBonis/microdetect-api/core/errors"
	"github.com/RodolfoBonis/microdetect-api/features/system/domain/entities"
)

type Detector interface {
	GetGPUInfo() (entities.GPU, *errors.AppError)
}
