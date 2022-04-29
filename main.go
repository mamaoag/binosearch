package main

import (
	"fmt"
	"log"
	"os"

	scanner "github.com/mamaoag/binosearch/services/scanner"
	resource "github.com/mamaoag/binosearch/services/url"
	wordlist "github.com/mamaoag/binosearch/services/wordlist"
)

func main() {
	const APPNAME string = "Binoscan"
	var baseUrl string
	var wordlistPath string
	// var resultsFound uint8 = 0

	fmt.Printf("%s - an api application scanner.\n", APPNAME)
	fmt.Print("Enter your base url > ")
	fmt.Scan(&baseUrl)

	if len(os.Args) == 1 {
		wordlistPath = wordlist.DEFAULT
	} else {
		wordlistPath = os.Args[1]
	}

	dir, err := wordlist.SetWordlist(wordlistPath)

	if err != nil {
		log.Fatalln(err)
	}

	for _, path := range dir {
		url := resource.Parse(baseUrl, path)
		scanner.ScanEndpoint(url)
	}

	fmt.Printf("Scanning Complete.\n")
}
