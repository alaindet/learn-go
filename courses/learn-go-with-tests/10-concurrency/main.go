package main

type UrlChecker func(string) bool

type urlCheckResult struct {
	url    string
	online bool
}

func CheckWebsites(checker UrlChecker, urls []string) map[string]bool {

	checked := make(map[string]bool, len(urls))
	checkedCh := make(chan urlCheckResult)
	checkedChSet := (chan<- urlCheckResult)(checkedCh)
	checkedChGet := (<-chan urlCheckResult)(checkedCh)

	// Write into channel
	for _, urlToCheck := range urls {
		go func(url string, ch chan<- urlCheckResult) {
			online := checker(url)
			r := urlCheckResult{url, online}
			checkedCh <- r // Send to channel
		}(urlToCheck, checkedChSet)
	}

	// Read from channel
	for i := 0; i < len(urls); i++ {
		r := <-checkedChGet // Receive from channel
		checked[r.url] = r.online
	}

	return checked
}
