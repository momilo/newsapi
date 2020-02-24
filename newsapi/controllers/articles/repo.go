package articles

import (
	"newsapi/newsapi/models"
)

// Repository is required by handlers to fulfil their function. It needs to be set before any queries are routed
var repo Repository

type Repository interface {
	GetArticleBody(uid string) ([]byte, bool)
	GetCategories() []string
	GetArticlesListByCategory(category string) []*models.Article
	GetArticlesList() []*models.Article
	GetSources() map[string]string
	AddNewSource(name string, url string) bool
	DelSource(name string) bool
}

// Sets the above
func SetRepo(r Repository) {
	repo = r
}
