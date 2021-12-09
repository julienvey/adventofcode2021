package main

import (
	"fmt"
)

func main() {
	input := readInput("day9/input.txt")
	fmt.Printf("Solution Day 9, Part 1: %d\n", SolvePuzzlePart1(input))
}

func SolvePuzzlePart1(cave [][]int) int {
	var sum int
	for i := 0; i < len(cave); i++ {
		for j := 0; j < len(cave[i]); j++ {
			sum += isLowPoint(cave, i, j)
		}
	}
	return sum
}

func isLowPoint(cave [][]int, x, y int) int {
	val := cave[x][y]
	xmax := len(cave) - 1
	ymax := len(cave[0]) - 1
	if (x-1 < 0 || val < cave[x-1][y]) &&
		(y-1 < 0 || val < cave[x][y-1]) &&
		(x+1 > xmax || val < cave[x+1][y]) &&
		(y+1 > ymax || val < cave[x][y+1]) {
		return val + 1
	}
	return 0
}
