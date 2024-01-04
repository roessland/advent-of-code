package main

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const eps = 0.001

func TestLineLineIntersection2D(t *testing.T) {
	type testCase struct {
		name                            string
		h1x, h1y, h1z, h1vx, h1vy, h1vz int
		h2x, h2y, h2z, h2vx, h2vy, h2vz int
		qx, qy                          float64
		intersects                      bool
	}

	testCases := []testCase{
		{
			"test1",
			// Solve equation {{2, 1, 0}; {1, 5, 0}; {0, 0, 7}} x = {{44}; {91}; {0}} for x
			// No solutions
			19, 13, 0, -2, 1, 0,
			18, 19, 0, -1, -1, 0,
			14.333, 15.333,
			true,
		},
		{
			"test2",
			19, 13, 0, -2, 1, 0,
			20, 25, 0, -2, -2, 0,
			11.667, 16.667,
			true,
		},
		{
			"test3",
			19, 13, 0, -2, 1, 0,
			12, 31, 0, -1, -2, 0,
			6.2, 19.4,
			true,
		},
		{
			"north arrow",
			// North arrow
			0, 0, 0, 1, 1, 0,
			1, 1, 0, -1, 1, 0,
			0.5, 0.5,
			true,
		},
		{
			"parallel lines",
			// Parallel lines
			0, 0, 0, -1, 1, 0,
			1, 1, 0, -1, 1, 0,
			0.5, 0.5,
			false,
		},
		{
			"east arrow",
			// East arrow
			0, 1, 0, 1, -1, 0,
			0, 0, 0, 1, 1, 0,
			0.5, 0.5,
			true,
		},
		{
			"east arrow (flipped)",
			// East arrow
			0, 0, 0, 1, 1, 0,
			0, 1, 0, 1, -1, 0,
			0.5, 0.5,
			true,
		},
		{
			"south arrow",
			0, 1, 0, 1, -1, 0,
			1, 1, 0, -1, -1, 0,
			0.5, 0.5,
			true,
		},
		{
			"south arrow (flipped)",
			1, 1, 0, -1, -1, 0,
			0, 1, 0, 1, -1, 0,
			0.5, 0.5,
			true,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%s %d,%d,%d -- %d,%d,%d", tc.name, tc.h1x, tc.h1y, tc.h1z, tc.h2x, tc.h2y, tc.h2z)
		t.Run(name, func(t *testing.T) {
			h1 := Hailstone{
				Pos: NewBigRatVec3(tc.h1x, tc.h1y, tc.h1z),
				Vel: NewBigRatVec3(tc.h1vx, tc.h1vy, tc.h1vz),
			}

			h2 := Hailstone{
				Pos: NewBigRatVec3(tc.h2x, tc.h2y, tc.h2z),
				Vel: NewBigRatVec3(tc.h2vx, tc.h2vy, tc.h2vz),
			}

			pRat, intersects := LineLineIntersection2D(h1, h2)
			if tc.intersects {
				require.True(t, intersects)
				qx, _ := pRat.X.Float64()
				qy, _ := pRat.Y.Float64()
				assert.InDeltaf(t, tc.qx, qx, eps, "x was %f", qx)
				assert.InDeltaf(t, tc.qy, qy, eps, "y was %f", qy)
			}
		})
	}
}

func TestBigRatVec3(t *testing.T) {
	t.Run("sub", func(t *testing.T) {
		u := NewBigRatVec3(1, 2, 3)
		v := NewBigRatVec3(4, 5, 6)
		uSubV := u.Sub(v).ToFloatVec3()
		require.InDelta(t, -3.0, uSubV.X, eps)
		require.InDelta(t, -3.0, uSubV.Y, eps)
		require.InDelta(t, -3.0, uSubV.Z, eps)
	})
}

func TestInsideArea(t *testing.T) {
	areaMin := big.NewRat(7, 1)
	areaMax := big.NewRat(17, 1)

	require.True(t, InsideArea3D(NewBigRatVec3(7, 7, 7), areaMin, areaMax))
	require.True(t, InsideArea3D(NewBigRatVec3(17, 17, 17), areaMin, areaMax))
}

func TestLineLineIntersectionEquations(t *testing.T) {
	xA := 0.0
	yA := 1.0
	xAv := 1.0
	yAv := -1.0
	xB := 0.0
	yB := 0.0
	xBv := 1.0
	yBv := 1.0

	x := 0.5
	y := 0.5

	parametricA := func(s float64) (float64, float64) {
		return xA + s*xAv, yA + s*yAv
	}

	parametricB := func(t float64) (float64, float64) {
		return xB + t*xBv, yB + t*yBv
	}

	Ax, Ay := parametricA(0.5)
	Bx, By := parametricB(0.5)

	require.InDelta(t, Ax, Bx, eps)
	require.InDelta(t, Ay, By, eps)

	// Test symmetric form
	{
		// A
		{
			LHS := (x - xA) * yAv
			RHS := (y - yA) * xAv
			require.InDelta(t, LHS, RHS, eps)
		}
		// B
		{
			LHS := (x - xB) * yBv
			RHS := (y - yB) * xBv
			require.InDelta(t, LHS, RHS, eps)
		}
	}

	// test matrix form
	{
		I := (x-xA)*yAv - (y-yA)*xAv
		II := (x-xB)*yBv - (y-yB)*xBv
		require.InDelta(t, 0, I, eps)
		require.InDelta(t, 0, II, eps)
	}

	// test next matrix form
	{
	}
}
