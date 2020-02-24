package repository

import (
	"github.com/rs/zerolog/log"
	"newsapi/newsapi/models"
)

// CacheAll iterates through the repository's source URLs, fetching all articles and storing them in the repository
func (r *ArticlesRepo) CacheAll() {
	if len(r.sourceUrls) == 0 {
		log.Info().Msg("Failed to CacheAll - no article sources available in the repository")
		return
	}
	log.Debug().Msg("Caching all articles")
	r.mu.RLock()
	defer r.mu.RUnlock()
	for sourceName, sourceUrl := range r.sourceUrls {
		articles := models.FetchArticles(sourceUrl)
		if articles == nil {
			return
		}
		log.Debug().Msgf("Fetched %d articles from %s", len(articles), sourceUrl)
		go r.StoreArticles(articles, sourceName)
	}
}
