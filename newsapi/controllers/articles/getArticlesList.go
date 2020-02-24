package articles

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newsapi/newsapi/models"
)

// GetArticlesList handles a request responding with a list of all articles fetched by the api. If "cat" query param
// is specified, it will provide only the articles which fall within the specified category
func GetArticlesList(c *gin.Context) {
	if repo == nil {
		sendInternalError(c)
		return
	}

	category := c.DefaultQuery("cat", "")

	var response []*models.Article

	if category != "" {
		response = repo.GetArticlesListByCategory(category)
		if response == nil || len(response) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidCategory})
			return
		}
	} else {
		response = repo.GetArticlesList()
	}
	c.JSON(200, response)
}
