package scraper

import (
	"context"
	"time"

	"github.com/mmcdole/gofeed"

	gen "newscraper/internal/generator"
	sources "newscraper/internal/sources"
)

func Scrape(s sources.Source) ([]gen.Item, error) {
	parser := gofeed.NewParser()
	parser.Client = sources.Client
	parser.UserAgent = "newscrapper/1.0"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feed, err := parser.ParseURLWithContext(s.URL, ctx)
	if err != nil {
		return nil, err
	}

	const limit = 15
	var items []gen.Item
	for i, item := range feed.Items {
		if i >= limit {
			break
		}
		var date string
		if item.PublishedParsed != nil {
			date = item.PublishedParsed.Format("2006-01-02 15:04")
		}
		items = append(items, gen.Item{
			Title: item.Title,
			Link:  item.Link,
			Site:  feed.Title,
			Date:  date,
		})
	}
	return items, nil
}
