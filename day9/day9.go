package main

import (
	"fmt"
	"sort"
)

func main() {
	input := readInput("day9/input.txt")
	fmt.Printf("Solution Day 9, Part 1: %d\n", SolvePuzzlePart1(input))
	fmt.Printf("Solution Day 9, Part 2: %d\n", SolvePuzzlePart2(input))
}

func SolvePuzzlePart1(cave [][]int) int {
	var sum int
	for x := 0; x < len(cave); x++ {
		for y := 0; y < len(cave[x]); y++ {
			sum += isLowPoint(cave, x, y)
		}
	}
	return sum
}

func SolvePuzzlePart2(cave [][]int) int {
	var isPartOfBassin = make(map[int]map[int]bool)
	for x := 0; x < len(cave); x++ {
		isPartOfBassin[x] = make(map[int]bool)
	}

	xlen := len(cave)
	ylen := len(cave[0])
	var bassinSizes []int
	for x := 0; x < len(cave); x++ {
		for y := 0; y < len(cave[x]); y++ {
			if cave[x][y] == 9 || isPartOfBassin[x][y] {
				continue
			}
			bassinSizes = append(bassinSizes, growBasin(cave, &isPartOfBassin, x, y, xlen, ylen))
		}
	}

	sort.Ints(bassinSizes)
	sum := bassinSizes[len(bassinSizes)-3]
	for i := len(bassinSizes) - 2; i < len(bassinSizes); i++ {
		sum = sum * bassinSizes[i]
	}
	return sum
}

func growBasin(cave [][]int, isPartOfBassin *map[int]map[int]bool, x, y, xlen, ylen int) int {
	var sum int
	for newx := x + 1; newx < xlen; newx++ {
		growth := grow(cave, isPartOfBassin, newx, y, xlen, ylen)
		if growth == 0 {
			break
		}
		sum += growth
	}
	for newx := x - 1; newx >= 0; newx-- {
		growth := grow(cave, isPartOfBassin, newx, y, xlen, ylen)
		if growth == 0 {
			break
		}
		sum += growth
	}
	for newy := y + 1; newy < ylen; newy++ {
		growth := grow(cave, isPartOfBassin, x, newy, xlen, ylen)
		if growth == 0 {
			break
		}
		sum += growth
	}
	for newy := y - 1; newy >= 0; newy-- {
		growth := grow(cave, isPartOfBassin, x, newy, xlen, ylen)
		if growth == 0 {
			break
		}
		sum += growth
	}
	return sum
}

func grow(cave [][]int, isPartOfBassin *map[int]map[int]bool, x, y, xlen, ylen int) int {
	if (*isPartOfBassin)[x][y] || cave[x][y] == 9 {
		return 0
	}
	(*isPartOfBassin)[x][y] = true
	return growBasin(cave, isPartOfBassin, x, y, xlen, ylen) + 1
}

func isLowPoint(cave [][]int, x, y int) int {
	val := cave[x][y]
	xmax := len(cave) - 1
	ymax := len(cave[0]) - 1
	if (x-1 < 0 || val < cave[x-1][y]) &&
		(y-1 < 0 || val < cave[x][y-1]) &&
		(x+1 > xmax || val < cave[x+1][y]) &&
		(y+1 > ymax || val < cave[x][y+1]) {
		return val + 1
	}
	return 0
}
