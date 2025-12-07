package aocutil_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2025/aocutil"
	"github.com/stretchr/testify/assert"
)

func TestMultiset(t *testing.T) {
	t.Run("Can add same stuff twice", func(t *testing.T) {
		s := aocutil.NewImmutableMultiSet[int]()
		s = s.With(1)
		s = s.With(1)
		assert.Equal(t, s.Cardinality(), 2)
	})
}
