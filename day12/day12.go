package main

import (
	"fmt"
	"strings"
)

func main() {
	input := readInput("day12/input.txt")
	fmt.Printf("Solution Day 12, Part 1: %d\n", SolvePuzzlePart1(input))
	fmt.Printf("Solution Day 12, Part 1: %d\n", SolvePuzzlePart2(input))
}

const start = "start"
const end = "end"

func SolvePuzzlePart1(input []Entry) int {
	pathMap := map[string][]string{}
	for _, entry := range input {
		pathMap[entry.Source] = append(pathMap[entry.Source], entry.Dest)
		pathMap[entry.Dest] = append(pathMap[entry.Dest], entry.Source)
	}

	return len(growPath("start,", start, pathMap))
}

func growPath(path, current string, pathMap map[string][]string) []string {
	var paths []string
	for _, next := range pathMap[current] {
		switch {
		case next == end:
			paths = append(paths, path+end)
		case next == start:
			continue
		case strings.ToLower(next) == next:
			if !strings.Contains(path, ","+next+",") {
				paths = append(paths, growPath(path+next+",", next, pathMap)...)
			}
		case strings.ToUpper(next) == next:
			paths = append(paths, growPath(path+next+",", next, pathMap)...)

		}
	}
	return paths
}

func SolvePuzzlePart2(input []Entry) int {
	pathMap := map[string][]string{}
	for _, entry := range input {
		pathMap[entry.Source] = append(pathMap[entry.Source], entry.Dest)
		pathMap[entry.Dest] = append(pathMap[entry.Dest], entry.Source)
	}

	return len(growPath2("start,", start, pathMap, false))
}

func growPath2(path, current string, pathMap map[string][]string, usedSingle bool) []string {
	var paths []string
	for _, next := range pathMap[current] {
		switch {
		case next == end:
			paths = append(paths, path+end)
		case next == start:
			continue
		case strings.ToLower(next) == next:
			if !strings.Contains(path, ","+next+",") {
				paths = append(paths, growPath2(path+next+",", next, pathMap, usedSingle)...)
			} else if !usedSingle {
				paths = append(paths, growPath2(path+next+",", next, pathMap, true)...)
			}
		case strings.ToUpper(next) == next:
			paths = append(paths, growPath2(path+next+",", next, pathMap, usedSingle)...)

		}
	}
	return paths
}
