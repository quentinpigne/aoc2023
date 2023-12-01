package days

import (
	"fmt"
	"quentinpigne/aoc2023/utils"
	"strconv"
	"strings"
)

// Parts Computation

func ComputeDay1Part1(fileLines []string) int {
	count := 0
	for _, fileLine := range fileLines {
		intValue, _ := strconv.Atoi(BoundaryDigits(fileLine))
		count += intValue
	}
	return count
}

func ComputeDay1Part2(fileLines []string) int {
	count := 0
	for _, fileLine := range fileLines {
		convertedLine := ReplaceLiteralNumbers(fileLine)
		intValue, _ := strconv.Atoi(BoundaryDigits(convertedLine))
		count += intValue
	}
	return count
}

// Utilities Functions

func BoundaryDigits(fileLine string) string {
	onlyDigits := utils.Filter[rune]([]rune(fileLine), utils.IsDigit)
	return fmt.Sprintf("%c%c", onlyDigits[0], onlyDigits[len(onlyDigits)-1])
}

func ReplaceLiteralNumbers(fileLine string) string {
	firstReplacer := strings.NewReplacer("twone", "21", "oneight", "18", "threeight", "38", "fiveight", "58", "eightwo", "82", "eighthree", "83", "sevenine", "79", "nineight", "98")
	secondReplacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
	return secondReplacer.Replace(firstReplacer.Replace(fileLine))
}
