package day12_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-12/day12"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	one, two := day12.Part12("input.txt")
	assert.Equal(t, 1450816, one)
	assert.Equal(t, 865662, two)
}
