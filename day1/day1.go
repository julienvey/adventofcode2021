package main

import (
	"fmt"
	"github.com/julienvey/adventofcode2021/utils"
)

func main() {
	input := utils.ReadInput("day1/input.txt")
	fmt.Printf("Solution Day 1, Part 1: %d\n", SolvePuzzlePart1(input))
	fmt.Printf("Solution Day 1, Part 2: %d\n", SolvePuzzlePart2(input))
}

func SolvePuzzlePart2(input []int) int {
	if len(input) < 2 {
		return 0
	}

	var sum, localSum, last int
	for i := 1; i < len(input)-1; i++ {
		localSum = input[i-1] + input[i] + input[i+1]
		if last != 0 && localSum > last {
			sum++
		}
		last = localSum
	}
	return sum
}

func SolvePuzzlePart1(input []int) int {
	if len(input) < 2 {
		return 0
	}

	var sum, last int
	for i := 0; i < len(input)-1; i++ {
		last = input[i]
		if input[i+1] > last {
			sum++
		}
	}
	return sum
}
