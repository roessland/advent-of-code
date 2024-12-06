package day05_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/roessland/advent-of-code/2024/day-05/day05"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	t0 := time.Now()
	one, two := day05.Part12("input.txt")
	assert.Equal(t, 6242, one)
	assert.Equal(t, 5169, two)
	fmt.Println(time.Since(t0))

	{
		exOne, exTwo := day05.Part12("input-ex1.txt")
		assert.Equal(t, 143, exOne)
		assert.Equal(t, 123, exTwo)
	}
}
