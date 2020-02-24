package articles

import "github.com/gin-gonic/gin"

// GetSources handles a request responding with the list of currently-defined rss sources which the API is using (and their
// codenames)
func GetSources(c *gin.Context) {
	if repo == nil {
		sendInternalError(c)
		return
	}
	response := repo.GetSources()
	c.JSON(200, response)
}
