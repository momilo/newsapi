package routes

import (
	"github.com/gin-gonic/gin"
	"newsapi/newsapi/controllers/core"
)

// SetCoreRoutes defines the appropriate routes fot the core handles
func SetCoreRoutes(router *gin.Engine) {
	router.GET("/", core.GcpHealthCheck)
	router.GET("/hello", core.SayHello)
}
