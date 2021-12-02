package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	tests := []struct {
		input  []command
		result int
	}{
		{[]command{
			{"forward", 5},
			{"down", 5},
			{"forward", 8},
			{"up", 3},
			{"down", 8},
			{"forward", 2},
		}, 150},
		{[]command{
			{"forward", 5},
			{"down", 7},
			{"forward", 5},
			{"up", 3},
			{"down", 8},
			{"forward", 2},
		}, 144},
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
		input  []command
		result int
	}{
		{[]command{
			{"forward", 5},
			{"down", 5},
			{"forward", 8},
			{"up", 3},
			{"down", 8},
			{"forward", 2},
		}, 900},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v-%d", tt.input, tt.result), func(t *testing.T) {
			assert.Equal(t, tt.result, SolvePuzzlePart2(tt.input))
		})
	}
}
