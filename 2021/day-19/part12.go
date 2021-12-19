package main

import (
	"bufio"
	"fmt"
	. "github.com/roessland/gopkg/mathutil"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

type Scanner struct {
	Id       int
	Position Vec
	Rotation Mat
	Offsets  map[Vec]bool
}

func (s Scanner) String() string {
	return fmt.Sprintf("Scanner%d", s.Id)
}

func AllDistances(offsets map[Vec]bool) map[Vec]bool {
	distances := make(map[Vec]bool)
	for u := range offsets {
		for v := range offsets {
			distances[u.Sub(v).Abs().Sort()] = true
		}
	}
	return distances
}

func Distances(v Vec, offsets map[Vec]bool) map[Vec]bool {
	distances := make(map[Vec]bool)
	for u := range offsets {
		distances[u.Sub(v).Abs().Sort()] = true
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

func DistanceOverlap(dists1, dists2 map[Vec]bool) map[Vec]bool {
	common := map[Vec]bool{}
	for d := range dists1 {
		if dists2[d] {
			common[d] = true
		}
	}
	return common
}

func main() {

	f, err := os.Create("pprof.txt")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}

	defer pprof.StopCPUProfile()

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

	allDistances := make(map[*Scanner]map[Vec]bool)
	for _, scanner := range scanners {
		allDistances[scanner] = AllDistances(scanner.Offsets)
	}

	for len(scannersTodo) > 0 {
		scannerA := PickAnyFrom(scannersTodo)
		absDistsA := allDistances[scannerA]
		absPosA := AbsolutePositions(scannerA.Position, scannerA.Rotation, scannerA.Offsets)

		// Compare this scanner with all unknown scanners.
		for scannerB := range scannersUnknown {
			absDistsB := allDistances[scannerB]

			distanceOverlap := DistanceOverlap(absDistsA, absDistsB)

			// Optimization, not necessary
			if len(distanceOverlap) < 67 { // 12*11/2 is 67
				continue
			}

			// Just assume that some A offset and some B offset point
			// to the same beacon, and check if that assumption holds.
			for offA := range scannerA.Offsets {

				// Optimization, not necessary
				if len(DistanceOverlap(Distances(offA, scannerA.Offsets), distanceOverlap)) < 12 {
					continue
				}

				AX := scannerA.Rotation.MulVec(offA)
				OX := scannerA.Position.Add(AX)

				for offB := range scannerB.Offsets {

					// Optimization, not necessary
					if len(DistanceOverlap(Distances(offB, scannerB.Offsets), distanceOverlap)) < 12 {
						continue
					}

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
