package main

import (
	"github.com/julienvey/adventofcode2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	input := utils.ReadInput("fixtures/input.txt")
	assert.Equal(t, 37, SolvePuzzlePart1(input))
}

func TestCalcMedian(t *testing.T) {
	assert.Equal(t, 2, CalcMedian(16, 1, 2, 0, 4, 2, 7, 1, 2, 14))
	assert.Equal(t, 3, CalcMedian(16, 1, 2, 0, 4, 2, 7, 1, 2, 14, 18, 3, 4, 45, 21, 1, 0, 3, 4, 5))
	assert.Equal(t, 6, CalcMedian(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11))
}
