package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"final_project/data"
)

var pageTests = []struct {
	name               string
	url                string
	expectedStatusCode int
	handler            http.HandlerFunc
	sessionData        map[string]any
	expectedHTML       string
}{
	{
		name:               "home",
		url:                "/",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.HomePage,
	},
	{
		name:               "login",
		url:                "/login",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.LoginPage,
		expectedHTML:       `<h1 class="mt-5">Login</h1>`,
	},
	// Visiting log out with a session redirects you to GET /login
	{
		name:               "logout",
		url:                "/logout",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.LoginPage,
		sessionData: map[string]any{
			"userID": 1,
			"user":   data.User{},
		},
	},
}

func TestGetHandlers(t *testing.T) {
	templatesPath = "./templates"

	for _, pageTest := range pageTests {
		t.Run(pageTest.name, func(t *testing.T) {
			t.Helper()
			res := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", pageTest.url, nil)

			ctx := getContext(req)
			req = req.WithContext(ctx)

			if len(pageTest.sessionData) > 0 {
				for key, value := range pageTest.sessionData {
					testApp.Session.Put(ctx, key, value)
				}
			}

			pageTest.handler.ServeHTTP(res, req)

			// Check HTTP code
			if res.Code != pageTest.expectedStatusCode {
				t.Errorf(
					"expected HTTP code %d, but got %d",
					pageTest.expectedStatusCode,
					res.Code,
				)
			}

			// Check HTML
			if len(pageTest.expectedHTML) > 0 {
				html := res.Body.String()
				if !strings.Contains(html, pageTest.expectedHTML) {
					t.Errorf("Expected to find %q but did not", pageTest.expectedHTML)
				}
			}
		})
	}
}

func TestPostLoginPage(t *testing.T) {
	templatesPath = "./templates"

	postedData := url.Values{
		"email":    {"admin@example.com"},
		"password": {"abc123abc123abc123abc123"},
	}

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(postedData.Encode()))
	ctx := getContext(req)
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(testApp.PostLoginPage)
	handler.ServeHTTP(res, req)

	if res.Code != http.StatusSeeOther {
		t.Errorf("expected HTTP code 303, but got %d", res.Code)
	}

	if !testApp.Session.Exists(ctx, "userID") {
		t.Errorf("Did not find userID in the session")
	}
}
