package scraper

import (
	"fmt"
	"testing"
)

func TestScrape(t *testing.T) {
	testUrls := []string{
		"https://en.wikipedia.org/wiki/Child_development_stages",
	}

	Scrape(testUrls)
	fmt.Println(Scrape(testUrls))
}
