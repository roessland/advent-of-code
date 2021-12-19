package main

import (
	"fmt"
	. "github.com/roessland/gopkg/mathutil"
)

var rotations []Mat

func init() {
	// Rotate 90 degrees around right hand axis
	R := Mat{
		1, 0, 0,
		0, 0, 1,
		0, -1, 0,
	}

	// Rotate 90 degrees around upwards axis
	U := Mat{
		0, 1, 0,
		-1, 0, 0,
		0, 0, 1,
	}

	// I used a Silva compass to visualize these rotations.
	// It's lying in front of me on the table, pointing to my right.
	baseRotations := []Mat{
		Identity(),                      // Laying flat
		R.MulMat(R),                     // Laying flat upside down
		U.MulMat(R),                     // Standing on flat short edge
		U.MulMat(R).MulMat(R).MulMat(R), // Standing on rounded edge
		R,                               // Standing on 1:25k edge
		U.MulMat(U).MulMat(R),           // Standing on 1:50k edge
	}

	// For each base rotation, add all rotations around up axis.
	Us := []Mat{Identity(), U, U.MulMat(U), U.MulMat(U).MulMat(U)}
	for _, baseRot := range baseRotations {
		for _, Ui := range Us {
			rotations = append(rotations, baseRot.MulMat(Ui))
		}
	}
}

type Vec struct {
	a0, a1, a2 int
}

func (U Vec) Add(V Vec) (W Vec) {
	return Vec{
		U.a0 + V.a0,
		U.a1 + V.a1,
		U.a2 + V.a2,
	}
}

func (U Vec) Neg() (V Vec) {
	return Vec{-U.a0, -U.a1, -U.a2}
}

func (U Vec) Sort() Vec {
	if U.a1 < U.a0 {
		U.a0, U.a1 = U.a1, U.a0
	}
	if U.a2 < U.a1 {
		U.a1, U.a2 = U.a2, U.a1
	}
	if U.a1 < U.a0 {
		U.a0, U.a1 = U.a1, U.a0
	}
	return U
}

func (U Vec) Sub(V Vec) (W Vec) {
	return Vec{
		U.a0 - V.a0,
		U.a1 - V.a1,
		U.a2 - V.a2,
	}
}

func (U Vec) Dist1(V Vec) int {
	return AbsInt(U.a0-V.a0) + AbsInt(U.a1-V.a1) + AbsInt(U.a2-V.a2)
}

func (U Vec) Abs() Vec {
	return Vec{AbsInt(U.a0), AbsInt(U.a1), AbsInt(U.a2)}
}

func (U Vec) Norm1() int {
	return AbsInt(U.a0) + AbsInt(U.a1) + AbsInt(U.a2)
}

type Mat struct {
	a00, a01, a02 int
	a10, a11, a12 int
	a20, a21, a22 int
}

func Identity() Mat {
	return Mat{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

func (A Mat) String() string {
	return fmt.Sprintf("[%4d\t%4d\t%4d\n %4d\t%4d\t%4d\n %4d\t%4d\t%4d  ]\n",
		A.a00, A.a01, A.a02,
		A.a10, A.a11, A.a12,
		A.a20, A.a21, A.a22)
}

func (A Mat) T() Mat {
	return Mat{
		A.a00, A.a10, A.a20,
		A.a01, A.a11, A.a21,
		A.a02, A.a12, A.a22,
	}
}

func (A Mat) MulMat(B Mat) (C Mat) {
	C.a00 = A.a00*B.a00 + A.a01*B.a10 + A.a02*B.a20
	C.a01 = A.a00*B.a01 + A.a01*B.a11 + A.a02*B.a21
	C.a02 = A.a00*B.a02 + A.a01*B.a12 + A.a02*B.a22
	C.a10 = A.a10*B.a00 + A.a11*B.a10 + A.a12*B.a20
	C.a11 = A.a10*B.a01 + A.a11*B.a11 + A.a12*B.a21
	C.a12 = A.a10*B.a02 + A.a11*B.a12 + A.a12*B.a22
	C.a20 = A.a20*B.a00 + A.a21*B.a10 + A.a22*B.a20
	C.a21 = A.a20*B.a01 + A.a21*B.a11 + A.a22*B.a21
	C.a22 = A.a20*B.a02 + A.a21*B.a12 + A.a22*B.a22
	return C
}

func (A Mat) MulVec(B Vec) (C Vec) {
	C.a0 = A.a00*B.a0 + A.a01*B.a1 + A.a02*B.a2
	C.a1 = A.a10*B.a0 + A.a11*B.a1 + A.a12*B.a2
	C.a2 = A.a20*B.a0 + A.a21*B.a1 + A.a22*B.a2
	return C
}
