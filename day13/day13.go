package main

import (
	"fmt"
)

func main() {
	input := readInput("day13/input.txt")
	fmt.Printf("Solution Day 13, Part 1: %d\n", SolvePuzzlePart1(input))
}

func SolvePuzzlePart1(input Input) int {
	instruct := input.Instructions[0]
	for i, point := range input.Paper {
		if point[instruct.Direction] > instruct.Count {
			input.Paper[i][instruct.Direction] = instruct.Count - (point[instruct.Direction] - instruct.Count)
		}
	}

	res := make(map[string]bool)
	for _, point := range input.Paper {
		res[fmt.Sprintf("%d%d", point["x"], point["y"])] = true
	}
	return len(res)
}
