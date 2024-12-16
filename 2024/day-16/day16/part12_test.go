package day16_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-16/day16"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	one, two := day16.Part12("input.txt")
	assert.Equal(t, 75416, one)
	assert.Equal(t, 476, two)
}
