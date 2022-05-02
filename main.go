package main

import (
	"fmt"
	"log"
	"os"

	owasp "github.com/mamaoag/binosearch/services/owasp"
	scanner "github.com/mamaoag/binosearch/services/scanner"
	resource "github.com/mamaoag/binosearch/services/url"
	wordlist "github.com/mamaoag/binosearch/services/wordlist"
)

func main() {
	const APPNAME string = "Binoscan"
	var baseUrl string
	var wordlistPath string
	var endpointsFound []resource.Url
	var message string

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
		found := scanner.ScanEndpoint(url)

		if found.Path != "" {
			endpointsFound = append(endpointsFound, found)
		}
	}

	if len(endpointsFound) > 0 {
		fmt.Printf("\nScanning for OWASP API Security 2019 Issues\n")
		result, err := owasp.BrokenObjectLevelAuth(endpointsFound)

		if err != nil {
			log.Fatalln(err)
		}

		if result {
			message = "There are issues found. ❌"
		} else {
			message = "No issues found. ✅"
		}

		logResult("API1:2019", message)
	}

	fmt.Printf("Scanning Complete.\n")
}

func logResult(code string, message string) {
	log.Printf(
		"%d %d: [%s] %s\n",
		log.Ldate,
		log.Ltime,
		code,
		message,
	)
}
