package day05

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadInput(t *testing.T) {
	{
		exRanges, exIDs := ReadInput("input-ex.txt")
		require.EqualValues(t, len(exRanges), 4)
		require.EqualValues(t, len(exIDs), 6)
		require.EqualValues(t, 3, Part1(exRanges, exIDs))
		require.EqualValues(t, 14, Part2(exRanges))
	}
	{
		ranges, ids := ReadInput("input.txt")
		require.EqualValues(t, 505, Part1(ranges, ids))
		require.EqualValues(t, 344423158480189, Part2(ranges))
	}

}

func TestRange(t *testing.T) {
	require.Equal(t, []Range{{3, 5}}, Range{3, 5}.Sub(Range{10, 14}))         // disjunct
	require.EqualValues(t, []Range(nil), Range{0, 10}.Sub(Range{-99, 99}))    // contained
	require.Equal(t, []Range{{0, 6}, {8, 10}}, Range{0, 10}.Sub(Range{7, 7})) // contains
	require.Equal(t, []Range{{0, 6}}, Range{0, 10}.Sub(Range{7, 999}))        // 1
	require.Equal(t, []Range{{8, 10}}, Range{0, 10}.Sub(Range{-99, 7}))       // 2
}
