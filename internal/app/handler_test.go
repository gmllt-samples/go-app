package app

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleRequest_JSON(t *testing.T) {
	app := NewApp()

	req := httptest.NewRequest("GET", "/?status=201&response_size=1K&wait=10ms", nil)
	w := httptest.NewRecorder()

	app.HandleRequest(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 201 {
		t.Errorf("Expected status 201, got %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Expected JSON response, got %s", contentType)
	}

	var parsed struct {
		Status int    `json:"status"`
		URI    string `json:"uri"`
		Body   string `json:"body"`
	}
	if err := json.NewDecoder(res.Body).Decode(&parsed); err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
	}

	if parsed.Status != 201 {
		t.Errorf("JSON field 'status' is incorrect: got %d", parsed.Status)
	}

	if len(parsed.Body) < 1000 {
		t.Errorf("Expected body of at least 1K, got %d bytes", len(parsed.Body))
	}
}
