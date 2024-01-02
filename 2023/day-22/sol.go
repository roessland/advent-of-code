package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
	"sort"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

type Coord2 struct {
	X, Y int
}

type Coord3 struct {
	X, Y, Z int
}

type Brick struct {
	Name   string
	ID     int
	Origin Coord3
	Size   Coord3
}

// A is the smallest coordinate included in the brick.
func (br Brick) A() Coord3 {
	return br.Origin
}

// B is the largest coordinate included in the brick.
func (br Brick) B() Coord3 {
	return Coord3{br.Origin.X + br.Size.X - 1, br.Origin.Y + br.Size.Y - 1, br.Origin.Z + br.Size.Z - 1}
}

func (br Brick) Contains(c Coord3) bool {
	x, y, z := c.X, c.Y, c.Z
	a := br.A()
	b := br.B()
	return a.X <= x && x <= b.X && a.Y <= y && y <= b.Y && a.Z <= z && z <= b.Z
}

func (br Brick) ContainsXY(c Coord2) bool {
	x, y := c.X, c.Y
	a := br.A()
	b := br.B()
	return a.X <= x && x <= b.X && a.Y <= y && y <= b.Y
}

func (br Brick) ColsXY() []Coord2 {
	cols := []Coord2{}
	a := br.A()
	b := br.B()
	for x := a.X; x <= b.X; x++ {
		for y := a.Y; y <= b.Y; y++ {
			cols = append(cols, Coord2{x, y})
		}
	}
	return cols
}

// IntersectsXY returns true if the two bricks intersect in the XY plane.
// >       y
// >       ▲            ┃
// >       │           │┃
// >       └───▶x      │▼
// >  ┌─┐  ┏━━━┓┌───┐  ▼
// > ┏╋┓│  ┃┌─┐┃│┏━┓│
// > ┃└╋┘  ┃└─┘┃│┗━┛│  │
// > ┗━┛   ┗━━━┛└───┘  │┃
// > ┏━┓   ┌─┐   ┏━┓   ▼┃
// > ┃┌╋┐ ┏╋━╋┓ ┌╋─╋┐   ▼
// > ┗╋┛│ ┗╋━╋┛ └╋─╋┘
// >  └─┘  └─┘   ┗━┛   │
// > ┌─┐   ┏━┓         │┃
// > │┏╋┓ ┌╋┐┃         │▼
// > └╋┘┃ │┗╋┛         ▼
// >  ┗━┛ └─┘
// >                    ┃
// >                   │┃
// >                   ▼┃
// >                    ▼
func (p Brick) IntersectsXY(q Brick) bool {
	// Basic idea: Consider each square as a pair of intervals on the X and Y
	// axes, respectively. Both must intersect for the bricks to intersect.
	pA, pB, qA, qB := p.A(), p.B(), q.A(), q.B()
	pIntX, pIntY := Interval{pA.X, pB.X}, Interval{pA.Y, pB.Y}
	qIntX, qIntY := Interval{qA.X, qB.X}, Interval{qA.Y, qB.Y}
	if !pIntX.Intersects(qIntX) {
		return false
	}
	if !pIntY.Intersects(qIntY) {
		return false
	}
	return true
}

func (p Brick) Intersects(q Brick) bool {
	// Basic idea: Consider each square as a pair of intervals on the X and Y
	// axes, respectively. Both must intersect for the bricks to intersect.
	pA, pB, qA, qB := p.A(), p.B(), q.A(), q.B()
	pIntX, pIntY, pIntZ := Interval{pA.X, pB.X}, Interval{pA.Y, pB.Y}, Interval{pA.Z, pB.Z}
	qIntX, qIntY, qIntZ := Interval{qA.X, qB.X}, Interval{qA.Y, qB.Y}, Interval{qA.Z, qB.Z}
	if !pIntX.Intersects(qIntX) {
		return false
	}
	if !pIntY.Intersects(qIntY) {
		return false
	}
	if !pIntZ.Intersects(qIntZ) {
		return false
	}
	return true
}

// Interval is an inclusive interval.
type Interval struct {
	Start, End int
}

func (i Interval) Intersects(o Interval) bool {
	// i contains o
	if i.Start <= o.Start && o.End <= i.End {
		return true
	}
	// o contains i
	if o.Start <= i.Start && i.End <= o.End {
		return true
	}

	// i starts or ends in o
	if o.Start <= i.Start && i.Start <= o.End || o.Start <= i.End && i.End <= o.End {
		return true
	}

	// o starts or ends in i
	if i.Start <= o.Start && o.Start <= i.End || i.Start <= o.End && o.End <= i.End {
		panic("shouldn't happen")
	}
	return false
}

func (br Brick) ContainsXZ(c Coord2) bool {
	x, z := c.X, c.Y
	a := br.A()
	b := br.B()
	return a.X <= x && x <= b.X && a.Z <= z && z <= b.Z
}

