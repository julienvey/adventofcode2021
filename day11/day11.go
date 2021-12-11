package main

import (
	"fmt"
)

func main() {
	input := readInput("day11/input.txt")
	fmt.Printf("Solution Day 11, Part 1: %d\n", SolvePuzzlePart1(input))
}

func SolvePuzzlePart1(input [][]int) int {
	var sumFlashes int
	for step := 0; step < 100; step++ {
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				input[x][y] += 1
			}
		}

		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				sumFlashes += flashes(&input, x, y)
			}
		}
	}
	return sumFlashes
}

func flashes(input *[][]int, x int, y int) int {
	if (*input)[x][y] > 9 {
		(*input)[x][y] = 0
		return increaseNearby(input, x, y) + 1
	}
	return 0
}

func increaseNearby(input *[][]int, x int, y int) int {
	var sum int
	for xnext := x - 1; xnext <= x+1; xnext++ {
		for ynext := y - 1; ynext <= y+1; ynext++ {
			if (xnext == x && ynext == y) ||
				(xnext < 0 || xnext > 9 || ynext < 0 || ynext > 9) ||
				((*input)[xnext][ynext] == 0) {
				continue
			}
			(*input)[xnext][ynext] += 1
			sum += flashes(input, xnext, ynext)
		}
	}
	return sum
}
