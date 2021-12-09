package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	input := readInput("fixtures/input.txt")
	assert.Equal(t, 15, SolvePuzzlePart1(input))
}

func TestSolvePuzzlePart2(t *testing.T) {
	input := readInput("fixtures/input.txt")
	assert.Equal(t, 1134, SolvePuzzlePart2(input))
}
