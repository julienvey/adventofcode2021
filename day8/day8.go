package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input := readInput("day8/input.txt")
	fmt.Printf("Solution Day 8, Part 1: %d\n", SolvePuzzlePart1(input))
}

type Entry struct {
	Inputs  []string
	Outputs []string
}

func readInput(in string) []Entry {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	var entries []Entry
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		item := strings.Split(line, "|")
		if len(item) != 2 {
			log.Fatal("entry should have two items, separated by |")
		}

		var entry Entry
		for _, signal := range strings.Split(strings.TrimSpace(item[0]), " ") {
			entry.Inputs = append(entry.Inputs, strings.TrimSpace(signal))
		}
		for _, digit := range strings.Split(strings.TrimSpace(item[1]), " ") {
			entry.Outputs = append(entry.Outputs, strings.TrimSpace(digit))
		}
		entries = append(entries, entry)
	}
	return entries
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
