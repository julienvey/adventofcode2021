package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	tests := []struct {
		input  []string
		result int64
	}{
		{[]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111",
			"11100", "10000", "11001", "00010", "01010"}, 198},
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
		input  []string
		result int64
	}{
		{[]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111",
			"11100", "10000", "11001", "00010", "01010"}, 230},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v-%d", tt.input, tt.result), func(t *testing.T) {
			assert.Equal(t, tt.result, SolvePuzzlePart2(tt.input))
		})
	}
}
