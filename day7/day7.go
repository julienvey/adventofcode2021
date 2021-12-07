package main

import (
	"fmt"
	"github.com/julienvey/adventofcode2021/utils"
	"math"
	"sort"
)

func main() {
	input := utils.ReadInput("day7/input.txt")
	fmt.Printf("Solution Day 7, Part 1: %d\n", SolvePuzzlePart1(input))
}

func SolvePuzzlePart1(crabs []int) int {
	median := CalcMedian(crabs...)
	var fuel int
	for _, crab := range crabs {
		fuel += int(math.Abs(float64(median - crab)))
	}
	return fuel
}

func CalcMedian(crabs ...int) int {
	l := len(crabs)
	sort.Ints(crabs)
	if l%2 == 0 {
		return (crabs[l/2-1] + crabs[l/2]) / 2
	}
	return crabs[l/2]
}
