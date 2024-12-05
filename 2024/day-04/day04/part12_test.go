package day04_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-04/day04"
	"github.com/stretchr/testify/assert"
)

func TestIsXmas(t *testing.T) {
	m := day04.ReadInput("input-ex1.txt")
	assert.Equal(t, 1, m.IsXmas(2, 1))
}

func TestPart12(t *testing.T) {
	exOne, exTwo := day04.Part12("input-ex1.txt")
	assert.Equal(t, 18, exOne)
	assert.Equal(t, 9, exTwo)

	one, two := day04.Part12("input.txt")
	assert.Equal(t, 2530, one)
	assert.Equal(t, 1921, two)
}
