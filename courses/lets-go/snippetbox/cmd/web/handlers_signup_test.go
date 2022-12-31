package main

import (
	"net/http"
	"net/url"
	"testing"

	"snippetbox.dev/internal/assert"
	"snippetbox.dev/internal/models/mocks"
)

// Run this with
// go test -count=1 -v -run="^TestSignUp$" ./cmd/web
func TestSignUp(t *testing.T) {

	// Arrange
	app := newTestApplication(t)
	server := newTestServer(t, app.routes())
	defer server.Close()

	_, _, body := server.get(t, "/users/signup")
	validCSRFToken := extractCSRFToken(t, body)

	const (
		validName     = "Bob"
		validPassword = "validPa$$word"
		validEmail    = "bob@example.com"
		formTag       = `<form action="/users/signup" method="POST" novalidate>`
	)

	testCases := []struct {
		name              string
		inputUserName     string
		inputUserEmail    string
		inputUserPassword string
		csrfToken         string
		expectedCode      int
		expectedFormTag   string
	}{
		{
			name:              "Valid submission",
			inputUserName:     validName,
			inputUserEmail:    validEmail,
			inputUserPassword: validPassword,
			csrfToken:         validCSRFToken,
			expectedCode:      http.StatusSeeOther,
		},
		{
			name:              "Invalid CSRF Token",
			inputUserName:     validName,
			inputUserEmail:    validEmail,
			inputUserPassword: validPassword,
			csrfToken:         "wrongToken",
			expectedCode:      http.StatusBadRequest,
		},
		{
			name:              "Empty name",
			inputUserName:     "",
			inputUserEmail:    validEmail,
			inputUserPassword: validPassword,
			csrfToken:         validCSRFToken,
			expectedCode:      http.StatusUnprocessableEntity,
			expectedFormTag:   formTag,
		},
		{
			name:              "Empty email",
			inputUserName:     validName,
			inputUserEmail:    "",
			inputUserPassword: validPassword,
			csrfToken:         validCSRFToken,
			expectedCode:      http.StatusUnprocessableEntity,
			expectedFormTag:   formTag,
		},
		{
			name:              "Empty password",
			inputUserName:     validName,
			inputUserEmail:    validEmail,
			inputUserPassword: "",
			csrfToken:         validCSRFToken,
			expectedCode:      http.StatusUnprocessableEntity,
			expectedFormTag:   formTag,
		},
		{
			name:              "Invalid email",
			inputUserName:     validName,
			inputUserEmail:    "bob@example.",
			inputUserPassword: validPassword,
			csrfToken:         validCSRFToken,
			expectedCode:      http.StatusUnprocessableEntity,
			expectedFormTag:   formTag,
		},
		{
			name:              "Short password",
			inputUserName:     validName,
			inputUserEmail:    validEmail,
			inputUserPassword: "pa$$",
			csrfToken:         validCSRFToken,
			expectedCode:      http.StatusUnprocessableEntity,
			expectedFormTag:   formTag,
		},
		{
			name:              "Duplicate email",
			inputUserName:     validName,
			inputUserEmail:    mocks.MockDuplicateEmail,
			inputUserPassword: validPassword,
			csrfToken:         validCSRFToken,
			expectedCode:      http.StatusUnprocessableEntity,
			expectedFormTag:   formTag,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			form := url.Values{}
			form.Add("name", testCase.inputUserName)
			form.Add("email", testCase.inputUserEmail)
			form.Add("password", testCase.inputUserPassword)
			form.Add("csrf_token", testCase.csrfToken)

			code, _, body := server.postForm(t, "/users/signup", form)
			assert.Equal(t, code, testCase.expectedCode)

			if testCase.expectedFormTag != "" {
				assert.StringContains(t, body, testCase.expectedFormTag)
			}
		})
	}

}
