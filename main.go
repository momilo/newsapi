package main

import (
	"github.com/rs/zerolog/log"
	"newsapi/newsapi"
	"newsapi/newsapi/logging"
	"newsapi/newsapi/repository"
	"newsapi/newsapi/repository/ristrettodb"
)

func main() {
	logging.ConfigureGlobalLogger()

	// Initialise the underlying articles cache
	webCache, err := ristrettodb.New()
	if err != nil {
		log.Error().Err(err).Msg("Failed to initialise the underlying webCache")
		return
	}

	// Initialise articles repository with the desired underlying cache
	repo := repository.New(webCache)

	// Initialise the api
	api := newsapi.New(repo)
	if api == nil {
		log.Error().Err(err).Msg("Failed to initialise the api")
		return
	}
	api.Start()
}
