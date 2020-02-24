package repository

import (
	"container/list"
	"newsapi/newsapi/models"
	"sync"
)

type ArticlesRepo struct {
	mu         *sync.RWMutex              // mutex for all repo's maps; consider refactoring for better efficiency
	sourceUrls map[string]string          // name -> url
	articles   map[string]*models.Article // link -> article
	categories map[string]*list.List      // category -> list of pointers to articles
	webCache   dbInterface
}

type dbInterface interface {
	Get(articleId string) ([]byte, bool)
	Set(articleId string, articleBody []byte)
}

// New initialises a new Articles Repository, using the cache database provided
func New(db dbInterface) *ArticlesRepo {
	repo := &ArticlesRepo{
		mu:         new(sync.RWMutex),
		sourceUrls: make(map[string]string),
		articles:   make(map[string]*models.Article),
		categories: make(map[string]*list.List),
		webCache:   db,
	}
	return repo
}
