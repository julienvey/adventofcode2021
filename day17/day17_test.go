package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	assert.Equal(t, 45, SolvePuzzle("target area: x=20..30, y=-10..-5"))
}
