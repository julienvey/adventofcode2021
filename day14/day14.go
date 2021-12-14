package main

import (
	"fmt"
)

func main() {
	input := readInput("day14/input.txt")
	fmt.Printf("Solution Day 14, Part 1: %d\n", SolvePuzzle(input, 10))
	fmt.Printf("Solution Day 14, Part 2: %d\n", SolvePuzzle(input, 40))
}

func SolvePuzzle(input Input, steps int) int {
	count := map[string]int{}
	countChar := map[string]int{}
	for i := 0; i < len(input.Start)-1; i++ {
		count[fmt.Sprintf("%c%c", input.Start[i], input.Start[i+1])]++
		countChar[string(input.Start[i])]++
	}
	countChar[string(input.Start[len(input.Start)-1])]++

	for i := 0; i < steps; i++ {
		for key, val := range copy(count) {
			count[key] -= val
			mapValue := input.Map[key]
			countChar[mapValue] += val
			count[fmt.Sprintf("%c%s", key[0], mapValue)] += val
			count[fmt.Sprintf("%s%c", mapValue, key[1])] += val
		}
	}

	var mostCommon, leastCommon int
	for _, sum := range countChar {
		if sum > mostCommon {
			mostCommon = sum
		}
		if leastCommon == 0 || sum < leastCommon {
			leastCommon = sum
		}
	}
	return mostCommon - leastCommon
}

func copy(original map[string]int) map[string]int {
	newMap := map[string]int{}
	for k, v := range original {
		newMap[k] = v
	}
	return newMap
}
