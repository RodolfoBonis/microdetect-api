package config

import (
	"github.com/RodolfoBonis/microdetect-api/core/entities"
	"go.uber.org/fx"
)

type AppConfig struct {
	Port           string
	Keycloak       entities.KeyCloakDataEntity
	ServiceId      string
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

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Port:           EnvPort(),
		Keycloak:       EnvKeyCloak(),
		ServiceId:      EnvServiceId(),
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

var Module = fx.Module("config", fx.Provide(NewAppConfig))
