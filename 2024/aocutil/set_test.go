package aocutil_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2024/aocutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	t.Run("With doesn't mutate", func(t *testing.T) {
		s := aocutil.NewImmutableSet[int]()
		s.With(1)
		assert.Equal(t, s.Size(), 0)

		s2 := s.With(2).With(2, 3)
		assert.Equal(t, s2.Size(), 2)
		assert.False(t, s2.Has(1))
		assert.True(t, s2.Has(2), s2.String())
		assert.True(t, s2.Has(3))
	})

	t.Run("Intersect sets", func(t *testing.T) {
		s := aocutil.NewImmutableSet[int]()
		s = s.With(1)
		s = s.With(2)
		s2 := aocutil.NewImmutableSet[int]()
		s2 = s2.With(2)
		s2 = s2.With(3)
		s3 := s.Intersection(s2)
		require.EqualValues(t, []int{2}, s3.Values())
	})

	t.Run("Union sets", func(t *testing.T) {
		s1 := aocutil.NewImmutableSet[int](1, 2, 3)
		s2 := aocutil.NewImmutableSet[int](3, 4, 5)
		s3 := s1.Union(s2)
		assert.Equal(t, 5, s3.Size())
	})
}
