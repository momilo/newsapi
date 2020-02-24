package mocks

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func PerformRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func ValidResponseJson(t *testing.T, w *httptest.ResponseRecorder, wantCode int, wantResponse []byte) bool {
	if w.Code != wantCode {
		t.Error("wantcode: ", wantCode, " receivedCode: ", w.Code)
		return false
	}

	equal, err := jsonBytesEqual(wantResponse, []byte(w.Body.String()))
	if err != nil {
		t.Error("failed to compare jsons:", err)
		return false
	}
	if !equal {
		t.Error("wantResponse: ", string(wantResponse), " gotResponse: ", w.Body.String())
		return false
	}

	return true
}

func jsonBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}
