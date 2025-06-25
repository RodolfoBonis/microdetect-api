package auth

import (
	"github.com/RodolfoBonis/microdetect-api/features/auth/domain/usecases"
	"github.com/gin-gonic/gin"
)

// LoginHandler handles user login requests.
// @Summary User Login
// @Schemes
// @Description Authenticate user and return JWT tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body entities.RequestLoginEntity true "Login credentials"
// @Success 200 {object} entities.LoginResponseEntity "Successful login"
// @Failure 400 {object} errors.HttpError
// @Failure 401 {object} errors.HttpError
// @Failure 403 {object} errors.HttpError
// @Failure 409 {object} errors.HttpError
// @Failure 500 {object} errors.HttpError
// @Router /auth/login [post]
// @Example request {"email": "user@example.com", "password": "string"}
// @Example response {"accessToken": "jwt-token", "refreshToken": "refresh-token", "expiresIn": 3600}
func LoginHandler(authUc usecases.AuthUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authUc.ValidateLogin(c)
	}
}

// Routes registers authentication routes for the application.
func Routes(route *gin.RouterGroup, authUC usecases.AuthUseCase, protectFactory func(handler gin.HandlerFunc, role string) gin.HandlerFunc) {
	route.POST("/login", LoginHandler(authUC))
	route.POST("/logout", protectFactory(authUC.Logout, "user"))
	route.POST("/refresh", protectFactory(authUC.RefreshAuthToken, "user"))
}
