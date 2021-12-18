package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	part1, part2 := SolvePuzzle("target area: x=20..30, y=-10..-5")
	assert.Equal(t, 45, part1)
	assert.Equal(t, 112, part2)
}
