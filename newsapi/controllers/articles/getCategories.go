package articles

import "github.com/gin-gonic/gin"

// GetCategories handles a request responding with the list of article categories currently stored in the api
func GetCategories(c *gin.Context) {
	if repo == nil {
		sendInternalError(c)
		return
	}
	response := repo.GetCategories()
	c.JSON(200, response)
}
