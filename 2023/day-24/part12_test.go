package main

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	zero = big.NewRat(0, 1)
	one  = big.NewRat(1, 1)
	inf  = new(big.Rat).Mul(big.NewRat(math.MaxInt64, 1), big.NewRat(math.MaxInt64, 1))
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
			1, 0, 0, -1, 1, 0,
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
		{
			// Both move forwards in time. One in +x and the other in -x.
			"in time, collide in future",
			0, 0, 0, 1, 1, 0,
			1, 0, 0, -1, 1, 0,
			0.5, 0.5,
			true,
		},
		{
			// Both move forwards in time. One in +x and the other in -x.
			"in time, collide in past",
			0, -1, 0, 1, 1, 0,
			1, -1, 0, -1, 1, 0,
			0.5, -0.5,
			true,
		},
		{
			// Both move forwards in time. One in +x and the other in -x.
			"in time, same x-directions, collide in future",
			-10, 0, 0, 10, 1, 0, // starts way left, moves fast
			0, 0, 0, 1, 1, 0, // starts at 0, moves slow
			1.111111, 1.11111,
			true,
		},
		{
			// Both move forwards in time. One in +x and the other in -x.
			"in time, same x-directions, collide in future, flipped",
			10, 0, 0, -10, 1, 0, // starts way left, moves fast
			0, 0, 0, -1, 1, 0, // starts at 0, moves slow
			-1.111111, 1.11111,
			true,
		},
		{
			// Both move forwards in time. One in +x and the other in -x.
			"in time, one standing still",
			-1, 0, 0, 10, 1, 0, // starts left, moves fast
			0, 0, 0, 0, 1, 0, // starts at 0, stays there
			0.0, 0.1,
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

			xA := h1.Pos.X
			yA := h1.Pos.Y
			xAv := h1.Vel.X
			yAv := h1.Vel.Y
			xB := h2.Pos.X
			yB := h2.Pos.Y
			xBv := h2.Vel.X
			yBv := h2.Vel.Y
			pRat, intersects := LineLineIntersection2D(xA, yA, xAv, yAv, xB, yB, xBv, yBv)
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
		A := Mat2{
			big.NewRat(0, 1), big.NewRat(-1, 1),
			big.NewRat(-1, 1), big.NewRat(-1, 1),
		}
		b := BigRatVec2{
			big.NewRat(int64(xA*yAv-yA*xAv), 1),
			big.NewRat(int64(xB*yBv-yB*xBv), 1),
		}

		_, ok := A.Solve(b)
		require.True(t, ok)
	}
}

func TestMean(t *testing.T) {
	h1 := Hailstone{
		Pos: NewBigRatVec3(1, 2, 3),
		Vel: NewBigRatVec3(-1, -2, -3),
	}
	h2 := Hailstone{
		Pos: NewBigRatVec3(4, 5, 6),
		Vel: NewBigRatVec3(-4, -5, -6),
	}
	hMean := Mean([]Hailstone{h1, h2})

	fMean := hMean.ToFloat()

	require.InDelta(t, (1+4)/2.0, fMean.Pos.X, eps)
	require.InDelta(t, (2+5)/2.0, fMean.Pos.Y, eps)
	require.InDelta(t, (3+6)/2.0, fMean.Pos.Z, eps)
	require.InDelta(t, (-1-4)/2.0, fMean.Vel.X, eps)
	require.InDelta(t, (-2-5)/2.0, fMean.Vel.Y, eps)
	require.InDelta(t, (-3-6)/2.0, fMean.Vel.Z, eps)
}

