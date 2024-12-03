package day03_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-03/day03"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	one, two := day03.Part12("input.txt")
	assert.Equal(t, 183669043, one)
	assert.Equal(t, 59097164, two)
}
