package main

import (
	"fmt"
	"quentinpigne/aoc2023/days"
	"quentinpigne/aoc2023/utils"
)

func main() {
	fileLines := utils.ReadFileLines("inputs/day4.txt")
	fmt.Println(days.ComputeDay4Part1(fileLines))
}
