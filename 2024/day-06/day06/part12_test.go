package day06_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/roessland/advent-of-code/2024/day-06/day06"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	t0 := time.Now()
	one, two := day06.Part12("input.txt")
	assert.Equal(t, 4580, one)
	assert.Equal(t, 1480, two)
	fmt.Println(time.Since(t0))

	{
		exOne, exTwo := day06.Part12("input-ex1.txt")
		assert.Equal(t, 41, exOne)
		assert.Equal(t, 6, exTwo)
	}
}
