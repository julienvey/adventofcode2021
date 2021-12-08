package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := readInput("day8/input.txt")
	fmt.Printf("Solution Day 8, Part 1: %d\n", SolvePuzzlePart1(input))
	fmt.Printf("Solution Day 8, Part 2: %d\n", SolvePuzzlePart2(input))
}

func SolvePuzzlePart1(entries []Entry) int {
	var sum int
	for _, entry := range entries {
		for _, output := range entry.Outputs {
			switch len(output) {
			case 2, 3, 4, 7:
				sum++
			}
		}
	}
	return sum
}

func SolvePuzzlePart2(entries []Entry) int {
	var sum int
	index := make(map[int]string, 8)
	indexInt := make(map[int]string, 4)
	for _, entry := range entries {
		for _, input := range entry.Inputs {
			switch l := len(input); l {
			case 2:
				index[1] = input
				indexInt[l] = "1"
			case 3:
				index[7] = input
				indexInt[l] = "7"
			case 4:
				index[4] = input
				indexInt[l] = "4"
			case 7:
				index[8] = input
				indexInt[l] = "8"
			}
		}

		var fourDigit string
		for _, output := range entry.Outputs {
			switch l := len(output); l {
			case 5: // Two, Three or Five
				if containsAllChars(output, index[1]) {
					fourDigit += "3"
				} else if countCommonChars(output, index[4]) == 3 {
					fourDigit += "5"
				} else {
					fourDigit += "2"
				}
			case 6: // Zero or Six, Nine
				if containsAllChars(output, index[4]) {
					fourDigit += "9"
				} else if !containsAllChars(output, index[1]) {
					fourDigit += "6"
				} else {
					fourDigit += "0"
				}
			default:
				fourDigit += indexInt[l]
			}

		}
		outputInt, err := strconv.Atoi(fourDigit)
		if err != nil {
			log.Fatal(err)
		}
		sum += outputInt
	}
	return sum
}

func containsAllChars(input, index string) bool {
	for _, c := range index {
		if !strings.Contains(input, string(c)) {
			return false
		}
	}
	return true
}

func countCommonChars(a, b string) int {
	var sum int
	for _, ca := range a {
		for _, cb := range b {
			if ca == cb {
				sum++
			}
		}
	}
	return sum
}
