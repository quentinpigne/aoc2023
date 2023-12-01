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
	return 0
}

// Utilities Functions

func BoundaryDigits(fileLine string) string {
	onlyDigits := utils.Filter[rune]([]rune(fileLine), utils.IsDigit)
	return fmt.Sprintf("%c%c", onlyDigits[0], onlyDigits[len(onlyDigits)-1])
}
