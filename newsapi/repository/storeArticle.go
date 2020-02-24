package repository

import (
	"container/list"
	"github.com/rs/zerolog/log"
	"newsapi/newsapi/models"
)

// storeArticle stores the article details in the repository and fetches and stores its body in the cache
// TODO consider adding appropriate handling if there have been changes/updates in the article bodies
func (r *ArticlesRepo) storeArticle(a models.Article) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, found := r.articles[a.Link]; found {
		log.Debug().Msgf("Article with link [%s] is already stored in the repository, ignoring", a.Link)
		return
	}

	a.FillMissingDetails()

	r.articles[a.Link] = &a
	go r.addToCategory(&a)
	go r.cacheArticleBody(&a)

	log.Debug().Msgf("Successfully added to the repo an article with UID [%s], category [%s], and url [%s]", a.UID, a.Category, a.Link)
}

// StoreArticles stores the articles provided in the repository
func (r *ArticlesRepo) StoreArticles(articles []models.Article, sourceName string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, _ := range articles {
		articles[i].Source = sourceName
		go r.storeArticle(articles[i])
	}
}

// addToCategory adds the article in the list of articles for the appropriate category
func (r *ArticlesRepo) addToCategory(a *models.Article) {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, found := r.categories[a.Category]
	if !found {
		r.categories[a.Category] = list.New()
	}
	r.categories[a.Category].PushBack(a)
}

func (r *ArticlesRepo) cacheArticleBody(a *models.Article) {
	body, err := a.FetchArticleBody()
	if err != nil {
		log.Error().Err(err).Str("url", a.Link).Msg("Failed to fetch article's body")
		return
	}
	r.webCache.Set(a.UID, body)
	log.Debug().Msgf("Successfully stored in cache the html body of the article num %s", a.UID)
}
