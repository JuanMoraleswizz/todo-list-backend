package main

import (
	"net/http"
	"testing"
)

func TestCorsMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	middleware := corsMiddleware(handler)

	req, _ := http.NewRequest("OPTIONS", "/", nil)
	rr := &mockResponseWriter{headers: make(http.Header)}

	middleware.ServeHTTP(rr, req)

	if rr.headers.Get("Access-Control-Allow-Origin") != "*" {
		t.Error("CORS headers not set correctly")
	}
	if rr.statusCode != http.StatusNoContent {
		t.Errorf("Expected status 204 for OPTIONS, got %d", rr.statusCode)
	}
}

type mockResponseWriter struct {
	headers    http.Header
	statusCode int
}

func (m *mockResponseWriter) Header() http.Header {
	return m.headers
}

func (m *mockResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.statusCode = statusCode
}