package services

import (
	"bufio"
	"os"
)

const DEFAULT = "./wordlist/default.txt"

func SetWordlist(path string) ([]string, error) {

	wordlist, err := readFile(path)

	if err != nil {
		return nil, err
	}

	return wordlist, err
}

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
