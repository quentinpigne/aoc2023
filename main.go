package main

import (
	"fmt"
	"quentinpigne/aoc2023/days"
	"quentinpigne/aoc2023/utils"
)

func main() {
	fileLines := utils.ReadFileLines("inputs/day1.txt")
	fmt.Println(days.ComputeDay1Part2(fileLines))
}
