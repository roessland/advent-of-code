package day15_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-15/day15"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	one, two := day15.Part12("input.txt")
	assert.Equal(t, 0, one)
	assert.Equal(t, 0, two)
}
