package day14_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-14/day14"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	one, two := day14.Part12("input.txt")
	assert.Equal(t, 0, one)
	assert.Equal(t, 0, two)
}
