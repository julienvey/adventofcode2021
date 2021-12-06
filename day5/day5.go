package main

import (
	"fmt"
	"github.com/julienvey/adventofcode2021/utils"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readInput("day5/input.txt")
	fmt.Printf("Solution Day 5, Part 1: %d\n", SolvePuzzlePart1(input))
	fmt.Printf("Solution Day 5, Part 2: %d\n", SolvePuzzlePart2(input))
}

type Entry struct {
	x1, y1, x2, y2 int
}

func readInput(in string) []Entry {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}
	var entries []Entry
	regex := regexp.MustCompile(`^([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)$`)
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		res := regex.FindAllStringSubmatch(line, -1)
		x1, err := strconv.Atoi(res[0][1])
		if err != nil {
			log.Fatal(err)
		}
		y1, err := strconv.Atoi(res[0][2])
		if err != nil {
			log.Fatal(err)
		}
		x2, err := strconv.Atoi(res[0][3])
		if err != nil {
			log.Fatal(err)
		}
		y2, err := strconv.Atoi(res[0][4])
		if err != nil {
			log.Fatal(err)
		}
		entries = append(entries, Entry{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		})
	}
	return entries
}

func SolvePuzzlePart1(entries []Entry) int {
	puzzlemap := initPuzzleMap(entries)
	fillPuzzleMap(entries, &puzzlemap, true)
	return countOverlap(puzzlemap)
}

func SolvePuzzlePart2(entries []Entry) int {
	puzzlemap := initPuzzleMap(entries)
	fillPuzzleMap(entries, &puzzlemap, false)
	return countOverlap(puzzlemap)
}

func countOverlap(puzzlemap [][]int) int {
	var sum int
	for x, _ := range puzzlemap {
		for y, _ := range puzzlemap[x] {
			if puzzlemap[x][y] > 1 {
				sum++
			}
		}
	}
	return sum
}

func initPuzzleMap(entries []Entry) [][]int {
	var xmax, ymax int
	for _, entry := range entries {
		if entry.x1 > xmax {
			xmax = entry.x1
		}
		if entry.x2 > xmax {
			xmax = entry.x2
		}
		if entry.y1 > ymax {
			ymax = entry.y1
		}
		if entry.y2 > ymax {
			ymax = entry.y2
		}
	}
	puzzlemap := make([][]int, xmax+1)
	for i := range puzzlemap {
		puzzlemap[i] = make([]int, ymax+1)
	}
	return puzzlemap
}

func fillPuzzleMap(entries []Entry, puzzlemap *[][]int, ignoreDiagonal bool) {
	for _, entry := range entries {
		if entry.x1 == entry.x2 || entry.y1 == entry.y2 {
			for x := utils.Min(entry.x1, entry.x2); x <= utils.Max(entry.x1, entry.x2); x++ {
				for y := utils.Min(entry.y1, entry.y2); y <= utils.Max(entry.y1, entry.y2); y++ {
					(*puzzlemap)[x][y]++
				}
			}
		} else if !ignoreDiagonal {
			ymin := utils.Min(entry.y1, entry.y2)
			xmin := utils.Min(entry.x1, entry.x2)
			xmax := utils.Max(entry.x1, entry.x2)
			if xmin == entry.x1 { // croissant sur x
				if ymin == entry.y1 { // croissant sur y
					for i := 0; i <= xmax-xmin; i++ {
						(*puzzlemap)[entry.x1+i][entry.y1+i]++
					}
				} else { // decroissant sur y
					for i := 0; i <= xmax-xmin; i++ {
						(*puzzlemap)[entry.x1+i][entry.y1-i]++
					}
				}
			} else { // decroissant sur x
				if ymin == entry.y1 { // croissant sur y
					for i := 0; i <= xmax-xmin; i++ {
						(*puzzlemap)[entry.x1-i][entry.y1+i]++
					}
				} else { // decroissant sur y
					for i := 0; i <= xmax-xmin; i++ {
						(*puzzlemap)[entry.x1-i][entry.y1-i]++
					}
				}
			}
		}
	}
}