func Fall(bricks_ []Brick) (fallen []Brick, pt1Solution int, supportSets map[int]map[int]bool) {
	AssertNoIntersections(bricks_)
	bricks := make([]Brick, len(bricks_))
	copy(bricks, bricks_)

	idxs := make([]int, len(bricks))
	for i := range idxs {
		idxs[i] = i
	}

	//> ┌──────────┐   .─.     .─.
	//> │Col (0,0) │  ( A )──▶( B )
	//> └──────────┘   `─'     `─'
	//> ┌──────────┐           .─.     .─.
	//> │Col (1,0) │          ( B )──▶( C )
	//> └──────────┘           `─'     `─'
	//> ┌──────────┐                   .─.     .─.
	//> │Col (1,1) │                  ( C )──▶( D )
	//> └──────────┘                   `─'     `─'
	//> ┌──────────┐   .─.     .─.     .─.     .─.
	//> │  Result  │  ( A )──▶( B )──▶( C )──▶( D )
	//> └──────────┘   `─'     `─'     `─'     `─'

	bricksInCol := BricksSortedByColumns(bricks)

	supportSets = map[int]map[int]bool{}

	// Starting at the bottom, move each brick down until it hits another brick
	for brID, br := range bricks {
		// Find index of previous brick in each col
		prevBrickInCol := map[Coord2]int{}
		for _, col := range br.ColsXY() {
			colBricks := bricksInCol[col]
			prevIdx := slices.Index(colBricks, brID) - 1
			if prevIdx > 0 {
				prevBrickInCol[col] = colBricks[prevIdx]
			}
		}

		for col, prevBrickID := range prevBrickInCol {
			fmt.Printf("prev for %v is %v: %v\n", br.Name, col, bricks[prevBrickID].Name)
		}

		// Find highest Z of support
		highestZ := 0
		for _, prevBrickID := range prevBrickInCol {
			support := bricks[prevBrickID]
			if support.B().Z > highestZ {
				highestZ = support.B().Z
			}
		}

		// Find support set
		supportSet := map[int]bool{}
		for _, prevBrickID := range prevBrickInCol {
			support := bricks[prevBrickID]
			if support.B().Z == highestZ {
				supportSet[prevBrickID] = true
			}
		}
		supportSets[brID] = supportSet

		log.Print(fmt.Sprintln("Brick", br.Name, "is supported by", supportSet, "at", highestZ))

		// Move brick down
		if br.Origin.Z > highestZ+1 {
			fmt.Println("moving from ", br.Origin.Z, "to", highestZ+1)
			bricks[brID].Origin.Z = highestZ + 1
		}
	}

	// Find essential bricks, those that are the sole support in any support set
	essentialBricks := map[int]bool{}
	for _, supportSet := range supportSets {
		if len(supportSet) == 1 {
			for supportID := range supportSet {
				essentialBricks[supportID] = true
			}
		}
	}

	for brID := range essentialBricks {
		log.Print("Essential brick:", bricks[brID].Name)
	}

	AssertNoIntersections(bricks)
	// Subtract one for the floor trick
	return bricks, len(bricks) - len(essentialBricks) - 1, supportSets
}

func AssertNoIntersections(bricks []Brick) {
	for i, p := range bricks {
		for j, q := range bricks {
			if i == j {
				continue
			}
			if p.Intersects(q) {
				panic(fmt.Sprintf("Bricks %v and %v intersect, %#v, %#v", p.Name, q.Name, p, q))
			}
		}
	}
}

func BricksSortedByColumns(bricks []Brick) map[Coord2][]int {
	cols := FindXYCOlumns(bricks)
	bricksInCol := map[Coord2][]int{}
	for i, br := range bricks {
		for _, col := range cols {
			if br.ContainsXY(col) {
				bricksInCol[col] = append(bricksInCol[col], i)
			}
		}
	}

	// Sort them by Z
	for _, brickIdxs := range bricksInCol {
		sort.Slice(brickIdxs, func(i, j int) bool {
			br1 := bricks[brickIdxs[i]]
			br2 := bricks[brickIdxs[j]]
			return br1.A().Z < br2.A().Z
		})
	}
	return bricksInCol
}

func FindXYCOlumns(bricks []Brick) []Coord2 {
	cols := map[Coord2]bool{}
	for _, br := range bricks {
		a := br.A()
		b := br.B()
		for x := a.X; x <= b.X; x++ {
			for y := a.Y; y <= b.Y; y++ {
				cols[Coord2{x, y}] = true
			}
		}
	}

	lst := []Coord2{}
	for col := range cols {
		lst = append(lst, col)
	}

	sort.Slice(lst, func(i, j int) bool {
		if lst[i].Y == lst[j].Y {
			return lst[i].X < lst[j].X
		}
		return lst[i].Y < lst[j].Y
	})

	return lst
}

func intToName(i int) string {
	return fmt.Sprintf("%c%c", 'A'+i%26, 'A'+i%26)
	// fst := i / 135
	// lst := i % 135
	// return fmt.Sprintf("%c%c", '!'+fst, '!'+lst)
}

func ReadInput() []Brick {
	bricks := []Brick{}
	id := 0

	for _, line := range aocutil.ReadLines("input.txt") {

		nums := aocutil.GetIntsInString(line)
		brick := Brick{
			Name:   intToName(id),
			Origin: Coord3{nums[0], nums[1], nums[2]},
			Size:   Coord3{nums[3] - nums[0] + 1, nums[4] - nums[1] + 1, nums[5] - nums[2] + 1},
			ID:     id,
		}
		bricks = append(bricks, brick)
		id++
	}

	// Add a brick representing the floor
	bricks = append(bricks, Brick{
		Name:   "💚",
		Origin: Coord3{0, 0, 0},
		Size:   Coord3{10, 10, 1},
		ID:     id,
	})

	slices.SortFunc(bricks, func(p, q Brick) int {
		return cmp.Compare(p.Origin.Z, q.Origin.Z)
	})

	for _, br := range bricks {
		log.Printf("%#v", br)
	}

	return bricks
}
