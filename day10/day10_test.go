package main

import (
	"github.com/julienvey/adventofcode2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	input := utils.ReadInputString("fixtures/input.txt")
	assert.Equal(t, 26397, SolvePuzzlePart1(input))
}

func TestSolvePuzzlePart2(t *testing.T) {
	input := utils.ReadInputString("fixtures/input.txt")
	assert.Equal(t, 288957, SolvePuzzlePart2(input))
}
