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
	fmt.Printf("Solution Day 3, Part 1: %d\n", SolvePuzzlePart1(input))
	fmt.Printf("Solution Day 3, Part 2: %d\n", SolvePuzzlePart2(input))
}

func readInput() []string {
	file, err := ioutil.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var input []string
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		input = append(input, line)
	}
	return input
}

func SolvePuzzlePart1(input []string) int64 {
	zero := make([]int, len(input[0]))
	one := make([]int, len(input[0]))
	for _, in := range input {
		if len(in) != len(zero) {
			log.Fatalf("wrong input size: %d, expected :%d\n", len(in), len(zero))
		}
		for pos, char := range in {
			switch char {
			case '0':
				zero[pos]++
			case '1':
				one[pos]++
			default:
				log.Fatalf("%c not valid input char", char)
			}
		}
	}
	var gamma, epsilon string
	for i := 0; i < len(zero); i++ {
		if zero[i] > one[i] {
			gamma = fmt.Sprintf("%s%c", gamma, '1')
			epsilon = fmt.Sprintf("%s%c", epsilon, '0')
		} else {
			gamma = fmt.Sprintf("%s%c", gamma, '0')
			epsilon = fmt.Sprintf("%s%c", epsilon, '1')
		}
	}
	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilonInt, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return gammaInt * epsilonInt
}

func SolvePuzzlePart2(input []string) int64 {
	return getRating(input, 0, oxygenCompare) * getRating(input, 0, co2Compare)
}

type compare func(map[uint8]int) uint8

func getRating(input []string, pos int, fn compare) int64 {
	var mostCommon = make(map[uint8]int, 2)
	for _, in := range input {
		mostCommon[in[pos]]++
	}

	filter := fn(mostCommon)
	var filtered []string
	for _, in := range input {
		if in[pos] == filter {
			filtered = append(filtered, in)
		}
	}
	if len(filtered) == 1 {
		i, err := strconv.ParseInt(filtered[0], 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		return i
	} else {
		return getRating(filtered, pos+1, fn)
	}
}

func oxygenCompare(mostCommon map[uint8]int) uint8 {
	if mostCommon['1'] < mostCommon['0'] {
		return '0'
	}
	return '1'
}

func co2Compare(mostCommon map[uint8]int) uint8 {
	if mostCommon['0'] > mostCommon['1'] {
		return '1'
	}
	return '0'
}
