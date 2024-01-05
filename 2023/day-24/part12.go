package main

import (
	"fmt"
	"math"
	"math/big"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

func main() {
	Part1()
}

type Hailstone struct {
	Pos BigRatVec3
	Vel BigRatVec3
}

type FloatVec2 struct {
	X, Y float64
}

type FloatVec3 struct {
	X, Y, Z float64
}

func (u FloatVec3) Sub(v FloatVec3) FloatVec3 {
	return FloatVec3{u.X - v.X, u.Y - v.Y, u.Z - v.Z}
}

func (u FloatVec3) Norm2() float64 {
	return u.X*u.X + u.Y*u.Y + u.Z*u.Z
}

func (u FloatVec3) Mul(alpha float64) FloatVec3 {
	return FloatVec3{u.X * alpha, u.Y * alpha, u.Z * alpha}
}

func (u FloatVec3) Add(v FloatVec3) FloatVec3 {
	return FloatVec3{u.X + v.X, u.Y + v.Y, u.Z + v.Z}
}

type BigRatVec2 struct {
	X, Y *big.Rat
}

func (u BigRatVec2) ToFloatVec2() FloatVec2 {
	if u.X == nil || u.Y == nil {
		panic("tofloatvec2 is nil")
	}
	x, _ := u.X.Float64()
	y, _ := u.Y.Float64()
	return FloatVec2{x, y}
}

type BigRatVec3 struct {
	X, Y, Z *big.Rat
}

type BigRatVec6 struct {
	v1, v2, v3, v4, v5, v6 *big.Rat
}

func NewBigRatVec3(x, y, z int) BigRatVec3 {
	return BigRatVec3{big.NewRat(int64(x), 1), big.NewRat(int64(y), 1), big.NewRat(int64(z), 1)}
}

func (u BigRatVec3) Sub(v BigRatVec3) BigRatVec3 {
	x := new(big.Rat).Sub(u.X, v.X)
	y := new(big.Rat).Sub(u.Y, v.Y)
	z := new(big.Rat).Sub(u.Z, v.Z)
	return BigRatVec3{x, y, z}
}

// Length2 returns the squared length of the vector.
func (u BigRatVec3) Length2() *big.Rat {
	x2 := new(big.Rat).Mul(u.X, u.X)
	y2 := new(big.Rat).Mul(u.Y, u.Y)
	z2 := new(big.Rat).Mul(u.Z, u.Z)
	sum := new(big.Rat)
	sum.Add(sum, x2)
	sum.Add(sum, y2)
	sum.Add(sum, z2)
	return sum
}

func (u BigRatVec3) Dot(v BigRatVec3) *big.Rat {
	x := new(big.Rat).Mul(u.X, v.X)
	y := new(big.Rat).Mul(u.Y, v.Y)
	z := new(big.Rat).Mul(u.Z, v.Z)
	sum := new(big.Rat)
	sum.Add(sum, x)
	sum.Add(sum, y)
	sum.Add(sum, z)
	return sum
}

// Quo return the quotient u/alpha
// If alpha == 0, Quo panics.
func (u BigRatVec3) Quo(alpha *big.Rat) BigRatVec3 {
	return BigRatVec3{
		new(big.Rat).Quo(u.X, alpha),
		new(big.Rat).Quo(u.Y, alpha),
		new(big.Rat).Quo(u.Z, alpha),
	}
}

func (u BigRatVec3) ToFloatVec3() FloatVec3 {
	x, _ := u.X.Float64()
	y, _ := u.Y.Float64()
	z, _ := u.Z.Float64()
	if math.IsInf(x, 0) || math.IsInf(y, 0) || math.IsInf(z, 0) {
		msg := fmt.Sprintf("Could not convert to float: %v", u)
		panic(msg)
	}
	return FloatVec3{x, y, z}
}

func (u BigRatVec3) String() string {
	return fmt.Sprintf("(%v, %v, %v)", u.X.RatString(), u.Y.RatString(), u.Z.RatString())
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
		hs[i].Pos.Z = big.NewRat(0, 1)
		hs[i].Vel.Z = big.NewRat(0, 1)
	}

	areaMin := big.NewRat(200000000000000, 1)
	areaMax := big.NewRat(400000000000000, 1)
	// areaMin := big.NewRat(7, 1)
	// areaMax := big.NewRat(27, 1)
	count := 0

	for iA, hA := range hs {
		for iB, hB := range hs {
			if iA >= iB {
				continue
			}
			fmt.Println("Checking", hA.Pos, hB.Pos)
			p, intersects := SegmentSegmentIntersection2D(hA, hB)
			fmt.Println("\tIntersects:", intersects)
			if !intersects {
				continue
			}
			if intersects {
				fmt.Println("\tIntersection:", p.ToFloatVec2())
			}
			fmt.Println("Checking", hA, hB, p, intersects)
			if !InsideArea2D(p, areaMin, areaMax) {
				fmt.Println("\tIntersection is outside area")
				continue
			}
			if !IsFuture2D(hA, p) {
				fmt.Println("\tIntersection is not in A's future")
				continue
			}
			if !IsFuture2D(hB, p) {
				fmt.Println("\tIntersection is not in B's future")
				continue
			}

			fmt.Println("WILL CROSS INSIDE TEST AREA")
			count++
		}
	}
	fmt.Println("Part 1:", count)
}

