package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
)

func StupidOutput(entries []Entry, o io.Writer) {
	for _, v := range entries {
		fmt.Fprintf(o, "%s\n%s\n", v.Term, v.Def)
	}
}

func NormalOutput(entries []Entry, o io.Writer) {
	for _, v := range entries {
		fmt.Fprintf(o, "\"%s\", \"%s\"\n", v.Term, v.Def)
	}
}

func JSONOutput(entries []Entry, i bool, o io.Writer) {
	if i {
		j, err := json.MarshalIndent(entries, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(o, "%s\n", j)
	} else {
		j, err := json.Marshal(entries)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(o, "%s\n", j)
	}

}

func XMLOutput(entries []Entry, i bool, o io.Writer) {
	if i {
		x, err := xml.MarshalIndent(entries, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(o, "%s\n", x)
	} else {
		x, err := xml.Marshal(entries)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(o, "%s\n", x)
	}
}
