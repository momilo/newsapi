package articles

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"newsapi/newsapi/mocks"
	"testing"
)

func TestModifySource(t *testing.T) {
	mockRepo, router := mocks.GetMockSetup()
	repo = mockRepo
	router.GET("/v1/categories", ModifySource)

	wantResponseRaw, _ := json.Marshal(gin.H{"status": "success", "message": "Successfully added a new source. " +
		"Articles from it should become available shortly."})
	wantCode := http.StatusOK

	truebool := true
	requestBody := modifySourceMsg{
		Name: "newSource",
		Url:  "http://somesource.com",
		Add:  &truebool,
	}
	requestBodyRaw, _ := json.Marshal(requestBody)

	w := mocks.PerformRequest(router, "GET", "/v1/categories", bytes.NewReader(requestBodyRaw))
	if !mocks.ValidResponseJson(t, w, wantCode, wantResponseRaw) {
		t.Fail()
	}
	// TODO: add table tests for various queries
}
