package errors

import (
	"github.com/RodolfoBonis/microdetect-api/core/config"
	"github.com/RodolfoBonis/microdetect-api/core/entities"
)

// swaggo:generate
type HttpError struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
	StackTrace string `json:"stack_trace,omitempty"`
}

func (e *HttpError) ToMap() map[string]interface{} {
	stackTrace := callers()
	return map[string]interface{}{
		"code":        e.StatusCode,
		"message":     e.Message,
		"stack_trace": stackTrace.String(),
	}
}

func NewHTTPError(statusCode int, message string) *HttpError {
	httpError := &HttpError{
		StatusCode: statusCode,
		Message:    message,
	}

	if config.EnvironmentConfig() == entities.Environment.Development {
		stacktrace := callers()
		httpError.StackTrace = stacktrace.String()
	}

	return httpError
}
