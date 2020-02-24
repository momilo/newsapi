package routes

import (
	"github.com/gin-gonic/gin"
	"newsapi/newsapi/controllers/articles"
)

// SetArticlesRoutes defines routes fot the Articles handles
func SetArticlesRoutes(g *gin.RouterGroup) {
	g.GET("/articles", articles.GetArticlesList)
	g.GET("/article/:id", articles.GetArticle)
	g.GET("/categories", articles.GetCategories)
	g.GET("/sources", articles.GetSources)
	g.POST("/source", articles.ModifySource)
}
