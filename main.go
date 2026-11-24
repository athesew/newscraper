package main

import (
	"sync"

	scraper "newscraper/internal/scraper"
	sources "newscraper/internal/sources"
)

func main() {
	var wg sync.WaitGroup

	for _, src := range sources.Sources {
		wg.Add(1)
		go func(s sources.Source) {
			defer wg.Done()
			scraper.Scrape(s)
		}(src)
	}

	wg.Wait()
}
