package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

func StupidOutput(entries []Entry) {
	for _, v := range entries {
		fmt.Printf("%s\n%s\n", v.Term, v.Def)
	}
}

func NormalOutput(entries []Entry) {
	for _, v := range entries {
		fmt.Printf("\"%s\", \"%s\"\n", v.Term, v.Def)
	}
}

func JSONOutput(entries []Entry, i bool) {
	if i {
		j, err := json.MarshalIndent(entries, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", j)
	} else {
		j, err := json.Marshal(entries)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", j)
	}

}

func XMLOutput(entries []Entry, i bool) {
	if i {
		x, err := xml.MarshalIndent(entries, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", x)
	} else {
		x, err := xml.Marshal(entries)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", x)
	}
}
