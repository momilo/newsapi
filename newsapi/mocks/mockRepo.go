package mocks

// TODO: refactor the mocks package to be more flexible

import (
	"newsapi/newsapi/models"
	"time"
)

type MockRepo struct{}

func (r *MockRepo) GetArticleBody(uid string) ([]byte, bool) {
	body := `<!DOCTYPE html> <html lang="en"> <head> <meta charset="UTF-8"> <title>I am an article</title> </head> <body> This is an exciting article </body> </html>`
	return []byte(body), true
}
func (r *MockRepo) GetCategories() []string {
	return []string{"oneCategory", "anotherCategory"}
}
func (r *MockRepo) GetArticlesListByCategory(category string) []*models.Article {
	article := getMockArticle()
	article.Category = category
	article2 := article
	article2.Title = "Second Article"
	return []*models.Article{&article, &article2}
}
func (r *MockRepo) GetArticlesList() []*models.Article {
	article := getMockArticle()
	article2 := article
	article2.Title = "Second Article"
	return []*models.Article{&article, &article2}
}
func (r *MockRepo) GetSources() map[string]string {
	m := map[string]string{
		"sourceOne": "http://excitingsourceone.com",
		"sourceTwo": "http://excitingsourcetwo.com",
	}
	return m
}

func (r *MockRepo) AddNewSource(name string, url string) bool {
	return true
}
func (r *MockRepo) DelSource(name string) bool {
	return true
}

func getMockArticle() models.Article {
	return models.Article{
		Title:       "First Article",
		Category:    "someCategory",
		Link:        "http://excitingUrl.com",
		Description: "I am description",
		PubDate:     time.Now().Format(time.RFC1123Z),
		Date:        time.Time{},
		UID:         "thisIsSuperUnique",
		Source:      "greatSource",
	}
}
