package main

import (
	"net/http"
	"testing"

	"snippetbox.dev/internal/assert"
)

func TestPing(t *testing.T) {

	// Arrange
	app := newTestApplication(t)
	server := newTestServer(t, app.routes())
	defer server.Close()

	// Act
	statusCode, _, body := server.get(t, "/ping")

	// Assert
	assert.Equal(t, statusCode, http.StatusOK)
	assert.Equal(t, body, "OK")
}
