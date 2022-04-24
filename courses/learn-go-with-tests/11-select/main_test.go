package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// This server just return an empty 200 OK HTTP response
// and has an artificial configurable delay
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}),
	)
}

func TestUrlRacer(t *testing.T) {

	t.Run(
		"returns the faster server to respond",
		func(t *testing.T) {
			slowServer := makeDelayedServer(20 * time.Millisecond)
			defer slowServer.Close()

			fastServer := makeDelayedServer(0 * time.Millisecond)
			defer fastServer.Close()

			expected := fastServer.URL
			result, _ := UrlRacer(slowServer.URL, fastServer.URL)

			if result != expected {
				t.Errorf("Result: %q, Expected: %q", result, expected)
			}
		},
	)

	t.Run(
		"returns error if server does not response in 100 milliseconds",
		func(t *testing.T) {
			slowServer := makeDelayedServer(150 * time.Millisecond)
			defer slowServer.Close()

			fastServer := makeDelayedServer(120 * time.Millisecond)
			defer fastServer.Close()

			timeout := 100 * time.Millisecond
			_, err := ConfigurableUrlRacer(slowServer.URL, fastServer.URL, timeout)

			if err == nil {
				t.Error("expected error but didnt' get one")
			}
		},
	)
}
