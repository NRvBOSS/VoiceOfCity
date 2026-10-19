package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/health", nil)

	Health(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status code: got %d, want %d",
			recorder.Code, http.StatusOK)
	}

	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type: got %q, want %q",
			got, "application/json")
	}

	var response HealthResponse
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("JSON decode failed: %v", err)
	}

	if response.Status != "ok" {
		t.Errorf("status: got %q, want %q", response.Status, "ok")
	}

	if response.Service != "citysesi-api" {
		t.Errorf("service: got %q, want %q",
			response.Service, "citysesi-api")
	}
}