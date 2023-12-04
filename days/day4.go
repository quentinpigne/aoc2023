package days

import (
	"quentinpigne/aoc2023/utils"
	"slices"
	"strconv"
	"strings"
)

// Parts Computation

func ComputeDay4Part1(fileLines []string) int {
	count := 0
	for _, fileLine := range fileLines {
		cardScore := 0
		_, cardWinningNumbers, cardNumbers := getCardSetWithId(fileLine)
		for _, winningNumber := range cardWinningNumbers {
			if slices.Contains(cardNumbers, winningNumber) {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore *= 2
				}
			}
		}
		count += cardScore
	}
	return count
}

func ComputeDay4Part2(fileLines []string) int {
	count := 0
	return count
}

// Utilities Functions

func getCardSetWithId(line string) (int, []int, []int) {
	card := strings.Split(line, ": ")
	cardId, _ := strconv.Atoi(strings.Split(card[0], " ")[1])
	cardWinningNumbers := decomposeToInt(strings.Split(card[1], " | ")[0])
	cardNumbers := decomposeToInt(strings.Split(card[1], " | ")[1])
	return cardId, cardWinningNumbers, cardNumbers
}

func decomposeToInt(s string) []int {
	splittedString := strings.Split(s, " ")
	notEmtpySplittedString := utils.Filter[string](splittedString, func(s string) bool {
		return s != ""
	})
	return utils.Map[string, int](notEmtpySplittedString, func(s string) int {
		value, _ := strconv.Atoi(s)
		return value
	})
}
