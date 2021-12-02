package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	tests := []struct {
		input  []int
		result int
	}{
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{1, 2, 3, 4, 5, 5, 5, 4, 6}, 5},
		{[]int{5, 5}, 0},
		{[]int{5, 4, 3, 4, 5}, 2},
		{[]int{}, 0},
		{[]int{5}, 0},
		{[]int{3, 4, 5}, 2},
		{[]int{5, 4, 3}, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v-%d", tt.input, tt.result), func(t *testing.T) {
			assert.Equal(t, tt.result, SolvePuzzlePart1(tt.input))
		})
	}
}

func TestSolvePuzzlePart2(t *testing.T) {
	tests := []struct {
		input  []int
		result int
	}{
		{[]int{250, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 4},
		{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 5},
		{[]int{199, 199, 208, 210, 200, 207, 240, 269, 260, 263}, 6},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v-%d", tt.input, tt.result), func(t *testing.T) {
			assert.Equal(t, tt.result, SolvePuzzlePart2(tt.input))
		})
	}
}
