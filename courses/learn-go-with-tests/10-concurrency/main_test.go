package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {

	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	expected := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	result := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected: %v Result: %v", expected, result)
	}
}

func slowStubWebsiteChecker(url string) bool {
	_ = url
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {

	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "some-fake-url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
