package newsapi

import (
	"context"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Start launches the previously-initialised newsApi
func (api *newsApi) Start() {
	if !api.ready {
		log.Error().Msg("Failed to start API - not ready")
		return
	}
	log.Info().Msg("Starting newsApi")

	// launch a healthMonitor to monitor for SIGTERM etc.
	closeServerCtx, closeServerFunc := context.WithCancel(context.Background())
	go api.healthMonitor(&closeServerCtx, &closeServerFunc)

	// shutdownCleanup will be triggered (cleanly) by healthMonitor on SIGTERM, otherwise -
	// if this function unexpectedly returns (e.g. because of an error in ListenAndServe, triggering closeServerFunc
	defer api.shutdownCleanup()

	// Cache all articles from the initial sources provided
	api.repo.CacheAll()

	// launch a goroutine adding new articles to the cache, as they become available through known sources
	go api.repo.AutoCache(closeServerCtx)

	go func() {
		defer closeServerFunc()
		log.Info().Msgf("Listening for incoming requests on port %s", api.server.Addr)
		err := api.server.ListenAndServe()
		if err != nil {
			switch err {
			case http.ErrServerClosed:
				log.Info().Err(err).Msg("HTTP server closed")
			default:
				log.Error().Err(err).Msg("HTTP server stopped - fatal error")
			}
			return
		}
	}()

	// Block until the healthMonitor determines that the server should be closed
	<-closeServerCtx.Done()
}
