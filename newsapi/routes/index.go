package routes

import (
	"github.com/gin-gonic/gin"
	"newsapi/newsapi/controllers/articles"
)

// RegisterRoutes registers appropriate routes with the router and injects the provided repo into the handling functions
func RegisterRoutes(router *gin.Engine, repo articles.Repository) {
	SetCoreRoutes(router)

	v1 := router.Group("/v1")
	SetArticlesRoutes(v1)
	// Provide repository handle to the articles controllers
	articles.SetRepo(repo)
}
