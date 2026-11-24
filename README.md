<p align="center">
<img src="icon/icon.jpg" alt="RSSHub" width="100">
</p>
<h1 align="center">newscraper</h1>
<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/license-MIT-red?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/build-passing-brightgreen?style=flat-square" alt="Build">
</p>

### introduction

newscraper is a command-line utility written in Go that concurrently gathers news from multiple RSS/XML feeds and generates a single-page, HTML document. the program provides:

- simple addition of sources
- a standard, minimalist, auto-generated website
- concurrency, which allows for even faster news scraping

newscraper perfectly suitable for daily use

### how to use

make changes to `sources.json` like this
```json
[
  { "name": "SecurityLab", "url": "https://www.securitylab.ru/_Services/Export/RSS/" },
  { "name": "Habr News", "url": "https://habr.com/ru/rss/news/" },
  { "name": "Xakep", "url": "https://xakep.ru/feed/" }
]
```
then simply run the .go file
```bash
go run cmd/main.go
```

### how to build

the building presents no difficulties
```bash
cd cmd/
go build -o newscraper
cd ..
./cmd/newscraper
```
if you want to run the binary using `./newscraper`, you need to either move `template.html` to the same location as the binary or embed it inside the binary