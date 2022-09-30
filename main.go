package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Entry struct {
	Term string `json:"term" xml:"term"`
	Def  string `json:"definition" xml:"definition"`
}

func main() {
	scrapeURL := flag.String("url", "", "quizlet URL to be parsed")
	jsonFormat := flag.Bool("json", false, "sets output format to json")
	xmlFormat := flag.Bool("xml", false, "sets output format to xml")
	stupidFormat := flag.Bool("stupid", false, "sets output format to stupid")
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if !strings.Contains(*scrapeURL, "https://quizlet.com/") {
		fmt.Println("URL must be study set from quizlet.com page!")
		os.Exit(1)
	}

	entries, err := getEntries(*scrapeURL)
	if err != nil {
		log.Fatal(err)
	}

	if *jsonFormat {
		jsonOutput(entries)
	} else if *xmlFormat {
		xmlOutput(entries)
	} else if *stupidFormat {
		stupidOutput(entries)
	} else {
		normalOutput(entries)
	}
}

func stupidOutput(entries []Entry) {
	for _, v := range entries {
		fmt.Printf("%s\n%s\n", v.Term, v.Def)
	}
}

func normalOutput(entries []Entry) {
	for _, v := range entries {
		fmt.Printf("\"%s\", \"%s\"\n", v.Term, v.Def)
	}
}

func jsonOutput(entries []Entry) {
	j, err := json.Marshal(entries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", j)
}

func xmlOutput(entries []Entry) {
	j, err := xml.Marshal(entries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", j)
}

func getEntries(url string) ([]Entry, error) {
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
