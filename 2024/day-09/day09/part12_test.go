package day09_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/roessland/advent-of-code/2024/day-09/day09"
	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	t0 := time.Now()
	one, two := day09.Part12("input.txt")
	assert.Equal(t, 6337367222422, one)
	assert.Equal(t, 6361380647183, two)
	fmt.Println(time.Since(t0))
}
