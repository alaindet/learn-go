package main

import (
	"net/http"
	"strconv"
	"testing"

	"snippetbox.dev/internal/assert"
	"snippetbox.dev/internal/models/mocks"
)

func TestSnippetView(t *testing.T) {

	// Arrange
	app := newTestApplication(t)
	server := newTestServer(t, app.routes())
	defer server.Close()

	// Act
	snippet := mocks.MockSnippet
	testCases := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid ID",
			url:            "/snippets/view/" + strconv.Itoa(snippet.ID),
			expectedStatus: http.StatusOK,
			expectedBody:   snippet.Content,
		},
		{
			name:           "Non-existing ID",
			url:            "/snippets/view/2",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Invalid ID",
			url:            "/snippets/view/-123",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Decimal ID",
			url:            "/snippets/view/12.34",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			status, _, body := server.get(t, testCase.url)
			assert.Equal(t, status, testCase.expectedStatus)
			if testCase.expectedBody != "" {
				assert.StringContains(t, body, testCase.expectedBody)
			}
		})
	}
}
