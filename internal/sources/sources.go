package sources

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type Source struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

const config = "sources.json"

var Sources = []Source{
	{Name: "SecurityLab", URL: "https://www.securitylab.ru/_Services/Export/RSS/"},
	{Name: "Habr News", URL: "https://habr.com/ru/rss/news/?fl=ru"},
	{Name: "Xakep", URL: "https://xakep.ru/feed/"},
}

var Client = &http.Client{Timeout: 5 * time.Second}

func Load() error {
	raw, err := os.ReadFile(config)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var loaded []Source
	if err := json.Unmarshal(raw, &loaded); err != nil {
		return err
	}

	if len(loaded) > 0 {
		Sources = loaded
	}
	return nil
}
