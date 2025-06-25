package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/RodolfoBonis/microdetect-api/core/logger"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MonitoringMiddleware struct {
	logger logger.Logger
}

func NewMonitoringMiddleware(logger logger.Logger) *MonitoringMiddleware {
	return &MonitoringMiddleware{logger: logger}
}

func (m *MonitoringMiddleware) SentryMiddleware() gin.HandlerFunc {
	return sentrygin.New(sentrygin.Options{Repanic: true})
}

func (m *MonitoringMiddleware) LogMiddleware(ctx *gin.Context) {
	start := time.Now()
	requestId := uuid.NewString()
	ctx.Set("requestID", requestId)
	var responseBody = logger.HandleResponseBody(ctx.Writer)
	var requestBody = logger.
		HandleRequestBody(ctx.Request)
	ctx.Writer = responseBody

	// Adiciona o IP ao contexto do request
	ctxWithIP := context.WithValue(ctx.Request.Context(), "ip", ctx.ClientIP())
	ctx.Request = ctx.Request.WithContext(ctxWithIP)

	m.logger.Info(ctx.Request.Context(), "Request started", logger.Fields{
		"request_id":   requestId,
		"ip":           ctx.ClientIP(),
		"method":       ctx.Request.Method,
		"url":          ctx.Request.URL.String(),
		"user_agent":   ctx.Request.UserAgent(),
		"request_body": requestBody,
	})

	ctx.Next()

	latency := time.Since(start)
	status := ctx.Writer.Status()
	logFields := logger.Fields{
		"request_id":    requestId,
		"ip":            ctx.ClientIP(),
		"method":        ctx.Request.Method,
		"url":           ctx.Request.URL.String(),
		"user_agent":    ctx.Request.UserAgent(),
		"status":        status,
		"latency_ms":    latency.Milliseconds(),
		"request_body":  requestBody,
		"response_body": responseBody.Body.String(),
	}

	if isSuccessStatusCode(status) {
		m.logger.Info(ctx.Request.Context(), "Request completed", logFields)
	} else {
		m.logger.Error(ctx.Request.Context(), "Request failed", logFields)
	}
}

func isSuccessStatusCode(statusCode int) bool {
	switch statusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent:
		return true
	default:
		return false
	}
}
