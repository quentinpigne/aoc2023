package days

import (
	"regexp"
	"strconv"
)

// Parts Computation

func ComputeDay3Part1(fileLines []string) int {
	count := 0
	for i := 0; i < len(fileLines); i++ {
		for _, j := range findAllNumbers(fileLines[i]) {
			frame := getFrame(fileLines[i], j)
			value, _ := strconv.Atoi(fileLines[i][j[0]:j[1]])
			if i > 0 && containsSymbol(fileLines[i-1][frame[0]:frame[1]]) {
				count += value
				continue
			}
			if containsSymbol(fileLines[i][frame[0]:frame[1]]) {
				count += value
				continue
			}
			if i < len(fileLines)-2 && containsSymbol(fileLines[i+1][frame[0]:frame[1]]) {
				count += value
				continue
			}
		}
	}
	return count
}

func ComputeDay3Part2(fileLines []string) int {
	count := 0
	return count
}

// Utilities Functions

func findAllNumbers(sequence string) [][]int {
	re := regexp.MustCompile("(\\d+)")
	return re.FindAllStringIndex(sequence, -1)
}

func getFrame(line string, index []int) []int {
	startPos := index[0]
	endPos := index[1]
	if startPos > 0 {
		startPos -= 1
	}
	if endPos < len(line)-2 {
		endPos += 1
	}
	return []int{startPos, endPos}
}

func containsSymbol(sequence string) bool {
	re := regexp.MustCompile("[^.\\d]")
	return re.MatchString(sequence)
}
