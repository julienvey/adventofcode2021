package main

import (
	"github.com/julienvey/adventofcode2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	input := utils.ReadInput("fixtures/input.txt")
	assert.Equal(t, 5934, getGrowth(input, 80))
}

func TestSolvePuzzlePart2(t *testing.T) {
	input := utils.ReadInput("fixtures/input.txt")
	assert.Equal(t, 26984457539, getGrowth(input, 256))
}