func TestSSE(t *testing.T) {
	hs := ReadInput("input.txt")
	n := len(hs)

	px := NewBigRatVecN(n)
	py := NewBigRatVecN(n)
	pz := NewBigRatVecN(n)

	vx := NewBigRatVecN(n)
	vy := NewBigRatVecN(n)
	vz := NewBigRatVecN(n)

	for i := range hs {
		px[i] = hs[i].Pos.X
		py[i] = hs[i].Pos.Y
		pz[i] = hs[i].Pos.Z
		vx[i] = hs[i].Vel.X
		vy[i] = hs[i].Vel.Y
		vy[i] = hs[i].Vel.Z
	}

	fmt.Println("px", px)
	fmt.Println("vx", vx)
	//	 287838354624648//1
	//	 412952398656862//1
	//	 148587016955395//1
	//	              -5//1
	//	            -250//1
	//	             217//1
	px0 := big.NewRat(287838354624648, 1)
	py0 := big.NewRat(412952398656862, 1)
	pz0 := big.NewRat(148587016955395, 1)
	vx0 := big.NewRat(-5, 1)
	vy0 := big.NewRat(-250, 1)
	vz0 := big.NewRat(217, 1)
	//
	// px0 := big.NewRat(287838354624631, 1)
	// py0 := big.NewRat(412952398656777, 1)
	// pz0 := big.NewRat(148587016955475, 1)
	// vx0 := big.NewRat(-5, 1)
	// vy0 := big.NewRat(-250, 1)
	// vz0 := big.NewRat(217, 1)

	fmt.Println("SumSquaredError X", SumSquaredError(px, vx, px0, vx0))

	fmt.Println("SumSquaredError Y", SumSquaredError(py, vy, py0, vy0))

	fmt.Println("SumSquaredError Z", SumSquaredError(pz, vz, pz0, vz0))

	sum := big.NewRat(0, 1)
	sum = sum.Add(sum, px0)
	sum = sum.Add(sum, py0)
	sum = sum.Add(sum, pz0)
	fmt.Println("X")
	fmt.Println("Sum for part 2 is:", sum)
	// 849377770236883 too low
	// 849377770236905 ? Yes!
}

func SumSquaredError(p, v []*big.Rat, p0, v0 *big.Rat) *big.Rat {
	sum := big.NewRat(0, 1)
	for i := range p {
		fmt.Println("checking p", p[i], "v", v[i])
		fmt.Println("versus p0", p0, "v0", v0)
		sum = sum.Add(sum, SquaredError(p[i], v[i], p0, v0))
	}
	return sum
}

func SquaredError(pi, vi, p0, v0 *big.Rat) *big.Rat {
	e := Error(pi, vi, p0, v0)
	return e.Mul(e, e)
}

func Error(pi, vi, p0, v0 *big.Rat) *big.Rat {
	// Same line when projected to the X, T plane
	if pi.Cmp(p0) == 0 && vi.Cmp(v0) == 0 {
		return new(big.Rat).Set(zero)
	}

	// First compute tPrime = [t1, t2, ... tn]
	//
	// Relabel the axes to fit into LineLineIntersection2D.
	// X,Y,Z becomes X. T becomes Y.
	//
	//   ┌────────────────────────┐
	//   │ t                      │
	//   │ ▲          ▲           │
	//   │ │        ┌─┴┐          │
	//   │ └───▶x ┌─┘  └─┐        │
	//   │      ┌─┘      └v0      │
	//   │    vi┘          └─┐    │
	//   │  ┌─┘              └─┐  │
	//  .┴.─┘                  └─.┴.
	// (pi )────────────────────(p0 )
	//  `─'                      `─'
	//
	xA := p0
	yA := new(big.Rat).Set(zero)
	xAv := v0
	yAv := new(big.Rat).Set(one)
	xB := pi
	yB := new(big.Rat).Set(zero)
	xBv := vi
	yBv := new(big.Rat).Set(one)

	intersection, intersects := LineLineIntersection2D(xA, yA, xAv, yAv, xB, yB, xBv, yBv)
	if !intersects {
		fmt.Println("no intersection", xA, yA, xAv, yAv, xB, yB, xBv, yBv)
		return new(big.Rat).Set(inf)
	}

	// Time and location of rock when it intersects with hailstone path
	xPrime, tPrimeX := intersection.X, intersection.Y
	fmt.Println("tPrime", tPrimeX)

	// Location of hailstone at t=tPrime
	xRock := new(big.Rat).Add(p0, new(big.Rat).Mul(v0, tPrimeX))
	xHailstone := new(big.Rat).Add(pi, new(big.Rat).Mul(vi, tPrimeX))

	if tPrimeX.Cmp(zero) < 0 {
		fmt.Println("PAST PAST PAST")
		fmt.Println("xPrime: ", xPrime)
		fmt.Println("tPrime", tPrimeX)
		fmt.Println("pi", pi, vi, p0, v0)
		fmt.Println("intersection was in the past at time", tPrimeX)
		fmt.Println("At that time, hailstone was at", xHailstone)
		fmt.Println("rock was at ", xRock)
	} else {
		fmt.Println("intersection was in the future at time", tPrimeX)
		fmt.Println("At that time, hailstone was at", xHailstone)
		fmt.Println("rock was at ", xRock)
	}

	// Signed distance from hailstone to rock
	return new(big.Rat).Sub(xPrime, xHailstone)
}
