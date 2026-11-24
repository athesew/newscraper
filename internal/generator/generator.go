package generator

import (
	"html/template"
	"os"
)

type Item struct {
	Title string
	Link  string
	Date  string
	Site  string
}

type Data struct {
	Title string
	Items []Item
}

func GenerateHTML(data Data, outputPath string) error {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
