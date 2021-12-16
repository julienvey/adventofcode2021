package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart2(t *testing.T) {
	assert.Equal(t, 3, SolvePuzzlePart2("C200B40A82"))
	assert.Equal(t, 54, SolvePuzzlePart2("04005AC33890"))
	assert.Equal(t, 7, SolvePuzzlePart2("880086C3E88112"))
	assert.Equal(t, 1, SolvePuzzlePart2("D8005AC2A8F0"))
	assert.Equal(t, 0, SolvePuzzlePart2("F600BC2D8F"))
	assert.Equal(t, 0, SolvePuzzlePart2("9C005AC2F8F0"))
	assert.Equal(t, 1, SolvePuzzlePart2("9C0141080250320F1802104A08"))
}
