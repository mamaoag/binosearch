package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"crypto/tls"
)

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func scanEndpoint(url string, path string) uint8 {

	var fullUrl string

	if strings.Contains(url, "http") {
		fullUrl = fmt.Sprintf("%s%s", url, path)
	} else {
		fullUrl = fmt.Sprintf("https://%s%s", url, path)
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(fullUrl)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 404 {
		log.Printf("%d %d: GET %s: Status - %d\n", log.Ldate, log.Ltime, path, resp.StatusCode)
		return 0
	}

	log.Printf("%s %d %d: GET %s: Status - %d%s\n", "\033[32m", log.Ldate, log.Ltime, path, resp.StatusCode, "\033[0m")
	return 1
}

func main() {
	const APPNAME string = "Binoscan"
	var baseUrl string
	var resultsFound uint8 = 0

	fmt.Printf("%s - an api application scanner.\n", APPNAME)
	fmt.Print("Enter your base url: ")

	fmt.Scan(&baseUrl)

	paths, err := readFile("api-wordlist.txt")

	if err != nil {
		log.Fatalln(err)
	}

	for _, path := range paths {
		resultsFound = resultsFound + scanEndpoint(baseUrl, path)
	}

	fmt.Printf("Scanning Complete. There are %d routes are available \n", resultsFound)
}
