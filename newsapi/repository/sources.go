package repository

import (
	"github.com/rs/zerolog/log"
	"net/url"
)

// AddNewSource adds a new RSS feed source to the repository under the name provided
func (r *ArticlesRepo) AddNewSource(name string, url string) bool {
	if !isUrl(url) {
		return false
	}
	r.mu.Lock()
	r.sourceUrls[name] = url
	// TODO: check if it already exists instead of just overwriting
	r.mu.Unlock()

	log.Info().Msgf("Successfully added a new source to the API with name %s and url %s", name, url)
	return true
}

// GetSources returns RSS feed sources currently registered with, and monitored by, the API
func (r *ArticlesRepo) GetSources() map[string]string {
	return r.sourceUrls
}

// DelSource deletes an RSS feed source currently monitored by the API
func (r *ArticlesRepo) DelSource(name string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.sourceUrls[name]; !exists {
		log.Info().Msgf("Failed to delete a source with name %s", name)
		return false
	}
	delete(r.sourceUrls, name)
	log.Info().Msgf("Successfully deleted a source with name %s", name)
	// TODO: consider deleting also articles currently stored which came from this source
	return true
}

// isUrl checks if the string provided is a valid URL address
func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
