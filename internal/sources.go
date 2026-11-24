package scraper

import (
	"net/http"
	"time"
)

type Source struct {
	Name string
	URL  string
}

var Sources = []Source{
	{Name: "SecurityLab", URL: "https://www.securitylab.ru/_Services/Export/RSS/"},
	{Name: "Habr News", URL: "https://habr.com/ru/rss/news/?fl=ru"},
	{Name: "Xakep", URL: "https://xakep.ru/feed/"},
}

var Client = &http.Client{Timeout: 5 * time.Second}
