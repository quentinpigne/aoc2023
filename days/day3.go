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
	for i := 0; i < len(fileLines); i++ {
		for _, j := range findAllGears(fileLines[i]) {
			gearRatio := 1
			numberOfMatch := 0
			frame := getFrame(fileLines[i], j)
			if i > 0 {
				for _, k := range findAllNumbers(fileLines[i-1]) {
					value, _ := strconv.Atoi(fileLines[i-1][k[0]:k[1]])
					if numberInFrame(k, frame) {
						gearRatio *= value
						numberOfMatch += 1
					}
				}
			}
			for _, k := range findAllNumbers(fileLines[i]) {
				value, _ := strconv.Atoi(fileLines[i][k[0]:k[1]])
				if numberInFrame(k, frame) {
					gearRatio *= value
					numberOfMatch += 1
				}
			}
			if i < len(fileLines)-1 {
				for _, k := range findAllNumbers(fileLines[i+1]) {
					value, _ := strconv.Atoi(fileLines[i+1][k[0]:k[1]])
					if numberInFrame(k, frame) {
						gearRatio *= value
						numberOfMatch += 1
					}
				}
			}
			if numberOfMatch == 2 {
				count += gearRatio
			}
		}
	}
	return count
}

// Utilities Functions

func findAllNumbers(sequence string) [][]int {
	re := regexp.MustCompile("(\\d+)")
	return re.FindAllStringIndex(sequence, -1)
}

func findAllGears(sequence string) [][]int {
	re := regexp.MustCompile("([*])")
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

func numberInFrame(k []int, frame []int) bool {
	return (frame[0] <= k[0] && frame[1]-1 >= k[0]) || (frame[0] <= k[1]-1 && frame[1]-1 >= k[1]-1)
}

func containsSymbol(sequence string) bool {
	re := regexp.MustCompile("[^.\\d]")
	return re.MatchString(sequence)
}

func containsDigit(sequence string) bool {
	re := regexp.MustCompile("\\d")
	return re.MatchString(sequence)
}
