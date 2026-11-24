package main

import (
	"fmt"
	"sync"

	gen "newscraper/internal/generator"
	scraper "newscraper/internal/scraper"
	sources "newscraper/internal/sources"

	"github.com/pkg/browser"
)

func main() {
	if err := sources.Load(); err != nil {
		fmt.Println("[error] не могу загрузить sources.json:", err)
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var gi []gen.Item

	for _, src := range sources.Sources {
		wg.Add(1)
		go func(s sources.Source) {
			defer wg.Done()
			items, err := scraper.Scrape(s)
			if err != nil {
				fmt.Println("[error]", err)
				return
			}
			mu.Lock()
			gi = append(gi, items...)
			mu.Unlock()
		}(src)
	}

	wg.Wait()

	data := gen.Data{
		Title: "newscrapper",
		Items: gi,
	}

	if err := gen.GenerateHTML(data, "news.html"); err != nil {
		fmt.Println("[error]", err)
		return
	}

	err := browser.OpenFile("news.html")
	if err != nil {
		fmt.Println("[error]", err)
	}

	fmt.Println("\033[33m[completed]\033[0m")
}
