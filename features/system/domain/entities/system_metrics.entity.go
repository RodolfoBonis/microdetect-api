package entities

import (
	"time"
)

type SystemMetrics struct {
	Timestamp     time.Time
	CPUPercent    float64
	MemoryPercent float64
	GPUMetrics    string
}
