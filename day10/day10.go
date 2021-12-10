package main

import (
	"fmt"
	"github.com/julienvey/adventofcode2021/utils"
	"sort"
)

func main() {
	input := utils.ReadInputString("day10/input.txt")
	fmt.Printf("Solution Day 10, Part 1: %d\n", SolvePuzzlePart1(input))
	fmt.Printf("Solution Day 10, Part 2: %d\n", SolvePuzzlePart2(input))
}

var match = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
var reverse = map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
var score1 = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
var score2 = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}

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
					sum += score1[c]
					goto newLine
				}
				seq = seq[:len(seq)-1]
			}
		}
	newLine:
	}
	return sum
}

func SolvePuzzlePart2(input []string) int {
	var completeScores []int
	for _, line := range input {
		var completeSum int
		seq := make([]rune, 0)

		for _, c := range line {
			switch c {
			case '(', '[', '{', '<':
				seq = append(seq, c)
			default:
				last := seq[len(seq)-1] // C
				if last != reverse[c] {
					goto newLine
				}
				seq = seq[:len(seq)-1]
			}
		}

		for i := len(seq) - 1; i >= 0; i-- {
			completeSum = completeSum*5 + score2[match[seq[i]]]
		}
		completeScores = append(completeScores, completeSum)
	newLine:
	}
	sort.Ints(completeScores)
	return completeScores[len(completeScores)/2]
}
