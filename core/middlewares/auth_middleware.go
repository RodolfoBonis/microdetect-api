package middlewares

import (
	"encoding/json"
	"strings"

	"github.com/RodolfoBonis/microdetect-api/core/config"
	"github.com/RodolfoBonis/microdetect-api/core/entities"
	"github.com/RodolfoBonis/microdetect-api/core/errors"
	"github.com/RodolfoBonis/microdetect-api/core/logger"
	"github.com/RodolfoBonis/microdetect-api/core/services"
	"github.com/gin-gonic/gin"

	jsonToken "github.com/golang-jwt/jwt/v4"
)

// NewProtectMiddleware creates a new authentication middleware.
func NewProtectMiddleware(logger logger.Logger) func(handler gin.HandlerFunc, role string) gin.HandlerFunc {
	return func(handler gin.HandlerFunc, role string) gin.HandlerFunc {
		return func(c *gin.Context) {
			ctx := c.Request.Context()
			requestID, _ := c.Get("requestID")
			keycloakDataAccess := config.EnvKeyCloak()
			authHeader := c.GetHeader("Authorization")

			if len(authHeader) < 1 {
				err := errors.NewAppError(entities.ErrInvalidToken, "Missing token", nil, nil)
				httpError := err.ToHTTPError()
				logger.LogError(ctx, "Auth failed: missing token", err)
				c.AbortWithStatusJSON(httpError.StatusCode, httpError)
				c.Abort()
				return
			}

			accessToken := strings.Split(authHeader, " ")[1]

			rptResult, err := services.AuthClient.RetrospectToken(
				c,
				accessToken,
				keycloakDataAccess.ClientID,
				keycloakDataAccess.ClientSecret,
				keycloakDataAccess.Realm,
			)

			if err != nil {
				appError := errors.NewAppError(entities.ErrMiddleware, err.Error(), nil, err)
				httpError := appError.ToHTTPError()
				logger.LogError(ctx, "Auth failed: token introspection error", appError)
				c.AbortWithStatusJSON(httpError.StatusCode, httpError)
				c.Abort()
				return
			}

			isTokenValid := *rptResult.Active

			if !isTokenValid {
				err := errors.NewAppError(entities.ErrInvalidToken, "Token invalid", nil, nil)
				httpError := err.ToHTTPError()
				logger.LogError(ctx, "Auth failed: token invalid", err)
				c.AbortWithStatusJSON(httpError.StatusCode, httpError)
				c.Abort()
				return
			}

			token, _, err := services.AuthClient.DecodeAccessToken(
				c,
				accessToken,
				keycloakDataAccess.Realm,
			)

			if err != nil {
				appError := errors.NewAppError(entities.ErrMiddleware, err.Error(), nil, err)
				httpError := appError.ToHTTPError()
				logger.LogError(ctx, "Auth failed: decode token error", appError)
				c.AbortWithStatusJSON(httpError.StatusCode, httpError)
				c.Abort()
				return
			}

			claims := token.Claims.(jsonToken.MapClaims)

			jsonData, _ := json.Marshal(claims)

			var userClaim entities.JWTClaim

			err = json.Unmarshal(jsonData, &userClaim)
			if err != nil {
				appError := errors.NewAppError(entities.ErrMiddleware, err.Error(), nil, err)
				httpError := appError.ToHTTPError()
				logger.LogError(ctx, "Auth failed: unmarshal claims error", appError)
				c.AbortWithStatusJSON(httpError.StatusCode, httpError)
				c.Abort()
				return
			}

			keyCloakData := config.EnvKeyCloak()
			client := userClaim.ResourceAccess[keyCloakData.ClientID].(map[string]interface{})
			rolesBytes, err := json.Marshal(client["roles"])
			err = json.Unmarshal(rolesBytes, &userClaim.Roles)
			if err != nil {
				appError := errors.NewAppError(entities.ErrMiddleware, err.Error(), nil, err)
				httpError := appError.ToHTTPError()
				logger.LogError(ctx, "Auth failed: unmarshal roles error", appError)
				c.AbortWithStatusJSON(httpError.StatusCode, httpError)
				c.Abort()
				return
			}

			containsRole := userClaim.Roles.Contains(role)

			if !containsRole {
				appError := errors.NewAppError(entities.ErrUnauthorized, "Missing required role", nil, nil)
				httpError := appError.ToHTTPError()
				logger.LogError(ctx, "Auth failed: missing required role", appError)
				c.AbortWithStatusJSON(httpError.StatusCode, httpError)
				c.Abort()
				return
			}

			logger.Info(ctx, "Auth success", map[string]interface{}{
				"request_id": requestID,
				"ip":         c.ClientIP(),
				"role":       role,
				"user_roles": userClaim.Roles,
				"user_id":    userClaim.ID,
				"email":      userClaim.Email,
			})

			c.Set("claims", userClaim)
			handler(c)
		}
	}
}
