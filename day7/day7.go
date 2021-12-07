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
	fmt.Printf("Solution Day 7, Part 2: %d\n", SolvePuzzlePart2(input))
}

func SolvePuzzlePart1(crabs []int) int {
	median := CalcMedian(crabs...)
	var fuel int
	for _, crab := range crabs {
		fuel += int(math.Abs(float64(median - crab)))
	}
	return fuel
}

func SolvePuzzlePart2(crabs []int) int {
	avg := CalcAverage(crabs...)
	return utils.Min(getFuel(crabs, avg), getFuel(crabs, avg+1))
}

func getFuel(crabs []int, avg int) int {
	var fuel int
	for _, crab := range crabs {
		diff := int(math.Abs(float64(avg - crab)))
		fuel += (diff * (diff + 1)) / 2
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

func CalcAverage(crabs ...int) int {
	var sum int
	for _, crab := range crabs {
		sum += crab
	}
	return sum / len(crabs)
}
