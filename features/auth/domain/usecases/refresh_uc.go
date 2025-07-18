package usecases

import (
	"net/http"
	"strings"

	coreEntities "github.com/RodolfoBonis/microdetect-api/core/entities"
	"github.com/RodolfoBonis/microdetect-api/core/errors"
	"github.com/RodolfoBonis/microdetect-api/core/logger"
	"github.com/RodolfoBonis/microdetect-api/features/auth/domain/entities"
	"github.com/gin-gonic/gin"
)

// RefreshAuthToken renews the user's authentication tokens.
// @Summary Refresh Login Access Token
// @Schemes
// @Description Refresh the user's access and refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer refresh token"
// @Success 200 {object} entities.LoginResponseEntity "Tokens refreshed"
// @Failure 400 {object} errors.HttpError
// @Failure 401 {object} errors.HttpError
// @Failure 403 {object} errors.HttpError
// @Failure 409 {object} errors.HttpError
// @Failure 500 {object} errors.HttpError
// @Router /auth/refresh [post]
// @Example request {"Authorization": "Bearer <refresh-token>"}
// @Example response {"accessToken": "jwt-token", "refreshToken": "refresh-token", "expiresIn": 3600}
func (uc *authUseCaseImpl) RefreshAuthToken(c *gin.Context) {
	ctx := c.Request.Context()
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) < 1 {
		err := errors.NewAppError(coreEntities.ErrInvalidToken, "Invalid token", nil, nil)
		httpError := err.ToHTTPError()
		uc.Logger.LogError(ctx, "Refresh failed: missing token", err)
		c.AbortWithStatusJSON(httpError.StatusCode, httpError)
		c.Abort()
		return
	}
	refreshToken := strings.Split(authHeader, " ")[1]
	token, err := uc.KeycloakClient.RefreshToken(
		c,
		refreshToken,
		uc.KeycloakAccessData.ClientID,
		uc.KeycloakAccessData.ClientSecret,
		uc.KeycloakAccessData.Realm,
	)
	if err != nil {
		currentError := errors.UsecaseError(err.Error())
		httpError := currentError.ToHTTPError()
		uc.Logger.LogError(ctx, "Refresh failed", currentError)
		c.AbortWithStatusJSON(httpError.StatusCode, httpError)
		c.Abort()
		return
	}
	uc.Logger.Info(ctx, "Token refreshed successfully", logger.Fields{
		"ip": c.ClientIP(),
	})
	c.JSON(http.StatusOK, entities.LoginResponseEntity{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
	})
}
