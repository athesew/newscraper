package scraper

import (
	"context"
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

func Scrape(s Source) {
	parser := gofeed.NewParser()
	parser.Client = Client
	parser.UserAgent = "newscrapper/1.0"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feed, err := parser.ParseURLWithContext(s.URL, ctx)
	if err != nil {
		fmt.Printf("\033[33m[error]\033[0m %v\n", err)
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
		fmt.Printf("\033[35m[site]\033[0m %s\n\n", feed.Title)
	}
}
