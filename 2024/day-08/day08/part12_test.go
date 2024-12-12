package day08_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/roessland/advent-of-code/2024/day-08/day08"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	t0 := time.Now()
	one, two := day08.Part12("input.txt")
	assert.Equal(t, 364, one)
	assert.Equal(t, 1231, two)
	fmt.Println(time.Since(t0))

	{
		exOne, exTwo := day08.Part12("input-ex1.txt")
		assert.Equal(t, 14, exOne)
		assert.Equal(t, 34, exTwo)
	}
}
