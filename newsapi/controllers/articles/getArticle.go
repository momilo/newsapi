package articles

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetArticle handles a request which includes an article UID, responding with the article's contents
func GetArticle(c *gin.Context) {
	if repo == nil {
		sendInternalError(c)
		return
	}
	id := c.Param("id")
	body, ok := repo.GetArticleBody(id)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist})
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", body)
}

// TODO: implement serving clean mobile-friendly article, vs. the whole webpage. Ideally, serve over a more efficient protocol (e.g. gRPC)
