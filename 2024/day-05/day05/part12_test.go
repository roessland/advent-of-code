package day05_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/roessland/advent-of-code/2024/day-05/day05"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	one, two := day05.Part12("input.txt")
	assert.Equal(t, 6242, one)
	assert.Equal(t, 5169, two)

	t0 := time.Now()
	{
		exOne, exTwo := day05.Part12("input-ex1.txt")
		assert.Equal(t, 143, exOne)
		assert.Equal(t, 123, exTwo)
	}
	fmt.Println(time.Since(t0))
}
