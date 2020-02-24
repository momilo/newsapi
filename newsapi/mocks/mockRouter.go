package mocks

import (
	"github.com/gin-gonic/gin"
)

func GetMockSetup() (*MockRepo, *gin.Engine) {
	repo := new(MockRepo)
	router := gin.New()
	return repo, router
}
