package day20_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-20/day20"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	{
		one, two := day20.Part12("input.txt")
		assert.Equal(t, 1307, one)
		assert.Equal(t, 986545, two)
	}
}
