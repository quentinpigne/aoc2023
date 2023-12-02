package days

import (
	"quentinpigne/aoc2023/utils"
	"strconv"
	"strings"
)

// Parts Computation

func ComputeDay2Part1(fileLines []string) int {
	count := 0
	for _, fileLine := range fileLines {
		gameId, gameSets := getGameSetsWithId(fileLine)
		count += gameId
		for _, gameSet := range gameSets {
			if (!isSetValid(gameSet)) {
				count -= gameId
				break
			}
		}
	}
	return count
}

func ComputeDay2Part2(fileLines []string) int {
	return 0
}

// Utilities Functions

func getGameSetsWithId(gameLine string) (int, [][]string) {
	game := strings.Split(gameLine, ": ")
	gameId, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
	gameSets := utils.Map[string, []string](strings.Split(game[1], "; "), func(s string) []string {
		return strings.Split(s, ", ")
	})

	return gameId, gameSets
}

func isSetValid(gameSet []string) bool {
	for _, cubeSet := range gameSet {
		cubeColor := strings.Split(cubeSet, " ")[1]
		cubeCount, _ := strconv.Atoi(strings.Split(cubeSet, " ")[0])
		switch cubeColor {
			case "red":
				if (cubeCount > 12) {
					return false
				}
				break
			case "green":
				if (cubeCount > 13) {
					return false
				}
				break
			case "blue":
				if (cubeCount > 14) {
					return false
				}
				break
		}
	}
	return true
}
