package health

import (
	"net/http"

	"github.com/RodolfoBonis/microdetect-api/core/logger"
	"github.com/gin-gonic/gin"
)

func HealthRoutes(route *gin.RouterGroup, logger logger.Logger) {
	route.GET("/health_check", func(context *gin.Context) {
		logger.Info(context.Request.Context(), "Health check accessed")
		context.String(http.StatusOK, "This Service is Healthy")
	})
}

// healthCheck godoc
// @Summary Health Check
// @Schemes
// @Description Check if This service is healthy
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} errors.HttpError
// @Failure 401 {object} errors.HttpError
// @Failure 403 {object} errors.HttpError
// @Failure 409 {object} errors.HttpError
// @Failure 500 {object} errors.HttpError
// @Router /health_check [get]
func healthCheck(context *gin.Context) {
	context.String(http.StatusOK, "This Service is Healthy")
}
