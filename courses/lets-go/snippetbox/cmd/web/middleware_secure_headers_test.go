package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"snippetbox.dev/internal/assert"
)

func TestSecureHeaders(t *testing.T) {

	// Arrange
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Act
	secureHeaders(next).ServeHTTP(rec, req)
	res := rec.Result()

	// Assert
	headers := []struct {
		name     string
		expected string
	}{
		{
			"Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com",
		},
		{
			"Referrer-Policy",
			"origin-when-cross-origin",
		},
		{
			"X-Content-Type-Options",
			"nosniff",
		},
		{
			"X-Frame-Options",
			"deny",
		},
	}

	for _, header := range headers {
		assert.Equal(t, res.Header.Get(header.name), header.expected)
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
