package system

import (
	"github.com/RodolfoBonis/microdetect-api/features/system/di"
	"github.com/gin-gonic/gin"
)

func InjectRoutes(route *gin.RouterGroup) {
	var systemUc = di.SystemInjection()

	systemRoute := route.Group("/system")
	systemRoute.GET("/", systemUc.GetSystemStatus)
}
