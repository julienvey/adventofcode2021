package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	input := readInput("fixtures/input.txt")
	assert.Equal(t, 10, SolvePuzzlePart1(input))

	input2 := readInput("fixtures/input2.txt")
	assert.Equal(t, 19, SolvePuzzlePart1(input2))

	input3 := readInput("fixtures/input3.txt")
	assert.Equal(t, 226, SolvePuzzlePart1(input3))
}

func TestSolvePuzzlePart2(t *testing.T) {
	input := readInput("fixtures/input.txt")
	assert.Equal(t, 36, SolvePuzzlePart2(input))

	input2 := readInput("fixtures/input2.txt")
	assert.Equal(t, 103, SolvePuzzlePart2(input2))

	input3 := readInput("fixtures/input3.txt")
	assert.Equal(t, 3509, SolvePuzzlePart2(input3))
}