func InsideArea2D(p BigRatVec2, areaMin, areaMax *big.Rat) bool {
	if p.X.Cmp(areaMin) < 0 {
		return false
	}
	if p.X.Cmp(areaMax) > 0 {
		return false
	}
	if p.Y.Cmp(areaMin) < 0 {
		return false
	}
	if p.Y.Cmp(areaMax) > 0 {
		return false
	}
	return true
}

func InsideArea3D(p BigRatVec3, areaMin, areaMax *big.Rat) bool {
	if p.X.Cmp(areaMin) < 0 {
		return false
	}
	if p.X.Cmp(areaMax) > 0 {
		return false
	}
	if p.Y.Cmp(areaMin) < 0 {
		return false
	}
	if p.Y.Cmp(areaMax) > 0 {
		return false
	}
	if p.Z.Cmp(areaMin) < 0 {
		return false
	}
	if p.Z.Cmp(areaMax) > 0 {
		return false
	}
	return true
}

func IsFuture2D(hs Hailstone, coord BigRatVec2) bool {
	// x0 + t * x' = x
	// => t = (x - x0) / x'
	// sign(t) = sign(x - x0) == sign(x')
	dx := new(big.Rat).Set(hs.Pos.X)
	dx.Sub(dx, coord.X)
	dx.Neg(dx)

	if dx.Sign() == 0 {
		panic("dx is zero")
	}
	return dx.Sign() == hs.Vel.X.Sign()
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func SegmentSegmentIntersection3D(A, B Hailstone) (p BigRatVec3, intersects bool) {
	return LineLineIntersection3D(A, B)
}

func SegmentSegmentIntersection2D(hA, hB Hailstone) (q BigRatVec2, intersects bool) {
	return LineLineIntersection2D(hA, hB)
}

func LineLineIntersection2D(hA, hB Hailstone) (p BigRatVec2, intersects bool) {
	// Matrix form (z = 0), 2 dimensions:
	//
	// [ yA'  -xA' ]     [ x ]     [ xA yA' - yA xA' ]
	// [ yB'  -xB' ]  @  [ y ]  =  [ xB yB' - yB xB' ]
	//
	// Example:
	//  y = ax + c
	//  y = cx + d
	//  =>  ax + c = cx + d  =>  ax - cx = d - c  =>  x(a - c) = d-c
	//  =>  x = (d-c) / (a-c)
	//
	// Parametric equations:
	//  A: (x,y) = (xA, yA) + s * (xA', yA')
	//  B: (x,y) = (xB, yB) + t * (xB', yB')
	//
	// Convert to symmetric form:
	//  x = xA + s * xA'  =>  s = (x - xA) / xA'
	//  y = yA + s * yA'  =>  s = (y - yA) / yA'
	//  x = xB + t * xB'  =>  t = (x - xB) / xB'
	//  y = yB + t * yB'  =>  t = (y - yB) / yB'
	//
	//   I: (x - xA) / xA' = (y - yA) / yA'
	//  II: (x - xB) / xB' = (y - yB) / yB'
	//
	// Convert to matrix form:
	//  (x - xA) * yA' - (y - yA) * xA' = 0
	//  (x - xB) * yB' - (y - yB) * xB' = 0
	//
	//  x yA' - xA yA' - y xA' - yA xA' = 0
	//  x yB' - xB yB' + y xB' - yB xB' = 0
	//
	//
	//  x yA' + y xA' = xA yA' - yA xA'
	//  x yB' + y xB' = xB yB' - yB xB'
	//
	//  [ yA'  -xA' ]     [ x ]     [ xA yA' - yA xA' ]
	//  [ yB'  -xB' ]  @  [ y ]  =  [ xB yB' - yB xB' ]
	//
	// Solve manually:
	//
	//
	//
	a11 := hA.Vel.Y
	a12 := new(big.Rat).Neg(hA.Vel.X)
	a21 := hB.Vel.Y
	a22 := new(big.Rat).Neg(hB.Vel.X)

	A := Mat2{
		a11, a12,
		a21, a22,
	}

	//  [ xA yA' - yA xA' ]
	//  [ xB yB' - yB xB' ]
	b1 := new(big.Rat)
	b1.Add(b1, new(big.Rat).Mul(hA.Pos.X, hA.Vel.Y))
	b1.Sub(b1, new(big.Rat).Mul(hA.Pos.Y, hA.Vel.X))
	b2 := new(big.Rat)
	b2.Add(b2, new(big.Rat).Mul(hB.Pos.X, hB.Vel.Y))
	b2.Sub(b2, new(big.Rat).Mul(hB.Pos.Y, hB.Vel.X))
	b := BigRatVec2{X: b1, Y: b2}

	D := A.Det()

	if D.Num().Sign() == 0 {
		return BigRatVec2{}, false
	}

	if b.X == nil || b.Y == nil {
		panic("issa nil")
	}
	return A.Solve(b)
}

func (A Mat2) Solve(b BigRatVec2) (BigRatVec2, bool) {
	a11, a12, a21, a22 := A.a11, A.a12, A.a21, A.a22

	D := A.Det()

	Dx := Mat2{
		b.X, a12,
		b.Y, a22,
	}.Det()

	Dy := Mat2{
		a11, b.X,
		a21, b.Y,
	}.Det()

	if D.Num().Sign() == 0 {
		return BigRatVec2{}, false
	}

	x := new(big.Rat).Quo(Dx, D)
	y := new(big.Rat).Quo(Dy, D)

	return BigRatVec2{x, y}, true
}

func LineLineIntersection3D(hA, hB Hailstone) (p BigRatVec3, intersects bool) {
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
	// Matrix form (z = 0), 2 dimensions:
	//
	// [ yA'  -xA' ]     [ x ]     [ xA yA' - yA xA' ]
	// [ yB'  -xB' ]  @  [ y ]  =  [ xB yB' - yB xB' ]
	//
	// Example:
	// > Hailstone A: 19, 13, 30 @ -2, 1, -2
	// > Hailstone B: 18, 19, 22 @ -1, -1, -2
	// > Hailstones' paths will cross inside the test area (at x=14.333, y=15.333)
	// [ 1   2  ]               [ 19 * 1 + 13 * 2 ]
	// [ -1  1  ]               [ -18 * 1 + 19 * 1]
	// 14.333 + 2*15.333 = 44.999 = 45
	// -14.333 + 15.333 = 1 = 1

	A := Mat63{
		hA.Vel.Y, new(big.Rat).Neg(hA.Vel.X), big.NewRat(0, 1),
		hA.Vel.Z, big.NewRat(0, 1), new(big.Rat).Neg(hA.Vel.X),
		big.NewRat(0, 1), hA.Vel.Z, new(big.Rat).Neg(hA.Vel.Y),
		hB.Vel.Y, new(big.Rat).Neg(hB.Vel.X), big.NewRat(0, 1),
		hB.Vel.Z, big.NewRat(0, 1), new(big.Rat).Neg(hB.Vel.X),
		big.NewRat(0, 1), hB.Vel.Z, new(big.Rat).Neg(hB.Vel.Y),
	}

	AtA := A.AtA()

	// Solve A'A w = A'b for w
	// A 6x3
	// A' 3x6
	// A'A 3x3
	// b 6x1
	// A'b 3x1
	b1 := new(big.Rat).Add(new(big.Rat).Mul(hA.Pos.X, hA.Vel.Y), new(big.Rat).Mul(hA.Pos.Y, new(big.Rat).Neg(hA.Vel.X)))
	b2 := new(big.Rat).Add(new(big.Rat).Mul(hA.Pos.X, hA.Vel.Z), new(big.Rat).Mul(hA.Pos.Z, new(big.Rat).Neg(hA.Vel.X)))
	b3 := new(big.Rat).Add(new(big.Rat).Mul(hA.Pos.Y, hA.Vel.Z), new(big.Rat).Mul(hA.Pos.Z, new(big.Rat).Neg(hA.Vel.Y)))
	b4 := new(big.Rat).Add(new(big.Rat).Mul(hB.Pos.X, hB.Vel.Y), new(big.Rat).Mul(hB.Pos.Y, new(big.Rat).Neg(hB.Vel.X)))
	b5 := new(big.Rat).Add(new(big.Rat).Mul(hB.Pos.X, hB.Vel.Z), new(big.Rat).Mul(hB.Pos.Z, new(big.Rat).Neg(hB.Vel.X)))
	b6 := new(big.Rat).Add(new(big.Rat).Mul(hB.Pos.Y, hB.Vel.Z), new(big.Rat).Mul(hB.Pos.Z, new(big.Rat).Neg(hB.Vel.Y)))
	b := BigRatVec6{
		b1, b2, b3, b4, b5, b6,
	}
	// b := BigRatVec6{
	// 	hA.Pos.X*hA.Vel.Y - hA.Pos.Y*hA.Vel.X,
	// 	hA.Pos.X*hA.Vel.Z - hA.Pos.Z*hA.Vel.X,
	// 	hA.Pos.Y*hA.Vel.Z - hA.Pos.Z*hA.Vel.Y,
	// 	hB.Pos.X*hB.Vel.Y - hB.Pos.Y*hB.Vel.X,
	// 	hB.Pos.X*hB.Vel.Z - hB.Pos.Z*hB.Vel.X,
	// 	hB.Pos.Y*hB.Vel.Z - hB.Pos.Z*hB.Vel.Y,
	// }

	Atb := A.Atb(b)

	// LHS: A'A * w
	// RHS: A'b
	// Solve for w
	w := SolveSystem(AtA, Atb)
	fmt.Println("w", w)

	// [ yA'  -xA'    0 ]               [ xA yA' - yA xA' ]
	// [ zA'    0   -xA']               [ xA zA' - zA xA' ]
	// [ 0     zA'  -yA']     [ x ]     [ yA zA' - zA yA' ]
	// [ yB'  -xB'    0 ]  @  [ y ]  =  [ xB yB' - yB xB' ]
	// [ zB'    0   -xB']     [ z ]     [ xB zB' - zB xB' ]
	// [ 0     zB'  -yB']               [ yB zB' - zB yB' ]
	//
	if w == nil {
		return BigRatVec3{}, false
	}
	return *w, true
}

type Mat2 struct {
	a11, a12 *big.Rat
	a21, a22 *big.Rat
}

func (m Mat2) Det() *big.Rat {
	// a11 a22 - a12 a21
	s1 := new(big.Rat).Mul(m.a11, m.a22)
	s2 := new(big.Rat).Mul(m.a12, m.a21)
	s := new(big.Rat).Sub(s1, s2)
	return s
}

type Mat33 struct {
	a11, a12, a13 *big.Rat
	a21, a22, a23 *big.Rat
	a31, a32, a33 *big.Rat
}

func Det(m Mat33) *big.Rat {
	a11, a12, a13, a21, a22, a23, a31, a32, a33 := m.a11, m.a12, m.a13, m.a21, m.a22, m.a23, m.a31, m.a32, m.a33
	// s = s1 + s2 + s3
	s1 := new(big.Rat)
	s1 = s1.Add(s1, new(big.Rat).Mul(a11, new(big.Rat).Mul(a22, a33)))
	s2 := new(big.Rat)
	s2 = s2.Add(s2, new(big.Rat).Mul(a12, new(big.Rat).Mul(a23, a31)))
	s3 := new(big.Rat)
	s3 = s3.Add(s3, new(big.Rat).Mul(a13, new(big.Rat).Mul(a21, a32)))

	s := new(big.Rat)
	s.Add(s, s1)
	s.Add(s, s2)
	s.Add(s, s3)
	return s
}

func (m Mat33) String() string {
	return fmt.Sprintf("[%v, %v, %v; %v, %v, %v; %v, %v, %v]",
		m.a11.RatString(), m.a12.RatString(), m.a13.RatString(),
		m.a21.RatString(), m.a22.RatString(), m.a23.RatString(),
		m.a31.RatString(), m.a32.RatString(), m.a33.RatString(),
	)
}

// SolveSystem solves the system Ax=b for x, where A is a 3x3 matrix, b is a
// 3x1 vector, and x is a 3x1 vector, using Cramer's rule.
func SolveSystem(A Mat33, b BigRatVec3) *BigRatVec3 {
	fmt.Println("\nSolving", A, b)
	D := Det(A)
	if D.Num().Sign() == 0 {
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

	r := BigRatVec3{Dx, Dy, Dz}.Quo(D)

	return &r
}

func T(A Mat33) Mat33 {
	return Mat33{A.a11, A.a21, A.a31, A.a12, A.a22, A.a32, A.a13, A.a23, A.a33}
}

func MatMatMul(A, B Mat33) Mat33 {
	m11 := new(big.Rat)
	m11.Add(m11, new(big.Rat).Mul(A.a11, B.a11))
	m11.Add(m11, new(big.Rat).Mul(A.a12, B.a21))
	m11.Add(m11, new(big.Rat).Mul(A.a13, B.a31))

	m12 := new(big.Rat)
	m12.Add(m12, new(big.Rat).Mul(A.a11, B.a12))
	m12.Add(m12, new(big.Rat).Mul(A.a12, B.a22))
	m12.Add(m12, new(big.Rat).Mul(A.a13, B.a32))

	m13 := new(big.Rat)
	m13.Add(m13, new(big.Rat).Mul(A.a11, B.a13))
	m13.Add(m13, new(big.Rat).Mul(A.a12, B.a23))
	m13.Add(m13, new(big.Rat).Mul(A.a13, B.a33))

	m21 := new(big.Rat)
	m21.Add(m21, new(big.Rat).Mul(A.a21, B.a11))
	m21.Add(m21, new(big.Rat).Mul(A.a22, B.a21))
	m21.Add(m21, new(big.Rat).Mul(A.a23, B.a31))

	m22 := new(big.Rat)
	m22.Add(m22, new(big.Rat).Mul(A.a21, B.a12))
	m22.Add(m22, new(big.Rat).Mul(A.a22, B.a22))
	m22.Add(m22, new(big.Rat).Mul(A.a23, B.a32))

	m23 := new(big.Rat)
	m23.Add(m23, new(big.Rat).Mul(A.a21, B.a13))
	m23.Add(m23, new(big.Rat).Mul(A.a22, B.a23))
	m23.Add(m23, new(big.Rat).Mul(A.a23, B.a33))

	m31 := new(big.Rat)
	m31.Add(m31, new(big.Rat).Mul(A.a31, B.a11))
	m31.Add(m31, new(big.Rat).Mul(A.a32, B.a21))
	m31.Add(m31, new(big.Rat).Mul(A.a33, B.a31))

	m32 := new(big.Rat)
	m32.Add(m32, new(big.Rat).Mul(A.a31, B.a12))
	m32.Add(m32, new(big.Rat).Mul(A.a32, B.a22))
	m32.Add(m32, new(big.Rat).Mul(A.a33, B.a32))

	m33 := new(big.Rat)
	m33.Add(m33, new(big.Rat).Mul(A.a31, B.a13))
	m33.Add(m33, new(big.Rat).Mul(A.a32, B.a23))
	m33.Add(m33, new(big.Rat).Mul(A.a33, B.a33))

	return Mat33{
		m11, m12, m13,
		m21, m22, m23,
		m31, m32, m33,
	}

	// return Mat33{
	// 	A.a11*B.a11 + A.a12*B.a21 + A.a13*B.a31,
	// 	A.a11*B.a12 + A.a12*B.a22 + A.a13*B.a32,
	// 	A.a11*B.a13 + A.a12*B.a23 + A.a13*B.a33,
	// 	A.a21*B.a11 + A.a22*B.a21 + A.a23*B.a31,
	// 	A.a21*B.a12 + A.a22*B.a22 + A.a23*B.a32,
	// 	A.a21*B.a13 + A.a22*B.a23 + A.a23*B.a33,
	// 	A.a31*B.a11 + A.a32*B.a21 + A.a33*B.a31,
	// 	A.a31*B.a12 + A.a32*B.a22 + A.a33*B.a32,
	// 	A.a31*B.a13 + A.a32*B.a23 + A.a33*B.a33,
	// }
}

func (A Mat33) Mul(v BigRatVec3) BigRatVec3 {
	b1 := new(big.Rat)
	b1.Add(b1, new(big.Rat).Mul(A.a11, v.X))
	b1.Add(b1, new(big.Rat).Mul(A.a12, v.Y))
	b1.Add(b1, new(big.Rat).Mul(A.a13, v.Z))

	b2 := new(big.Rat)
	b2.Add(b2, new(big.Rat).Mul(A.a21, v.X))
	b2.Add(b2, new(big.Rat).Mul(A.a22, v.Y))
	b2.Add(b2, new(big.Rat).Mul(A.a23, v.Z))

	b3 := new(big.Rat)
	b3.Add(b3, new(big.Rat).Mul(A.a31, v.X))
	b3.Add(b3, new(big.Rat).Mul(A.a32, v.Y))
	b3.Add(b3, new(big.Rat).Mul(A.a33, v.Z))

	// return BigRatVec3{
	// 	A.a11*v.X + A.a12*v.Y + A.a13*v.Z,
	// 	A.a21*v.X + A.a22*v.Y + A.a23*v.Z,
	// 	A.a31*v.X + A.a32*v.Y + A.a33*v.Z,
	// }

	return BigRatVec3{b1, b2, b3}
}

type Mat63 struct {
	a11, a12, a13 *big.Rat
	a21, a22, a23 *big.Rat
	a31, a32, a33 *big.Rat
	a41, a42, a43 *big.Rat
	a51, a52, a53 *big.Rat
	a61, a62, a63 *big.Rat
}

// AtA returns the matrix A^T A, given A.
func (m Mat63) AtA() Mat33 {
	a11, a12, a13, a21, a22, a23, a31, a32, a33 := m.a11, m.a12, m.a13, m.a21, m.a22, m.a23, m.a31, m.a32, m.a33
	a41, a42, a43, a51, a52, a53, a61, a62, a63 := m.a41, m.a42, m.a43, m.a51, m.a52, m.a53, m.a61, m.a62, m.a63

	m11 := new(big.Rat)
	m11.Add(m11, new(big.Rat).Mul(a11, a11))
	m11.Add(m11, new(big.Rat).Mul(a21, a21))
	m11.Add(m11, new(big.Rat).Mul(a31, a31))
	m11.Add(m11, new(big.Rat).Mul(a41, a41))
	m11.Add(m11, new(big.Rat).Mul(a51, a51))
	m11.Add(m11, new(big.Rat).Mul(a61, a61))

	m12 := new(big.Rat)
	m12.Add(m12, new(big.Rat).Mul(a11, a12))
	m12.Add(m12, new(big.Rat).Mul(a21, a22))
	m12.Add(m12, new(big.Rat).Mul(a31, a32))
	m12.Add(m12, new(big.Rat).Mul(a41, a42))
	m12.Add(m12, new(big.Rat).Mul(a51, a52))
	m12.Add(m12, new(big.Rat).Mul(a61, a62))

	m13 := new(big.Rat)
	m13.Add(m13, new(big.Rat).Mul(a11, a13))
	m13.Add(m13, new(big.Rat).Mul(a21, a23))
	m13.Add(m13, new(big.Rat).Mul(a31, a33))
	m13.Add(m13, new(big.Rat).Mul(a41, a43))
	m13.Add(m13, new(big.Rat).Mul(a51, a53))
	m13.Add(m13, new(big.Rat).Mul(a61, a63))

	m21 := new(big.Rat)
	m21.Add(m21, new(big.Rat).Mul(a12, a11))
	m21.Add(m21, new(big.Rat).Mul(a22, a21))
	m21.Add(m21, new(big.Rat).Mul(a32, a31))
	m21.Add(m21, new(big.Rat).Mul(a42, a41))
	m21.Add(m21, new(big.Rat).Mul(a52, a51))
	m21.Add(m21, new(big.Rat).Mul(a62, a61))

	m22 := new(big.Rat)
	m22.Add(m22, new(big.Rat).Mul(a12, a12))
	m22.Add(m22, new(big.Rat).Mul(a22, a22))
	m22.Add(m22, new(big.Rat).Mul(a32, a32))
	m22.Add(m22, new(big.Rat).Mul(a42, a42))
	m22.Add(m22, new(big.Rat).Mul(a52, a52))
	m22.Add(m22, new(big.Rat).Mul(a62, a62))

	m23 := new(big.Rat)
	m23.Add(m23, new(big.Rat).Mul(a12, a13))
	m23.Add(m23, new(big.Rat).Mul(a22, a23))
	m23.Add(m23, new(big.Rat).Mul(a32, a33))
	m23.Add(m23, new(big.Rat).Mul(a42, a43))
	m23.Add(m23, new(big.Rat).Mul(a52, a53))
	m23.Add(m23, new(big.Rat).Mul(a62, a63))

	m31 := new(big.Rat)
	m31.Add(m31, new(big.Rat).Mul(a13, a11))
	m31.Add(m31, new(big.Rat).Mul(a23, a21))
	m31.Add(m31, new(big.Rat).Mul(a33, a31))
	m31.Add(m31, new(big.Rat).Mul(a43, a41))
	m31.Add(m31, new(big.Rat).Mul(a53, a51))
	m31.Add(m31, new(big.Rat).Mul(a63, a61))

	m32 := new(big.Rat)
	m32.Add(m32, new(big.Rat).Mul(a13, a12))
	m32.Add(m32, new(big.Rat).Mul(a23, a22))
	m32.Add(m32, new(big.Rat).Mul(a33, a32))
	m32.Add(m32, new(big.Rat).Mul(a43, a42))
	m32.Add(m32, new(big.Rat).Mul(a53, a52))
	m32.Add(m32, new(big.Rat).Mul(a63, a62))

	m33 := new(big.Rat)
	m33.Add(m33, new(big.Rat).Mul(a13, a13))
	m33.Add(m33, new(big.Rat).Mul(a23, a23))
	m33.Add(m33, new(big.Rat).Mul(a33, a33))
	m33.Add(m33, new(big.Rat).Mul(a43, a43))
	m33.Add(m33, new(big.Rat).Mul(a53, a53))
	m33.Add(m33, new(big.Rat).Mul(a63, a63))

	return Mat33{
		m11, m12, m13,
		m21, m22, m23,
		m31, m32, m33,
	}
}

// Atb returns the vector A^T b, given A and b.
func (m Mat63) Atb(v BigRatVec6) BigRatVec3 {
	a11, a12, a13, a21, a22, a23, a31, a32, a33 := m.a11, m.a12, m.a13, m.a21, m.a22, m.a23, m.a31, m.a32, m.a33
	a41, a42, a43, a51, a52, a53, a61, a62, a63 := m.a41, m.a42, m.a43, m.a51, m.a52, m.a53, m.a61, m.a62, m.a63
	v1, v2, v3, v4, v5, v6 := v.v1, v.v2, v.v3, v.v4, v.v5, v.v6

	r1 := new(big.Rat)
	r1.Add(r1, new(big.Rat).Mul(a11, v1))
	r1.Add(r1, new(big.Rat).Mul(a21, v2))
	r1.Add(r1, new(big.Rat).Mul(a31, v3))
	r1.Add(r1, new(big.Rat).Mul(a41, v4))
	r1.Add(r1, new(big.Rat).Mul(a51, v5))
	r1.Add(r1, new(big.Rat).Mul(a61, v6))

	r2 := new(big.Rat)
	r2.Add(r2, new(big.Rat).Mul(a12, v1))
	r2.Add(r2, new(big.Rat).Mul(a22, v2))
	r2.Add(r2, new(big.Rat).Mul(a32, v3))
	r2.Add(r2, new(big.Rat).Mul(a42, v4))
	r2.Add(r2, new(big.Rat).Mul(a52, v5))
	r2.Add(r2, new(big.Rat).Mul(a62, v6))

	r3 := new(big.Rat)
	r3.Add(r3, new(big.Rat).Mul(a13, v1))
	r3.Add(r3, new(big.Rat).Mul(a23, v2))
	r3.Add(r3, new(big.Rat).Mul(a33, v3))
	r3.Add(r3, new(big.Rat).Mul(a43, v4))
	r3.Add(r3, new(big.Rat).Mul(a53, v5))
	r3.Add(r3, new(big.Rat).Mul(a63, v6))

	return BigRatVec3{r1, r2, r3}
	// 	a11*v1 + a21*v2 + a31*v3 + a41*v4 + a51*v5 + a61*v6,
	// 	a12*v1 + a22*v2 + a32*v3 + a42*v4 + a52*v5 + a62*v6,
	// 	a13*v1 + a23*v2 + a33*v3 + a43*v4 + a53*v5 + a63*v6,
	// }
}

func ReadInput() []Hailstone {
	hailstones := []Hailstone{}
	for _, line := range aocutil.ReadLines("input.txt") {
		nums := aocutil.GetIntsInString(line)
		pos := NewBigRatVec3(nums[0], nums[1], nums[2])
		vel := NewBigRatVec3(nums[3], nums[4], nums[5])
		hailstones = append(hailstones, Hailstone{Pos: pos, Vel: vel})
	}
	return hailstones
}

// 1155 too low
