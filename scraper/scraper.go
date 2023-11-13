package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type PageContent struct {
	URL   string
	Title string
	Text  string
}

func Scrape(urls []string) []PageContent {
	c := colly.NewCollector()

	var contents []PageContent // Slice to hold all scraped data

	tempTexts := make(map[string]string)

	// Scrape paragraph texts and store in tempTexts
	c.OnHTML("p", func(e *colly.HTMLElement) {
		url := e.Request.URL.String()
		cleanText := strings.TrimSpace(e.Text)
		if cleanText != "" {
			tempTexts[url] += cleanText + "\n\n" // Adding paragraphs to Text
		}
	})

	// Scrape titles
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		url := e.Request.URL.String()
		data := PageContent{
			Title: e.Text,
			URL:   url,
			Text:  tempTexts[url], // Retrieve text from tempTexts
		}
		contents = append(contents, data) // Append the data to the slice
	})

	for _, url := range urls {
		if err := c.Visit(url); err != nil {
			fmt.Printf("Error visiting URL: %s\n", err)
		}
	}

	return contents
}
