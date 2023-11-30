package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}

	return fileLines
}
