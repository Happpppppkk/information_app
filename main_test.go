package main

import (
	"information_app/scraper"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	homeHandler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusOK)
	}
}

func TestContentHandler(t *testing.T) {
	scrapedContent = []scraper.PageContent{{Title: "Test", Text: "Test Test"}}

	req := httptest.NewRequest(http.MethodGet, "/content/Test", nil)
	res := httptest.NewRecorder()

	contentHandler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("got status %d but wanted %d for existing content", res.Code, http.StatusOK)
	}
}

//https://quii.gitbook.io/learn-go-with-tests/questions-and-answers/http-handlers-revisited
