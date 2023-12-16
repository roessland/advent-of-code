package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrt(t *testing.T) {
	t.Run("is correct for 1 equation", func(t *testing.T) {
		{
			// x = 1 (mod 2)
			// x = 1
			// N = 2
			x, N := crt([]int64{1}, []int64{2})

			assert.EqualValues(t, x, 1)
			assert.EqualValues(t, N, 2)
		}
	})

	t.Run("is correct for 2 equations", func(t *testing.T) {
		{
			// x = 2 (mod 3)
			// x = 3 (mod 4)
			x, N := crt([]int64{2, 3}, []int64{3, 4})
			assert.EqualValues(t, 11, x)
			assert.EqualValues(t, 12, N)
			assert.EqualValues(t, 2, x%3)
			assert.EqualValues(t, 3, x%4)
		}
	})

	t.Run("is correct for 3 equations", func(t *testing.T) {
		{
			// x = 3 (mod 37)
			// x = 7 (mod 29)
			// x = 1 (mod 3)
			// x ≡ 2704 (mod 3219).
			x, N := crt([]int64{3, 7, 1}, []int64{37, 29, 3})
			assert.EqualValues(t, 2704, x)
			assert.EqualValues(t, 3219, N)
			assert.EqualValues(t, 3, x%37)
			assert.EqualValues(t, 7, x%29)
			assert.EqualValues(t, 1, x%3)
		}
	})

	t.Run("is correct for 4 equations (1 coprime, bigger)", func(t *testing.T) {
		{
			// x = 3 (mod 74)
			// x = 3 (mod 37)
			// x = 7 (mod 29)
			// x = 1 (mod 3)
			// x ≡ 2704 (mod 3219).
			x, N := crt([]int64{3, 7, 1}, []int64{37, 29, 3})
			assert.EqualValues(t, 2704, x)
			assert.EqualValues(t, 3219, N)
			assert.EqualValues(t, 3, x%37)
			assert.EqualValues(t, 7, x%29)
			assert.EqualValues(t, 1, x%3)
		}
	})
}
