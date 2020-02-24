package repository

import (
	"github.com/rs/zerolog/log"
	"newsapi/newsapi/models"
	"sort"
)

// GetCategories provides a slice of categories of articles stored in the repository
func (r *ArticlesRepo) GetCategories() []string {
	categoryNames := make([]string, len(r.categories))
	i := 0
	r.mu.Lock()
	for key, _ := range r.categories {
		categoryNames[i] = key
		i++
	}
	r.mu.Unlock()
	sort.Strings(categoryNames)
	return categoryNames
}

// GetArticlesList provides a list of all articles stored in the repository
func (r *ArticlesRepo) GetArticlesList() []*models.Article {
	// TODO: refactor in line with GetCategories to avoid unnecessary allocations by append
	var articles []*models.Article
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, article := range r.articles {
		articles = append(articles, article)
	}

	// sort articles by publishing date
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Date.After(articles[j].Date)
	})
	return articles
}

// GetArticlesListByCategory provides a list of all articles stored in the repository which fall into the specified category
func (r *ArticlesRepo) GetArticlesListByCategory(category string) []*models.Article {
	r.mu.RLock()
	articlesList, found := r.categories[category]
	r.mu.RUnlock()
	if !found {
		return nil
	}
	var articles []*models.Article
	for a := articlesList.Front(); a != nil; a = a.Next() {
		article, ok := a.Value.(*models.Article)
		if !ok {
			log.Error().Msg("Failed to assert a categories list entry as a pointer to article - possibly corrupted list")
			continue
		}

		articles = append(articles, article)
	}
	return articles
}

// GetArticleBody fetches the body of an article stored in the repository's cache
func (r *ArticlesRepo) GetArticleBody(uid string) ([]byte, bool) {
	body, found := r.webCache.Get(uid)
	if !found {
		return nil, false
	}
	return body, true
}
