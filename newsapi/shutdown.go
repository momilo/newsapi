package newsapi

import (
	"context"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// shutdownCleanup performs the cleanup of the API server upon shutdown
func (api *newsApi) shutdownCleanup() {
	if !api.ready {
		return
	}
	api.ready = false

	log.Info().Msg("Server shutting down.")

	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	if err := api.server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to shut down the http server gracefully within the specified timeout")
	}
}

// HealthMonitor listens for TERM signals and, if one is received, (indirectly) activates a graceful shutdown
func (api *newsApi) healthMonitor(closeServerCtx *context.Context, closeServer *context.CancelFunc) {
	log.Info().Msg("Health monitor launched")
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	select {
	case <-interruptCh:
		log.Info().Msg("SIGTERM received - server shutdown initialised")
		(*closeServer)()
		return
	case <-(*closeServerCtx).Done():
		return
	}
}
