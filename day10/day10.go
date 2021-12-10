package main

import (
	"fmt"
	"github.com/julienvey/adventofcode2021/utils"
)

func main() {
	input := utils.ReadInputString("day10/input.txt")
	fmt.Printf("Solution Day 10, Part 1: %d\n", SolvePuzzlePart1(input))
}

var reverse = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var score = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func SolvePuzzlePart1(input []string) int {
	var sum int
	for _, line := range input {
		seq := make([]rune, 0)
		for _, c := range line {
			switch c {
			case '(', '[', '{', '<':
				seq = append(seq, c)
			default:
				last := seq[len(seq)-1] // C
				if last != reverse[c] {
					sum += score[c]
					goto newLine
				}
				seq = seq[:len(seq)-1]
			}
		}
	newLine:
	}
	return sum
}
