package system

import (
	"github.com/RodolfoBonis/microdetect-api/features/system/di"
	"github.com/gin-gonic/gin"
)

func InjectRoutes(route *gin.RouterGroup) {
	var systemUc = di.SystemInjection()

	systemRoute := route.Group("/system")
	systemRoute.GET("/storage", systemUc.GetStorage)
	systemRoute.GET("/memory", systemUc.GetMemory)
	systemRoute.GET("/status", systemUc.GetSystemStatus)
	systemRoute.GET("/cpu", systemUc.GetCPU)

}
