package config

import (
	"github.com/RodolfoBonis/microdetect-api/core/entities"
	"go.uber.org/fx"
)

// AppConfig holds the application configuration.
type AppConfig struct {
	Port     string
	Keycloak entities.KeyCloakDataEntity
	// ServiceID is the unique identifier for the service.
	ServiceID      string
	SentryDSN      string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	Environment    string
	ServiceName    string
	AmqpConnection string
}

// NewAppConfig creates and returns a new AppConfig instance.
func NewAppConfig() *AppConfig {
	return &AppConfig{
		Port:           EnvPort(),
		Keycloak:       EnvKeyCloak(),
		ServiceID:      EnvServiceID(),
		SentryDSN:      EnvSentryDSN(),
		DBHost:         EnvDBHost(),
		DBPort:         EnvDBPort(),
		DBUser:         EnvDBUser(),
		DBPassword:     EnvDBPassword(),
		DBName:         EnvDBName(),
		Environment:    EnvironmentConfig(),
		ServiceName:    EnvServiceName(),
		AmqpConnection: EnvAmqpConnection(),
	}
}

// Module provides the fx module for AppConfig.
var Module = fx.Module("config", fx.Provide(NewAppConfig))
