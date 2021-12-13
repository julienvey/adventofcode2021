package main

import (
	"fmt"
	"github.com/julienvey/adventofcode2021/utils"
)

func main() {
	input := readInput("day13/input.txt")
	fmt.Printf("Solution Day 13, Part 1: %d\n", SolvePuzzlePart1(input))
	fmt.Println("Solution Day 13, Part 2:")
	SolvePuzzlePart2(input)
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

func SolvePuzzlePart2(input Input) {
	for _, instruct := range input.Instructions {
		for i, point := range input.Paper {
			if point[instruct.Direction] > instruct.Count {
				input.Paper[i][instruct.Direction] = instruct.Count - (point[instruct.Direction] - instruct.Count)
			}
		}
	}

	var maxX, maxY int
	for _, point := range input.Paper {
		maxX = utils.Max(point["x"], maxX)
		maxY = utils.Max(point["y"], maxY)
	}
	res := make([][]bool, maxX+1)
	for i, _ := range res {
		res[i] = make([]bool, maxY+1)
	}

	for _, point := range input.Paper {
		res[point["x"]][point["y"]] = true
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if res[x][y] {
				fmt.Printf(" #")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
}
