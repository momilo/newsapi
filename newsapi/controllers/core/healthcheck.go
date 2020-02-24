package core

import (
	"github.com/gin-gonic/gin"
)

// GcpHealthCheck for GCP Load Balancer
func GcpHealthCheck(c *gin.Context) {
	if c.Request.URL.Path == "/" {
		c.AbortWithStatus(200)
	}
	return
}
