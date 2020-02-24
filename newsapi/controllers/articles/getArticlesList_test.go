package articles

import (
	"encoding/json"
	"net/http"
	"newsapi/newsapi/mocks"
	"testing"
)

func TestGetArticlesList(t *testing.T) {
	mockRepo, router := mocks.GetMockSetup()
	repo = mockRepo
	router.GET("/v1/articles", GetArticlesList)

	wantResponseRaw, _ := json.Marshal(repo.GetArticlesList())
	wantCode := http.StatusOK

	w := mocks.PerformRequest(router, "GET", "/v1/articles", nil)
	if !mocks.ValidResponseJson(t, w, wantCode, wantResponseRaw) {
		t.Fail()
	}
	// TODO: add table tests for various queries
}
