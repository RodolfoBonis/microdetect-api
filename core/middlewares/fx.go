package middlewares

import (
	"go.uber.org/fx"
)

var Module = fx.Module("middlewares",
	fx.Provide(
		NewMonitoringMiddleware,
		NewProtectMiddleware,
	),
)
