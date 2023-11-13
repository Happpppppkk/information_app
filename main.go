package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"information_app/scraper"
)

type PageData struct {
	Title   string
	Content []scraper.PageContent
}

var scrapedContent []scraper.PageContent

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Infant Information App",
		Content: scrapedContent,
	}

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func contentHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("content_template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, content := range scrapedContent {
		err = tmpl.Execute(w, content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	urls := []string{
		"https://en.wikipedia.org/wiki/Child_development_stages",
		"https://en.wikipedia.org/wiki/Infant_feeding",
		"https://en.wikipedia.org/wiki/Infant_sleep",
		"https://en.wikipedia.org/wiki/Infant_clothing",
		"https://en.wikipedia.org/wiki/Infant_crying",
	}

	scrapedContent = scraper.Scrape(urls)
	//fmt.Println(scrapedContent)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/content/", contentHandler)
	fmt.Println("Starting the web server on localhost port 8080, localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
