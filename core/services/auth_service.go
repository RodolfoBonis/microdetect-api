package services

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/RodolfoBonis/microdetect-api/core/config"
)

// AuthClient is the client used for authentication operations.
var AuthClient *gocloak.GoCloak

// InitializeOAuthServer initializes the OAuth server for authentication.
func InitializeOAuthServer() {
	keycloakDataAccess := config.EnvKeyCloak()

	AuthClient = gocloak.NewClient(keycloakDataAccess.Host)
}

// NewAuthClient creates a new authentication client.
func NewAuthClient(cfg *config.AppConfig) *gocloak.GoCloak {
	return gocloak.NewClient(cfg.Keycloak.Host)
}
