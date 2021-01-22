package concurrency

// WebsiteChecker checks a given website url and returns whether it passed or failed the check
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites runs the given website checker on each of the given urls and returns a map of url to result
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool, len(urls))
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultsChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultsChannel
		results[result.string] = result.bool
	}

	return results
}
