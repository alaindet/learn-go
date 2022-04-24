package main

import (
	"fmt"
	"net/http"
	"time"
)

func ConfigurableUrlRacer(a, b string, t time.Duration) (string, error) {
	select {
	// functions (ping, time.After) are evaluated asap
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(t):
		return "", fmt.Errorf("timed out (%d s) for %q and %q", t, a, b)
	}
}

func UrlRacer(a, b string) (string, error) {
	return ConfigurableUrlRacer(a, b, 3*time.Second)
}

// Returns a channel whose type is not important
// struct{} was chosen since it does not allocate anything
// so it doesn't tax the garbage collector later
func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
