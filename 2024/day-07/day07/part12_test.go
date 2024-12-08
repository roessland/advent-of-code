package day07_test

import (
	"fmt"
	"testing"
	"time"

	day07 "github.com/roessland/advent-of-code/2024/day-07/day07"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	t0 := time.Now()
	one, two := day07.Part12("input.txt")
	assert.Equal(t, 1620690235709, one)
	assert.Equal(t, 145397611075341, two)
	fmt.Println(time.Since(t0))

	{
		exOne, exTwo := day07.Part12("input-ex1.txt")
		assert.Equal(t, 3749, exOne)
		assert.Equal(t, 11387, exTwo)
	}
}

func TestCat(t *testing.T) {
	assert.Equal(t, 100, day07.Order(101))
	assert.Equal(t, 100, day07.Order(100))
	assert.Equal(t, 10, day07.Order(99))
	assert.Equal(t, 10, day07.Order(10))
	assert.Equal(t, 1, day07.Order(1))
	assert.Equal(t, 1, day07.Order(0))
	assert.Equal(t, 123, day07.Cat(1, 23))
	assert.Equal(t, 123, day07.Cat(12, 3))
	assert.Equal(t, 120, day07.Cat(12, 0))
	assert.Equal(t, 1259435, day07.Cat(125943, 5))
	assert.Equal(t, 1259435, day07.Cat(1, 259435))
	assert.Equal(t, 1259435, day07.Cat(12, 59435))
	assert.Equal(t, 1259435, day07.Cat(125, 9435))
}
