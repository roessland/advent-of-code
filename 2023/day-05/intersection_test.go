package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractIntersection(t *testing.T) {
	t.Parallel()

	t.Run("(10, 30), (30, 40)", func(t *testing.T) {
		a := Range{10, 30}
		b := Range{30, 40}
		intersection, remaining := ExtractIntersection(a, b)
		assert.Equal(t, 0, intersection.Size())
		assert.Len(t, remaining, 1)
	})

	t.Run("(40, 50), (30, 40)", func(t *testing.T) {
		a := Range{40, 50}
		b := Range{30, 40}
		intersection, remaining := ExtractIntersection(a, b)
		assert.Equal(t, 0, intersection.Size())
		assert.Len(t, remaining, 1)
		assert.Equal(t, 40, remaining[0].Start)
		assert.Equal(t, 50, remaining[0].End)
	})

	t.Run("(10, 40), (30, 40)", func(t *testing.T) {
		a := Range{10, 40}
		b := Range{30, 40}
		intersection, remaining := ExtractIntersection(a, b)
		assert.Equal(t, 10, intersection.Size())
		assert.Equal(t, 30, intersection.Start)
		assert.Equal(t, 40, intersection.End)
		assert.Len(t, remaining, 1)
		assert.Equal(t, 10, remaining[0].Start)
		assert.Equal(t, 30, remaining[0].End)
	})

	t.Run("(10, 40), (30, 50)", func(t *testing.T) {
		a := Range{10, 40}
		b := Range{30, 50}
		intersection, remaining := ExtractIntersection(a, b)
		assert.Equal(t, 10, intersection.Size())
		assert.Equal(t, 30, intersection.Start)
		assert.Equal(t, 40, intersection.End)
		assert.Len(t, remaining, 1)
		assert.Equal(t, 10, remaining[0].Start)
		assert.Equal(t, 30, remaining[0].End)
	})

	t.Run("(10, 50), (20, 40)", func(t *testing.T) {
		a := Range{10, 50}
		b := Range{20, 40}
		intersection, remaining := ExtractIntersection(a, b)
		assert.Equal(t, 20, intersection.Size())
		assert.Equal(t, 20, intersection.Start)
		assert.Equal(t, 40, intersection.End)
		assert.Len(t, remaining, 2)
		assert.Equal(t, 10, remaining[0].Start)
		assert.Equal(t, 20, remaining[0].End)
		assert.Equal(t, 40, remaining[1].Start)
		assert.Equal(t, 50, remaining[1].End)
	})

	t.Run("(20, 40), (10, 50)", func(t *testing.T) {
		a := Range{20, 40}
		b := Range{10, 50}
		intersection, remaining := ExtractIntersection(a, b)
		assert.Equal(t, 20, intersection.Size())
		assert.Equal(t, 20, intersection.Start)
		assert.Equal(t, 40, intersection.End)
		require.Len(t, remaining, 0)
	})

	t.Run("(10, 10), (10, 10)", func(t *testing.T) {
		a := Range{10, 10}
		b := Range{10, 10}
		intersection, remaining := ExtractIntersection(a, b)
		assert.Equal(t, 0, intersection.Size())
		require.Len(t, remaining, 0)
	})
}
