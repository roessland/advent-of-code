package day01_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-01/day01"
	"github.com/stretchr/testify/assert"
)

func TestPart1Example1(t *testing.T) {
	assert.Equal(t, 11, day01.Part1("input-ex1.txt"))
}

func TestPart1Example2(t *testing.T) {
	assert.Equal(t, 11, day01.Part1("input-ex2.txt"))
}

func TestPart2Example1(t *testing.T) {
	assert.Equal(t, 31, day01.Part2("input-ex1.txt"))
}

func TestPart1(t *testing.T) {
	// Too low
	assert.Equal(t, 1590491, day01.Part1("input.txt"))
}

func TestPart2(t *testing.T) {
	// Too low
	assert.Equal(t, 22588371, day01.Part2("input.txt"))
}
