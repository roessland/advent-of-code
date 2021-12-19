package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMulMat(t *testing.T) {
	A := Mat{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}

	B := Mat{
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
	}

	actualC := A.MulMat(B)

	expectC := Mat{
		30, 24, 18,
		84, 69, 54,
		138, 114, 90,
	}

	require.Equal(t, expectC, actualC)
}

func TestMulVec(t *testing.T) {
	A := Mat{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}

	B := Vec{
		9,
		6,
		3,
	}

	actualC := A.MulVec(B)

	expectC := Vec{
		30,
		84,
		138,
	}

	require.Equal(t, expectC, actualC)
}

func TestOrientationsAreUniqueAndOrthonormal(t *testing.T) {
	require.Len(t, rotations, 24)
	identity := Identity()
	uniq := map[Mat]struct{}{}
	for _, R := range rotations {
		// Orthonormal.
		require.Equal(t, identity, R.MulMat(R.T()))
		uniq[R] = struct{}{}
	}
	require.Len(t, uniq, 24)
}
