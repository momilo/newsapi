package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SayHello is a simple static page which - ideally - should include details about the API
func SayHello(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to News API. Imagine you can find "+
		"a beautifully-formatted documentation here.")
	// TODO: serve an appropriately-formatted HTML page with API details
	return
}
