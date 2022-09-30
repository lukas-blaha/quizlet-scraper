package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Entry struct {
	Term string `json:"term" xml:"term"`
	Def  string `json:"definition" xml:"definition"`
}

func GetEntries(url string) ([]Entry, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8"},
		"cache-control":             {"max-age=0"},
		"upgrade-insecure-requests": {"1"},
		"user-agent":                {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"},
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var entry Entry
	var entries []Entry

	doc.Find("div.SetPageTerms-term").Each(func(i int, s *goquery.Selection) {
		entry.Term = s.Find("a[class=SetPageTerm-wordText]").Text()
		entry.Def = s.Find("a[class=SetPageTerm-definitionText]").Text()
		entries = append(entries, entry)
	})

	return entries, nil
}
