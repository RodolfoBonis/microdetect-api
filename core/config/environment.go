package config

import (
	"fmt"
	"os"

	"github.com/RodolfoBonis/microdetect-api/core/entities"

	"github.com/joho/godotenv"
)

// GetEnv retrieves the value of the specified environment variable.
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	return defaultValue
}

// EnvPort returns the port from environment variables.
func EnvPort() string {
	return GetEnv("PORT", "8000")
}

// EnvKeyCloak returns the Keycloak configuration from environment variables.
func EnvKeyCloak() entities.KeyCloakDataEntity {
	return entities.KeyCloakDataEntity{
		ClientID:     GetEnv("CLIENT_ID", "test"),
		ClientSecret: GetEnv("CLIENT_SECRET", "test"),
		Realm:        GetEnv("REALM", "test"),
		Host:         GetEnv("KEYCLOAK_HOST", "localhost"),
	}
}

// EnvServiceID retrieves the service ID from the environment variables.
func EnvServiceID() string {
	return GetEnv("SERVICE_ID", "")
}

// EnvSentryDSN returns the Sentry DSN from environment variables.
func EnvSentryDSN() string {
	return GetEnv("SENTRY_DSN", "")
}

// EnvDBHost returns the database host from environment variables.
func EnvDBHost() string {
	return GetEnv("DB_HOST", "localhost")
}

// EnvDBPort returns the database port from environment variables.
func EnvDBPort() string {
	return GetEnv("DB_PORT", "5432")
}

// EnvDBUser returns the database user from environment variables.
func EnvDBUser() string {
	return GetEnv("DB_USER", "")
}

// EnvDBPassword returns the database password from environment variables.
func EnvDBPassword() string {
	return GetEnv("DB_SECRET", "")
}

// EnvDBName returns the database name from environment variables.
func EnvDBName() string {
	return GetEnv("DB_NAME", "")
}

// EnvironmentConfig returns the environment configuration.
func EnvironmentConfig() string {
	return GetEnv("ENV", entities.Environment.Development)
}

// EnvServiceName returns the service name from environment variables.
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

// EnvAmqpConnection returns the AMQP connection string from environment variables.
func EnvAmqpConnection() string {
	user := envUserAmqp()
	password := envPasswordAmqp()
	host := envHostAmqp()

	return fmt.Sprintf("amqp://%s:%s@%s/", user, password, host)
}

// LoadEnvVars loads all environment variables required by the application.
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
