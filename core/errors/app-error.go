package errors

import (
	"net/http"

	"github.com/RodolfoBonis/microdetect-api/core/entities"
)

// Error is the base interface for all custom errors in the system.
type Error interface {
	error
	Code() int
	Message() string
	StackTrace() string
	Context() map[string]interface{}
	Unwrap() error
	ToLogFields() map[string]interface{}
	ToHttpError() *HttpError
}

// AppError representa um erro de aplicação padronizado.
type AppError struct {
	Type    entities.AppErrorType
	Message string
	Fields  map[string]interface{}
	Cause   error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) HTTPStatus() int {
	if status, ok := entities.AppErrorTypeToHTTP[e.Type]; ok {
		return status
	}
	return http.StatusInternalServerError
}

// NewAppError cria um novo erro padronizado.
func NewAppError(errType entities.AppErrorType, msg string, fields map[string]interface{}, cause error) *AppError {
	if msg == "" {
		msg = entities.AppErrorTypeToString[errType]
	}
	return &AppError{
		Type:    errType,
		Message: msg,
		Fields:  fields,
		Cause:   cause,
	}
}

// ToLogFields returns a map with all error details for structured logging.
func (e *AppError) ToLogFields() map[string]interface{} {
	fields := map[string]interface{}{
		"error_code":    e.Type,
		"error_message": e.Message,
	}
	for k, v := range e.Fields {
		fields[k] = v
	}
	if e.Cause != nil {
		fields["cause"] = e.Cause.Error()
	}
	return fields
}

// ToHttpError converts the AppError to an HttpError.
func (e *AppError) ToHttpError() *HttpError {
	return NewHTTPError(e.HTTPStatus(), e.Message)
}

// Helpers for common errors
func EntityError(message string, ctx ...map[string]interface{}) *AppError {
	return NewAppError(entities.ErrEntity, message, ctx[0], nil)
}
func EnvironmentError(message string, ctx ...map[string]interface{}) *AppError {
	return NewAppError(entities.ErrEnvironment, message, ctx[0], nil)
}
func MiddlewareError(message string, ctx ...map[string]interface{}) *AppError {
	return NewAppError(entities.ErrMiddleware, message, ctx[0], nil)
}
func ModelError(message string, ctx ...map[string]interface{}) *AppError {
	return NewAppError(entities.ErrModel, message, ctx[0], nil)
}
func RepositoryError(message string, ctx ...map[string]interface{}) *AppError {
	return NewAppError(entities.ErrRepository, message, ctx[0], nil)
}
func RootError(message string, ctx ...map[string]interface{}) *AppError {
	return NewAppError(entities.ErrRoot, message, ctx[0], nil)
}
func ServiceError(message string, ctx ...map[string]interface{}) *AppError {
	return NewAppError(entities.ErrService, message, ctx[0], nil)
}
func UsecaseError(message string, ctx ...map[string]interface{}) *AppError {
	return NewAppError(entities.ErrUsecase, message, ctx[0], nil)
}
