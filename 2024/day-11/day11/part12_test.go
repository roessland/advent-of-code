package day11_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-11/day11"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	one, two := day11.Part12("input.txt")
	assert.Equal(t, 193607, one)
	assert.Equal(t, 229557103025807, two)
}
