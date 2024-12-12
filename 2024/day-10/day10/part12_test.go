package day10_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/roessland/advent-of-code/2024/day-10/day10"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	t0 := time.Now()
	one, two := day10.Part12("input.txt")
	assert.Equal(t, 667, one)
	assert.Equal(t, 1344, two)
	fmt.Println(time.Since(t0))

	{
		exOne, exTwo := day10.Part12("input-ex2.txt")
		assert.Equal(t, 36, exOne)
		assert.Equal(t, 81, exTwo)
	}
}
