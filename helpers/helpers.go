package helpers

import (
	"bufio"
	"log"
	"os"
)

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