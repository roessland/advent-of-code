package day19_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/day-19/day19"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	{
		one, two := day19.Part12("input-ex1.txt")
		assert.Equal(t, 6, one)
		assert.Equal(t, 16, two)
	}
	{
		one, two := day19.Part12("input.txt")
		assert.Equal(t, 347, one)
		assert.Equal(t, 347, two)
	}
}

func TestWays(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		patterns := []string{"a", "b"}
		assert.Equal(t, 1, day19.Ways(patterns, "a", 0))
		assert.Equal(t, 1, day19.Ways(patterns, "b", 0))
		assert.Equal(t, 1, day19.Ways(patterns, "ab", 0))
		assert.Equal(t, 1, day19.Ways(patterns, "ba", 0))
		assert.Equal(t, 1, day19.Ways(patterns, "aa", 0))
		assert.Equal(t, 1, day19.Ways(patterns, "bb", 0))
		assert.Equal(t, 1, day19.Ways(patterns, "abab", 0))
	})

	t.Run("complex", func(t *testing.T) {
		patterns := []string{"r", "b", "g", "br", "gb", "rb"}
		assert.Equal(t, 6, day19.Ways(patterns, "rrbgbr", 0))
	})
}
