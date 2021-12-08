package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	input := readInput("fixtures/input.txt")
	assert.Equal(t, 26, SolvePuzzlePart1(input))
}

func TestSolvePuzzlePart2(t *testing.T) {
	input := readInput("fixtures/input_single.txt")
	assert.Equal(t, 5353, SolvePuzzlePart2(input))

	input = readInput("fixtures/input.txt")
	assert.Equal(t, 61229, SolvePuzzlePart2(input))
}
