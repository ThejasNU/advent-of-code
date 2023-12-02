package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	Day1_1()
}

func ReadInput(inputFile string) []string {
	file, err := os.Open(inputFile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}
