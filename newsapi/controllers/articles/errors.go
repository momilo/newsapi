package articles

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	errNotExist        = "Article does not exist"
	errInvalidCategory = "No articles are available in the specified category"
	errAddingSource    = "Error in adding a new source"
	errDeletingSource  = "Error in deleting source"
	errInvalidBody     = "Invalid request body"
	errInternalError   = "Internal server error"
)

func sendInternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": errInternalError})
}
