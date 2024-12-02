package day02_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-02/day02"
	"github.com/stretchr/testify/assert"
)

func TestPart1Example1(t *testing.T) {
	assert.Equal(t, 2, day02.Part1("input-ex1.txt"))
}

func TestPart2Example1(t *testing.T) {
	assert.Equal(t, 4, day02.Part2("input-ex1.txt"))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 341, day02.Part1("input.txt"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 404, day02.Part2("input.txt"))
}
