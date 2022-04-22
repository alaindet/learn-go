package main

type WebsiteChecker func(string) bool

func CheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {

	checked := make(map[string]bool, len(urls))

	for _, url := range urls {
		checked[url] = checker(url)
	}

	return checked
}
