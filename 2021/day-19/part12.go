package main

import (
	"bufio"
	"fmt"
	. "github.com/roessland/gopkg/mathutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

type Scanner struct {
	Id       int
	Position Vec
	Rotation Mat
	Offsets  map[Vec]bool
}

func (s Scanner) String() string {
	return fmt.Sprintf("Scanner%d", s.Id)
}

func Distances(offsets map[Vec]bool) map[Vec]bool {
	distances := make(map[Vec]bool)
	for u := range offsets {
		for v := range offsets {
			distances[u.Sub(v).Abs().Sort()] = true
		}
	}
	return distances
}

func AbsolutePositions(pos Vec, rot Mat, offsets map[Vec]bool) map[Vec]bool {
	absolutePositions := make(map[Vec]bool)
	for offset := range offsets {
		pos := pos.Add(rot.MulVec(offset))
		absolutePositions[pos] = true
	}
	return absolutePositions
}

func PositionOverlap(absPos1, absPos2 map[Vec]bool) int {
	common := 0
	for p1 := range absPos1 {
		if absPos2[p1] {
			common++
		}
	}
	return common
}

func DistanceOverlap(dists1, dists2 map[Vec]bool) int {
	common := 0
	for d := range dists1 {
		if dists2[d] {
			common++
		}
	}
	return common
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

func main() {
	t0 := time.Now()

	scanners := ReadInput()
	LocateScanners(scanners)

	beaconPositions := make(map[Vec]struct{})
	for _, scanner := range scanners {
		for offset := range scanner.Offsets {
			pos := scanner.Position.Add(scanner.Rotation.MulVec(offset))
			beaconPositions[pos] = struct{}{}
		}
	}
	fmt.Println("Part 1:", len(beaconPositions))

	maxDist := 0
	for _, a := range scanners {
		for _, b := range scanners {
			maxDist = MaxInt(maxDist, a.Position.Dist1(b.Position))
		}
	}
	fmt.Println("Part 2:", maxDist)

	fmt.Println(time.Since(t0))
}

func LocateScanners(scanners []*Scanner) {
	var scannersDone = make(map[*Scanner]struct{})

	var scannersTodo = map[*Scanner]struct{}{
		scanners[0]: {},
	}

	var scannersUnknown = make(map[*Scanner]struct{})
	for _, scanner := range scanners[1:] {
		scannersUnknown[scanner] = struct{}{}
	}

	distances := make(map[*Scanner]map[Vec]bool)
	for _, scanner := range scanners {
		distances[scanner] = Distances(scanner.Offsets)
	}

	for len(scannersTodo) > 0 {
		scannerA := PickAnyFrom(scannersTodo)
		absDistsA := distances[scannerA]
		absPosA := AbsolutePositions(scannerA.Position, scannerA.Rotation, scannerA.Offsets)

		// Compare this scanner with all unknown scanners.
		for scannerB := range scannersUnknown {
			absDistsB := distances[scannerB]

			if DistanceOverlap(absDistsA, absDistsB) < 67 { // 12*11/2 is 67
				continue
			}

			// Just assume that some A offset and some B offset point
			// to the same beacon, and check if that assumption holds.
			for offA := range scannerA.Offsets {
				AX := scannerA.Rotation.MulVec(offA)
				OX := scannerA.Position.Add(AX)

				for offB := range scannerB.Offsets {
					for _, BRot := range rotations {
						BX := BRot.MulVec(offB)
						XB := BX.Neg()
						BPos := OX.Add(XB)
						absPosB := AbsolutePositions(BPos, BRot, scannerB.Offsets)
						if PositionOverlap(absPosA, absPosB) >= 12 {
							scannerB.Position = BPos
							scannerB.Rotation = BRot
							scannersTodo[scannerB] = struct{}{}
							delete(scannersUnknown, scannerB)
						}
					}
				}
			}
		}

		scannersDone[scannerA] = struct{}{}
		delete(scannersTodo, scannerA)
	}

	if len(scannersUnknown) != 0 {
		log.Println("didnt find them all")
	}
	return
}

func ReadInput() []*Scanner {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("lmao")
	}
	input := bufio.NewScanner(f)
	var scanners []*Scanner
	var scanner *Scanner
	id := 0
	for input.Scan() {
		line := input.Text()
		if strings.HasPrefix(line, "---") {
			scanner = &Scanner{Id: id, Rotation: Identity(), Offsets: make(map[Vec]bool)}
			id++
			continue
		}
		if line == "" {
			scanners = append(scanners, scanner)
			continue
		}
		numsStr := strings.Split(line, ",")
		a0, _ := strconv.Atoi(numsStr[0])
		a1, _ := strconv.Atoi(numsStr[1])
		a2, _ := strconv.Atoi(numsStr[2])
		scanner.Offsets[Vec{a0, a1, a2}] = true
	}

	scanners = append(scanners, scanner)
	return scanners
}

func PickAnyFrom(set map[*Scanner]struct{}) *Scanner {
	for scanner := range set {
		return scanner
	}
	panic("empty set")
}
