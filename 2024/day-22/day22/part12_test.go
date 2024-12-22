package day22

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart12(t *testing.T) {
	example1 := []int{
		123,      // 3
		15887950, // 0
		16495136, // 6
		527345,   // 5
		704524,   // 4
		1553684,  // 4
		12683156, // 6
		11100544,
		12249484,
		7753432,
		5908254,
	}

	for i := 0; i < len(example1)-1; i++ {
		assert.Equal(t, example1[i+1], Next(example1[i]))
	}

	assert.Equal(t, 8685429, Nth(1, 2000))
	assert.Equal(t, 4700978, Nth(10, 2000))
	assert.Equal(t, 15273692, Nth(100, 2000))
	assert.Equal(t, 8667524, Nth(2024, 2000))

	ex1, _ := Part12("input-ex1.txt")
	assert.Equal(t, 37327623, ex1)

	_, ex2 := Part12("input-ex2.txt")
	assert.Equal(t, 23, ex2)

	p1, p2 := Part12("input.txt")
	assert.Equal(t, 20506453102, p1)
	assert.Equal(t, 0, p2)

	assert.EqualValues(t, FourConsecutiveChanges{Price: 4, Changes: [4]int8{int8(-3), int8(6), int8(-1), int8(-1)}}, Take(EachChange(123), 1)[0])
	assert.EqualValues(t, FourConsecutiveChanges{Price: 4, Changes: [4]int8{int8(6), int8(-1), int8(-1), int8(0)}}, Take(EachChange(123), 2)[1])

	numFours := 0
	for range EachChange(123) {
		numFours++
	}
	assert.Equal(t, 2000-3, numFours)

	assert.Equal(t, 23, Part2([]int{1, 2, 3, 2024}))

	// 1 2 3 4 5 changes
	// Consecutive changes: 1 2 3 4, 2 3 4 5
	// n - 3
}

func Take(seq iter.Seq[FourConsecutiveChanges], n int) []FourConsecutiveChanges {
	var res []FourConsecutiveChanges
	for i := 0; i < n; i++ {
		seq(func(changes FourConsecutiveChanges) bool {
			res = append(res, changes)
			return len(res) < n
		})
	}
	return res
}
