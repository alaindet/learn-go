package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

// Returns an application stub
func newTestApplication(t *testing.T) *application {
	return &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}
}

// Wrap the test server in a struct in order to add methods to it
type testServer struct {
	*httptest.Server
}

// Creates a fake test server
func newTestServer(t *testing.T, handler http.Handler) *testServer {
	server := httptest.NewTLSServer(handler)
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	server.Client().Jar = jar

	// Disable following any 3xx redirect
	server.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{server}
}

// Helper to perform a GET request on the test server
// Returns the status code, the headers and the stringified body
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {

	res, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)
	return res.StatusCode, res.Header, string(body)
}
