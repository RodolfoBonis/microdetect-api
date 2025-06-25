package config

import (
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
)

// SentryConfig returns the Sentry configuration for the application.
func SentryConfig() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              EnvSentryDSN(),
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		os.Exit(1)
	}
}
