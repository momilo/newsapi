package articles

import (
	"encoding/json"
	"net/http"
	"newsapi/newsapi/mocks"
	"testing"
)

func TestGetSources(t *testing.T) {
	mockRepo, router := mocks.GetMockSetup()
	repo = mockRepo
	router.GET("/v1/sources", GetSources)

	wantResponseRaw, _ := json.Marshal(repo.GetSources())
	wantCode := http.StatusOK

	w := mocks.PerformRequest(router, "GET", "/v1/sources", nil)
	if !mocks.ValidResponseJson(t, w, wantCode, wantResponseRaw) {
		t.Fail()
	}
	// TODO: add table tests for various queries
}
