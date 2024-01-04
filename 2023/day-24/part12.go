package main

import (
	"fmt"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

func main() {
	hs := ReadInput()
	fmt.Println(hs)
}

type Hailstone struct {
	Pos IntVec3
	Vel IntVec3
}

type IntVec3 struct {
	X, Y, Z int
}

// FracVec3 represents [x, y, z]/d
type FracVec3 struct {
	X, Y, Z, D int
}

// >                        ▲
// >                       ╱
// >             y=7      ╱y=27
// >              ┌────────┐
// >  .─.         │     ╱  │
// > ( A )────────│────╳───│───────────▶
// >  `─'         │   ╱    │  y
// >           x=7└────────┘  ▲
// >              .─╱   x=27  │
// >             ( B )        └───▶x
// >              `─'
//
// pA = (xA, yA, zA)
// pB = (xB, yB, zB)
// vA = (xA', yA', zA')
// vB = (xB', yB', zB')
func Part1() {
	hs := ReadInput()
	for i := range hs {
		hs[i].Pos.Z = 0
		hs[i].Vel.Z = 0
	}

	areaMin := 7
	areaMax := 27
	count := 0

	for _, hA := range hs {
		for _, hB := range hs {
			p, intersects := SegmentSegmentIntersection(hA, hB)
			if intersects && p.X >= areaMin && p.X <= areaMax && p.Y >= areaMin && p.Y <= areaMax {
				count++
			}
		}
	}
	fmt.Println("Part 1:", count)
}

func SegmentSegmentIntersection(A, B Hailstone) (p FracVec3, intersects bool) {
	return LineLineIntersection(A, B)
}

func LineLineIntersection(hA, hB Hailstone) (p FracVec3, intersects bool) {
	// Parametric form:
	// ==> pA + tA * vA = pB + tB * vB,      tA, tB ∈ ℝ, > 0
	//
	// Convert to symmmetric form:
	//  x = xA + s * xA'  =>  s = (x - xA) / xA'
	//  y = yA + s * yA'  =>  s = (y - yA) / yA'
	//  z = zA + s * zA'  =>  s = (z - zA) / zA'
	// A: ==> (x - xA) / xA' = (y - yA) / yA' = (z - zA) / zA'
	// B: ==> (x - xB) / xB' = (y - yB) / yB' = (z - zB) / zB'
	//
	// This is a system of 6 equations with 3 unknowns:
	//
	// (x - xA) / xA' - (y - yA) / yA'                   = 0
	// (x - xA) / xA'                  - (z - zA) / zA'  = 0
	//                  (y - yA) / yA' - (z - zA) / zA'  = 0
	// (x - xB) / xB' - (y - yB) / yB'                   = 0
	// (x - xB) / xB'                  - (z - zB) / zB'  = 0
	//                  (y - yB) / yB' - (z - zB) / zB'  = 0
	//
	// Massage it into matrix form:
	//
	// (x - xA) / xA' - (y - yA) / yA'                   = 0  | * xA' * yA'
	// (x - xA) / xA'                  - (z - zA) / zA'  = 0  | * xA' * zA'
	//                  (y - yA) / yA' - (z - zA) / zA'  = 0  | * yA' * zA'
	// (x - xB) / xB' - (y - yB) / yB'                   = 0  | * xB' * yB'
	// (x - xB) / xB'                  - (z - zB) / zB'  = 0  | * xB' * zB'
	//                  (y - yB) / yB' - (z - zB) / zB'  = 0  | * yB' * zB'
	//
	// Step 1: Multiply to remove fractions:
	//
	// (x - xA) * yA' - (y - yA) * xA'                   = 0
	// (x - xA) * zA'                  - (z - zA) * xA'  = 0
	//                  (y - yA) * zA' - (z - zA) * yA'  = 0
	// (x - xB) * yB' - (y - yB) * xB'                   = 0
	// (x - xB) * zB'                  - (z - zB) * xB'  = 0
	//                  (y - yB) * zB' - (z - zB) * yB'  = 0
	//
	// Step 2: Expand, prepare for moving constants to the right hand side:
	//
	// x yA' - xA yA' - y xA' + yA xA' = 0
	// x zA' - xA zA' - z xA' + zA xA' = 0
	// y zA' - yA zA' - z yA' + zA yA' = 0
	// x yB' - xB yB' - y xB' + yB xB' = 0
	// x zB' - xB zB' - z xB' + zB xB' = 0
	// y zB' - yB zB' - z yB' + zB yB' = 0
	//
	// Step 3: Move constants to the right hand side:
	//
	// x yA' - y xA'         = xA yA' - yA xA'
	// x zA'         - z xA' = xA zA' - zA xA'
	//         y zA' - z yA' = yA zA' - zA yA'
	// x yB' - y xB'         = xB yB' - yB xB'
	// x zB'         - z xB' = xB zB' - zB xB'
	//         y zB' - z yB' = yB zB' - zB yB'
	//
	// Matrix form:
	//
	// [ yA'  -xA'    0 ]               [ xA yA' - yA xA' ]
	// [ zA'    0   -xA']               [ xA zA' - zA xA' ]
	// [ 0     zA'  -yA']     [ x ]     [ yA zA' - zA yA' ]
	// [ yB'  -xB'    0 ]  @  [ y ]  =  [ xB yB' - yB xB' ]
	// [ zB'    0   -xB']     [ z ]     [ xB zB' - zB xB' ]
	// [ 0     zB'  -yB']               [ yB zB' - zB yB' ]
	//
	// Example:
	// > Hailstone A: 19, 13, 30 @ -2, 1, -2
	// > Hailstone B: 18, 19, 22 @ -1, -1, -2
	// > Hailstones' paths will cross inside the test area (at x=14.333, y=15.333)
	// [ 1   2  ]               [ 19 * 1 + 13 * 2 ]
	// [ -1  1  ]               [ -18 * 1 + 19 * 1]
	// 14.333 + 2*15.333 = 44.999 = 45
	// -14.333 + 15.333 = 1 = 1
	//
	//
	//
	//
	// Solve A'A w = A'b for w
	x1 := hA.Pos.X
	y1 := hA.Pos.Y

	x2 := hA.Pos.X + hA.Vel.X
	y2 := hA.Pos.Y + hA.Vel.Y

	x3 := hB.Pos.X
	y3 := hB.Pos.Y

	x4 := hB.Pos.X + hB.Vel.X
	y4 := hB.Pos.Y + hB.Vel.Y

	PDenom := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if PDenom == 0 {
		return FracVec3{}, false
	}
	Px := (x1*y2-y1*x2)*(x3-x4) - (x1-x2)*(x3*y4-y3*x4)
	Py := (x1*y2-y1*x2)*(y3-y4) - (y1-y2)*(x3*y4-y3*x4)

	return FracVec3{Px, Py, 0, PDenom}, true
}

type Mat2 struct {
	a11, a12 int
	a21, a22 int
}

type Mat3 struct {
	a11, a12, a13 int
	a21, a22, a23 int
	a31, a32, a33 int
}

func Det(m Mat3) int {
	a11, a12, a13, a21, a22, a23, a31, a32, a33 := m.a11, m.a12, m.a13, m.a21, m.a22, m.a23, m.a31, m.a32, m.a33
	return a11*(a22*a33-a23*a32) - a12*(a21*a33-a23*a31) + a13*(a21*a32-a22*a31)
}

func Solve(A Mat3, b IntVec3) *FracVec3 {
	D := Det(A)
	if D == 0 {
		return nil
	}
	Ax := A
	Ax.a11 = b.X
	Ax.a21 = b.Y
	Ax.a31 = b.Z
	Dx := Det(Ax)

	Ay := A
	Ay.a12 = b.X
	Ay.a22 = b.Y
	Ay.a32 = b.Z
	Dy := Det(Ay)

	Az := A
	Az.a13 = b.X
	Az.a23 = b.Y
	Az.a33 = b.Z
	Dz := Det(Az)

	return &FracVec3{Dx, Dy, Dz, D}
}

func T(A Mat3) Mat3 {
	return Mat3{A.a11, A.a21, A.a31, A.a12, A.a22, A.a32, A.a13, A.a23, A.a33}
}

func MatMatMul(A, B Mat3) Mat3 {
	return Mat3{
		A.a11*B.a11 + A.a12*B.a21 + A.a13*B.a31,
		A.a11*B.a12 + A.a12*B.a22 + A.a13*B.a32,
		A.a11*B.a13 + A.a12*B.a23 + A.a13*B.a33,
		A.a21*B.a11 + A.a22*B.a21 + A.a23*B.a31,
		A.a21*B.a12 + A.a22*B.a22 + A.a23*B.a32,
		A.a21*B.a13 + A.a22*B.a23 + A.a23*B.a33,
		A.a31*B.a11 + A.a32*B.a21 + A.a33*B.a31,
		A.a31*B.a12 + A.a32*B.a22 + A.a33*B.a32,
		A.a31*B.a13 + A.a32*B.a23 + A.a33*B.a33,
	}
}

func ReadInput() []Hailstone {
	hailstones := []Hailstone{}
	for _, line := range aocutil.ReadLines("input-ex1.txt") {
		nums := aocutil.GetIntsInString(line)
		hailstones = append(hailstones, Hailstone{Pos: IntVec3{nums[0], nums[1], nums[2]}, Vel: IntVec3{nums[3], nums[4], nums[5]}})
	}
	return hailstones
}
