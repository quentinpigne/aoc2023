package main

import (
	"fmt"
	"quentinpigne/aoc2023/days"
	"quentinpigne/aoc2023/utils"
)

func main() {
	fileLines := utils.ReadFileLines("inputs/day2.txt")
	fmt.Println(days.ComputeDay2Part1(fileLines))
}
