package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/roessland/advent-of-code/2023/aocutil"
	"golang.org/x/exp/slices"
)

type Dir byte

const (
	Right Dir = 'R'
	Up    Dir = 'U'
	Left  Dir = 'L'
	Down  Dir = 'D'
)

func main() {
	input := ReadInput()
	part2(input)
}

type Instruction struct {
	Meters int
	Color  [3]byte
	Dir    Dir
}

type DigPlan []Instruction

func ReadInput() DigPlan {
	var digPlan DigPlan

	lines := aocutil.ReadLines("input.txt")
	re := regexp.MustCompile(`([UDLR]) (\d+) \(#([0-9a-f]{5})([0-4])\)`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		meters, err := strconv.ParseInt(match[3], 16, 32)
		if err != nil {
			panic(err)
		}
		var dir Dir
		dirNo := match[4]
		switch dirNo {
		case "0":
			dir = Right
		case "1":
			dir = Up
		case "2":
			dir = Left
		case "3":
			dir = Down
		default:
			panic("invalid dir")
		}
		digPlan = append(digPlan, Instruction{Dir: dir, Meters: int(meters)})
	}
	return digPlan
}

type Line struct {
	From, To Pos
}

func (l Line) Reversed() Line {
	return Line{l.To, l.From}
}

func (l Line) Dir() Dir {
	if l.From.Y == l.To.Y {
		if l.From.X < l.To.X {
			return Right
		} else {
			return Left
		}
	} else {
		if l.From.Y < l.To.Y {
			return Down
		} else {
			return Up
		}
	}
}

type Map struct {
	Lines  []Line
	Height int
	Width  int
}

func dx(dir Dir) int {
	switch dir {
	case Right:
		return 1
	case Left:
		return -1
	default:
		return 0
	}
}

func dy(dir Dir) int {
	switch dir {
	case Up:
		return -1
	case Down:
		return 1
	default:
		return 0
	}
}

func part2(input DigPlan) {
	m := &Map{}

	digLoop(input, m)

	sizeCW := measureTrenchSize(m.Lines)

	slices.Reverse(m.Lines)
	for i, line := range m.Lines {
		m.Lines[i] = line.Reversed()
	}

	sizeCCW := measureTrenchSize(m.Lines)
	if sizeCW > sizeCCW {
		fmt.Println("CW", sizeCW)
	} else {
		fmt.Println("CCW", sizeCCW)
	}
}

type Pos struct {
	Y, X int
}

func digLoop(input DigPlan, m *Map) {
	y, x := 0, 0
	for _, instr := range input {
		start := Pos{y, x}
		endY := dy(instr.Dir) * instr.Meters
		endX := dx(instr.Dir) * instr.Meters
		end := Pos{y + endY, x + endX}
		y, x = end.Y, end.X
		m.Lines = append(m.Lines, Line{start, end})
	}
}

func findVertices(lines []Line) []Pos {
	var vertices []Pos
	for i := range lines {
		a := lines[i]
		b := lines[(i+1)%len(lines)]
		vertices = append(vertices, findVertex(a, b))
	}
	return vertices
}

type Corner struct {
	A, B Dir
}

func findVertex(a, b Line) Pos {
	aDir, bDir := a.Dir(), b.Dir()
	corner := Corner{aDir, bDir}
	switch corner {
	case Corner{Right, Up}:
		return Pos{a.To.Y + 1, a.To.X + 1}
	case Corner{Right, Down}:
		return Pos{a.To.Y + 1, a.To.X}
	case Corner{Up, Left}:
		return Pos{a.To.Y, a.To.X + 1}
	case Corner{Up, Right}:
		return Pos{a.To.Y + 1, a.To.X + 1}
	case Corner{Left, Up}:
		return Pos{a.To.Y, a.To.X + 1}
	case Corner{Left, Down}:
		return Pos{a.To.Y, a.To.X}
	case Corner{Down, Right}:
		return Pos{a.To.Y + 1, a.To.X}
	case Corner{Down, Left}:
		return Pos{a.To.Y, a.To.X}
	default:
		panic("invalid corner")
	}
}

func measureTrenchSize(lines []Line) int {
	signedArea := 0
	vertices := findVertices(lines)

	for i := range vertices {
		from := vertices[i]
		to := vertices[(i+1)%len(vertices)]
		if from.Y == to.Y {
			goingRight := to.X > from.X
			w := to.X - from.X
			if w < 0 {
				w = -w
			}

			var h int = from.Y
			if goingRight {
				signedArea += w * h
			} else {
				signedArea -= w * h
			}
		}
	}
	if signedArea < 0 {
		return -signedArea
	}
	return signedArea
}
