package day18_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-18/day18"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	one, two := day18.Part12("input-ex1.txt", 6)
	assert.Equal(t, 22, one)
	assert.Equal(t, -1, two)
}
