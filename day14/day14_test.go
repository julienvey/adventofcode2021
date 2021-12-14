package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	input := readInput("fixtures/input.txt")
	assert.Equal(t, 1588, SolvePuzzle(input, 10))
}

func TestSolvePuzzlePart2(t *testing.T) {
	input := readInput("fixtures/input.txt")
	assert.Equal(t, 2188189693529, SolvePuzzle(input, 40))
}
