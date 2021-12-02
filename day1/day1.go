package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	fmt.Printf("Solution Day 1, Part 1: %d\n", SolvePuzzlePart1(input))
	fmt.Printf("Solution Day 1, Part 2: %d\n", SolvePuzzlePart2(input))
}

func readInput() []int {
	file, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var input []int
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		inputInt, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, inputInt)
	}
	return input
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
