package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func assertStrings(t testing.TB, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("Result: %q Expected: %q", result, expected)
	}
}

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	// Fake slow response
	go func(ctx context.Context) {
		var result string
		for _, char := range s.response {
			select {
			// Stop this goroutine if context is done
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(char)
			}
		}
		data <- result
	}(ctx)

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestServer(t *testing.T) {

	t.Run(
		"returns data from store",
		func(t *testing.T) {
			storeData := "hello, world"
			store := &SpyStore{
				response: storeData,
				t:        t,
			}
			server := Server(store)
			request := httptest.NewRequest(http.MethodGet, "/", nil)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)
			assertStrings(t, response.Body.String(), storeData)
		},
	)

	// t.Run(
	// 	"tells store to cancel work if request is cancelled",
	// 	func(t *testing.T) {
	// 		storeData := "hello, world"
	// 		store := &SpyStore{response: storeData}
	// 		server := Server(store)

	// 		request := httptest.NewRequest(http.MethodGet, "/", nil)

	// 		// Create a derived context with cancelling signal
	// 		// by extending the request context
	// 		cancellingCtx, cancel := context.WithCancel(request.Context())

	// 		// Run the function signaling cancellation after 10 milliseconds
	// 		time.AfterFunc(10*time.Millisecond, cancel)

	// 		// Create a new request context by extending the canceling context
	// 		request = request.WithContext(cancellingCtx)

	// 		response := httptest.NewRecorder()

	// 		server.ServeHTTP(response, request)

	// 		store.assertWasCancelled()
	// 	},
	// )
}
