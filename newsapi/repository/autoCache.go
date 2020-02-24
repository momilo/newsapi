package repository

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

// AutoCache regularly fetches and stores articles from the rss feed sources defined with the repo
func (r *ArticlesRepo) AutoCache(ctx context.Context) {
	log.Debug().Msg("Starting AutoCache")
	defer log.Debug().Msg("AutoCache stopped")

	ticker := time.NewTicker(2 * time.Minute)
	// TODO: caching frequency should be pulled from an environment variable

	for {
		select {
		case <-ticker.C:
			r.CacheAll()
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}
