package errors

import (
	"github.com/RodolfoBonis/microdetect-api/core/types"
	"runtime"
)

func callers() types.StackTrace {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(5, pcs[:])
	return pcs[0:n]
}
