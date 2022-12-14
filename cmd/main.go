package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scrapeURL := flag.String("url", "", "quizlet URL to be parsed")
	jsonFormat := flag.Bool("json", false, "sets output format to json")
	xmlFormat := flag.Bool("xml", false, "sets output format to xml")
	stupidFormat := flag.Bool("stupid", false, "sets output format to stupid")
	indent := flag.Bool("pretty", false, "enables indentation for json/xml marshalling")
	fileOutput := flag.String("file", "", "file to be data written to")
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if !strings.Contains(*scrapeURL, "https://quizlet.com/") {
		fmt.Println("URL must be study set from quizlet.com page!")
		os.Exit(1)
	}

	entries, err := GetEntries(*scrapeURL)
	if err != nil {
		log.Fatal(err)
	}

	if *indent && (!*jsonFormat && !*xmlFormat) {
		fmt.Println("\"-indent\" can be used only with -json/-xml option")
		os.Exit(1)
	}

	out := os.Stdout

	if *fileOutput != "" {
		out = FileOutput(*fileOutput)
	}

	switch {
	case *jsonFormat:
		JSONOutput(entries, *indent, out)
	case *xmlFormat:
		XMLOutput(entries, *indent, out)
	case *stupidFormat:
		StupidOutput(entries, out)
	default:
		NormalOutput(entries, out)
	}
}
