package newsapi

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"newsapi/newsapi/repository"
	"newsapi/newsapi/routes"
	"os"
	"strconv"
)

type newsApi struct {
	server *http.Server
	router *gin.Engine
	repo   *repository.ArticlesRepo
	ready  bool
}

// New returns a newly-initialised newsAPI
func New(repo *repository.ArticlesRepo) *newsApi {
	log.Info().Msg("Initialising newsApi")
	if repo == nil {
		log.Error().Msg("failed to initialise api - nil repository reference provided")
		return nil
	}

	// configure the server
	api := new(newsApi)
	gin.SetMode(gin.ReleaseMode)
	api.router = gin.New()

	// silence debug messages of the router

	api.server = &http.Server{Handler: api.router}
	api.SetPort(getPort())
	api.repo = repo

	// TODO: fetch any new sources directly from an independent source (e.g. GCP Datastore)
	// Initialise with default sources and cache them
	repo.AddNewSource("bbcuk", "http://feeds.bbci.co.uk/news/uk/rss.xml")
	repo.AddNewSource("reutersdomesticuk", "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml")
	routes.RegisterRoutes(api.router, api.repo)
	api.ready = true
	return api
}

// SetPort sets the port for the api (1-65535)
func (api *newsApi) SetPort(p int) {
	if p < 1 || p > 65535 {
		p = 80
	}
	api.server.Addr = ":" + strconv.Itoa(p)
}

// getPort attempts to retrieve the desired port number from environment variable, returning the default (80) if not found
func getPort() int {
	portString := os.Getenv("API_PORT")
	port, err := strconv.Atoi(portString)
	if err != nil {
		return 80
	}
	return port
}
