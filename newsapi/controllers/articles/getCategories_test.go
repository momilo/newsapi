package articles

import (
	"encoding/json"
	"net/http"
	"newsapi/newsapi/mocks"
	"testing"
)

func TestGetCategories(t *testing.T) {
	mockRepo, router := mocks.GetMockSetup()
	repo = mockRepo
	router.GET("/v1/categories", GetCategories)

	wantResponseRaw, _ := json.Marshal(repo.GetCategories())
	wantCode := http.StatusOK

	w := mocks.PerformRequest(router, "GET", "/v1/categories", nil)
	if !mocks.ValidResponseJson(t, w, wantCode, wantResponseRaw) {
		t.Fail()
	}
	// TODO: add table tests for various queries
}
