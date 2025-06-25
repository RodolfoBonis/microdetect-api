package config

import (
	"fmt"
	"os"

	"github.com/RodolfoBonis/microdetect-api/core/entities"

	"github.com/joho/godotenv"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	return defaultValue
}

func EnvPort() string {
	return GetEnv("PORT", "8000")
}

func EnvKeyCloak() entities.KeyCloakDataEntity {
	return entities.KeyCloakDataEntity{
		ClientID:     GetEnv("CLIENT_ID", "test"),
		ClientSecret: GetEnv("CLIENT_SECRET", "test"),
		Realm:        GetEnv("REALM", "test"),
		Host:         GetEnv("KEYCLOAK_HOST", "localhost"),
	}
}

func EnvServiceId() string {
	return GetEnv("SERVICE_ID", "")
}

func EnvSentryDSN() string {
	return GetEnv("SENTRY_DSN", "")
}

func EnvDBHost() string {
	return GetEnv("DB_HOST", "localhost")
}

func EnvDBPort() string {
	return GetEnv("DB_PORT", "5432")
}

func EnvDBUser() string {
	return GetEnv("DB_USER", "")
}

func EnvDBPassword() string {
	return GetEnv("DB_SECRET", "")
}

func EnvDBName() string {
	return GetEnv("DB_NAME", "")
}

func EnvironmentConfig() string {
	return GetEnv("ENV", entities.Environment.Development)
}

func EnvServiceName() string {
	return GetEnv("SERVICE_NAME", "API")
}

func envUserAmqp() string {
	return GetEnv("USER_AMQP", "guest")
}

func envPasswordAmqp() string {
	return GetEnv("PASSWORD_AMQP", "guest")
}

func envHostAmqp() string {
	return GetEnv("HOST_AMQP", "localhost:5672")
}

func EnvAmqpConnection() string {
	user := envUserAmqp()
	password := envPasswordAmqp()
	host := envHostAmqp()

	return fmt.Sprintf("amqp://%s:%s@%s/", user, password, host)
}

func LoadEnvVars() {
	env := EnvironmentConfig()
	if env == entities.Environment.Production || env == entities.Environment.Staging {
		fmt.Printf("Not using .env file in production or staging")
		return
	}

	filename := fmt.Sprintf(".env.%s", env)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = ".env"
	}

	err := godotenv.Load(filename)

	if err != nil {
		fmt.Printf(".env file not loaded")
		os.Exit(1)
	}
}
