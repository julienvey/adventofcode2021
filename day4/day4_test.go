package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	input := readInput("fixtures/input.txt")
	assert.Equal(t, 4512, SolvePuzzlePart1(input))
}

func TestSolvePuzzlePart2(t *testing.T) {
	input := readInput("fixtures/input.txt")
	assert.Equal(t, 1924, SolvePuzzlePart2(input))
}

func TestCheckWinner(t *testing.T) {
	tests := []struct {
		boards [][][]BoardNumber
		result bool
	}{
		{[][][]BoardNumber{
			{
				[]BoardNumber{{Marked: true}, {Marked: true}, {Marked: true}, {Marked: true}, {Marked: true}},
				[]BoardNumber{{Marked: true}, {Marked: true}, {Marked: true}, {Marked: true}, {Marked: true}},
				[]BoardNumber{{Marked: true}, {Marked: true}, {Marked: true}, {Marked: true}, {Marked: true}},
				[]BoardNumber{{Marked: true}, {Marked: true}, {Marked: true}, {Marked: true}, {Marked: true}},
				[]BoardNumber{{Marked: true}, {Marked: true}, {Marked: true}, {Marked: true}, {Marked: true}},
			},
		}, true},
		{[][][]BoardNumber{
			{
				[]BoardNumber{{Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
			},
		}, false},
		{[][][]BoardNumber{
			{
				[]BoardNumber{{Marked: true}, {Marked: true}, {Marked: false}, {Marked: true}, {Marked: true}},
				[]BoardNumber{{Marked: false}, {Marked: false}, {Marked: true}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: true}, {Marked: false}, {Marked: true}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: false}, {Marked: false}, {Marked: true}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: true}, {Marked: false}, {Marked: true}, {Marked: false}, {Marked: false}},
			},
		}, false},
		{[][][]BoardNumber{
			{
				[]BoardNumber{{Marked: true}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: true}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: true}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: true}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: true}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
			},
		}, true},
		{[][][]BoardNumber{
			{
				[]BoardNumber{{Marked: true}, {Marked: true}, {Marked: true}, {Marked: true}, {Marked: true}},
				[]BoardNumber{{Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
				[]BoardNumber{{Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}, {Marked: false}},
			},
		}, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v-%t", tt.boards, tt.result), func(t *testing.T) {
			winner, _ := checkWinner(tt.boards)
			if tt.result {
				assert.NotNil(t, winner)
			} else {
				assert.Nil(t, winner)
			}
		})
	}
}

func TestGetScore(t *testing.T) {
	tests := []struct {
		board  [][]BoardNumber
		draw   int
		result int
	}{
		{[][]BoardNumber{
			{BoardNumber{Number: 14, Marked: true}, BoardNumber{Number: 21, Marked: true}, BoardNumber{Number: 17, Marked: true}, BoardNumber{Number: 24, Marked: true}, BoardNumber{Number: 4, Marked: true}},
			{BoardNumber{Number: 10}, BoardNumber{Number: 16}, BoardNumber{Number: 15}, BoardNumber{Number: 9, Marked: true}, BoardNumber{Number: 19}},
			{BoardNumber{Number: 18}, BoardNumber{Number: 8}, BoardNumber{Number: 23, Marked: true}, BoardNumber{Number: 26}, BoardNumber{Number: 20}},
			{BoardNumber{Number: 22}, BoardNumber{Number: 11, Marked: true}, BoardNumber{Number: 13}, BoardNumber{Number: 6}, BoardNumber{Number: 5, Marked: true}},
			{BoardNumber{Number: 2, Marked: true}, BoardNumber{Number: 0, Marked: true}, BoardNumber{Number: 12}, BoardNumber{Number: 3}, BoardNumber{Number: 7, Marked: true}},
		}, 24, 4512},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v-%d", tt.board, tt.result), func(t *testing.T) {
			assert.Equal(t, tt.result, getScore(tt.board, tt.draw))
		})
	}
}
