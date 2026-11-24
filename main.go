package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/mmcdole/gofeed"
)

type source struct {
	Name string
	URL  string
}

var sources = []source{
	{Name: "SecurityLab", URL: "https://www.securitylab.ru/_Services/Export/RSS/"},
	{Name: "Habr News", URL: "https://habr.com/ru/rss/news/?fl=ru"},
	{Name: "Xakep", URL: "https://xakep.ru/feed/"},
}

var client = &http.Client{Timeout: 5 * time.Second}

func main() {
	var wg sync.WaitGroup

	for _, src := range sources {
		wg.Add(1)
		go func(s source) {
			defer wg.Done()
			scrape(s)
		}(src)
	}

	wg.Wait()
}

func scrape(s source) {
	parser := gofeed.NewParser()
	parser.Client = client
	parser.UserAgent = "newscrapper/1.0"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feed, err := parser.ParseURLWithContext(s.URL, ctx)
	if err != nil {
		fmt.Printf("\033[33m[ error ]\033[0m %v\n", err)
		return
	}

	const limit = 15
	for i, item := range feed.Items {
		if i >= limit {
			break
		}
		fmt.Printf("\033[36m[info]\033[0m %s\n", item.Title)
		fmt.Printf("\033[34m[src]\033[0m %s\n", item.Link)
		if item.PublishedParsed != nil {
			fmt.Printf("\033[33m[date]\033[0m %s\n", item.PublishedParsed.Format("2006-01-02 15:04"))
		}
		fmt.Printf("\033[35m[site]\033[0m %s", feed.Title)
		fmt.Println("")
		fmt.Println("")
	}
}
