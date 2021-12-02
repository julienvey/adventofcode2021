package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type command struct {
	action string
	units  int
}

func main() {
	input := readInput()
	fmt.Printf("Solution Day 2, Part 1: %d\n", SolvePuzzlePart1(input))
	fmt.Printf("Solution Day 2, Part 2: %d\n", SolvePuzzlePart2(input))
}

func readInput() []command {
	file, err := ioutil.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var input []command
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		item := strings.Split(line, " ")
		if len(item) != 2 {
			log.Fatal("command should have two items, action and units")
		}

		itemUnits, err := strconv.Atoi(item[1])
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, command{
			action: item[0],
			units:  itemUnits,
		})
	}
	return input
}

func SolvePuzzlePart1(input []command) int {
	var pos, depth int
	for _, cmd := range input {
		switch cmd.action {
		case "forward":
			pos += cmd.units
		case "down":
			depth += cmd.units
		case "up":
			depth -= cmd.units
		}
	}
	return depth * pos
}

func SolvePuzzlePart2(input []command) int {
	var pos, depth, aim int
	for _, cmd := range input {
		switch cmd.action {
		case "forward":
			pos += cmd.units
			depth += aim * cmd.units
		case "down":
			aim += cmd.units
		case "up":
			aim -= cmd.units
		}
	}
	return depth * pos
}
