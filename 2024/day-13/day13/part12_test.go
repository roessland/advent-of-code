package day13_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-13/day13"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	// one, _ := day13.Part12("input-ex1.txt")
	// assert.Equal(t, 480, one)

	one, two := day13.Part12("input.txt")
	assert.Equal(t, 33481, one)
	assert.Equal(t, 92572057880885, two)
}
