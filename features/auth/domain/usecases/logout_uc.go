package usecases

import (
	"net/http"
	"strings"

	"github.com/RodolfoBonis/microdetect-api/core/entities"
	"github.com/RodolfoBonis/microdetect-api/core/errors"
	"github.com/RodolfoBonis/microdetect-api/core/logger"
	"github.com/gin-gonic/gin"
)

// Logout invalidates the user's refresh token and ends the session.
// @Summary User Logout
// @Schemes
// @Description Invalidate the refresh token and logout the user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer refresh token"
// @Success 200 {object} bool "Logout successful"
// @Failure 400 {object} errors.HttpError
// @Failure 401 {object} errors.HttpError
// @Failure 403 {object} errors.HttpError
// @Failure 409 {object} errors.HttpError
// @Failure 500 {object} errors.HttpError
// @Router /auth/logout [post]
// @Example request {"Authorization": "Bearer <refresh-token>"}
// @Example response true
func (uc *authUseCaseImpl) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) < 1 {
		err := errors.NewAppError(entities.ErrInvalidToken, "Missing token", nil, nil)
		httpError := err.ToHTTPError()
		uc.Logger.LogError(ctx, "Logout failed: missing token", err)
		c.AbortWithStatusJSON(httpError.StatusCode, httpError)
		c.Abort()
		return
	}
	refreshToken := strings.Split(authHeader, " ")[1]
	err := uc.KeycloakClient.Logout(
		c,
		uc.KeycloakAccessData.ClientID,
		uc.KeycloakAccessData.ClientSecret,
		uc.KeycloakAccessData.Realm,
		refreshToken,
	)
	if err != nil {
		currentError := errors.NewAppError(entities.ErrUsecase, err.Error(), nil, err)
		httpError := currentError.ToHTTPError()
		uc.Logger.LogError(ctx, "Logout failed", currentError)
		c.AbortWithStatusJSON(httpError.StatusCode, httpError)
		c.Abort()
		return
	}
	uc.Logger.Info(ctx, "Logout successful", logger.Fields{
		"ip": c.ClientIP(),
	})
	c.JSON(http.StatusOK, true)
}
