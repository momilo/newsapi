package ristrettodb

import (
	"github.com/dgraph-io/ristretto"
)

// risCache is a Ristretto-based implementation of dbInterface, required by the API repository to cache articles
type risCache struct {
	*ristretto.Cache
}

// New initialises the risCache with default parameters.
func New() (*risCache, error) {
	cache := new(risCache)
	c, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 5000, // number of keys to track frequency of
		MaxCost:     2000, // maximum cost of Cache
		BufferItems: 64,   // number of keys per Get buffer (default value is 64 but further optimisation per use-case might be useful)
	})
	// TODO: read all these settings from environment variables
	if err != nil {
		return nil, err
	}
	cache.Cache = c
	return cache, nil
}

// Get retrieves the body of a previously-cached article, returning false it if could not be found or properly retrieved
func (c *risCache) Get(articleId string) ([]byte, bool) {
	bodyRaw, found := c.Cache.Get(articleId)
	if !found {
		return nil, false
	}
	body, ok := bodyRaw.([]byte)
	if !ok {
		return nil, false
	}
	return body, true
}

// Set stores the html body of an article under the article's UID. All articles are presumed to be of equal weight, regardless
// of their size
func (c *risCache) Set(articleId string, articleBody []byte) {
	// TODO: Implement caching of underlying images and other resources used by the article page
	c.Cache.Set(articleId, articleBody, 1)
}
