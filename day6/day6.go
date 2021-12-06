package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := readInput("day6/input.txt")
	fmt.Printf("Solution Day 6, Part 1: %d\n", SolvePuzzlePart1(input))
}

func readInput(in string) []int {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	var input []int
	for _, line := range strings.Split(string(file), ",") {
		if line == "" {
			continue
		}
		inputInt, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, inputInt)
	}
	return input
}

func SolvePuzzlePart1(input []int) int {
	for d := 0; d < 80; d++ {
		var dayGrowth int
		for i := 0; i < len(input); i++ {
			var grow int
			input[i], grow = growSquid(input[i])
			dayGrowth += grow
		}
		for i := 0; i < dayGrowth; i++ {
			input = append(input, 8)
		}
	}
	return len(input)
}

func growSquid(squid int) (int, int) {
	if squid == 0 {
		return 6, 1
	}
	return squid - 1, 0
}
