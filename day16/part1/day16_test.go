package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePuzzlePart1(t *testing.T) {
	assert.Equal(t, 6, SolvePuzzlePart1("D2FE28"))
	assert.Equal(t, 16, SolvePuzzlePart1("8A004A801A8002F478"))
	assert.Equal(t, 12, SolvePuzzlePart1("620080001611562C8802118E34"))
	assert.Equal(t, 23, SolvePuzzlePart1("C0015000016115A2E0802F182340"))
	assert.Equal(t, 31, SolvePuzzlePart1("A0016C880162017C3686B18A3D4780"))
}
