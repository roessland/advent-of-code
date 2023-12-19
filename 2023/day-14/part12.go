package main

import (
	"bytes"
	"fmt"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

const (
	rotN = 0
	rotW = 1
	rotS = 2
	rotE = 3
)

func main() {
	part1()
	part2()
}

func part1() {
	plat := ReadInput()
	plat.Float()
	fmt.Println(plat)
	fmt.Println(plat.Load())
}

func part2() {
	plat := ReadInput()

	cycle := map[string]int{}

	i := 0
	var cycleStart, cycleEnd int
	maxIts := 1000000000
	for i = 0; i < maxIts; i++ {
		str := plat.String()

		if j, ok := cycle[str]; ok {
			fmt.Println("cycle detected at i=", i, "was equal to i=", j)
			cycleStart = j
			cycleEnd = i
			break
		}
		cycle[plat.String()] = i

		plat.SpinCycle()
	}

	for maxIts-i > cycleEnd-cycleStart {
		i += cycleEnd - cycleStart
	}

	for i < maxIts {
		plat.SpinCycle()
		i++
	}

	fmt.Println("after", i, "iterations")
	fmt.Println(plat.Load())

	fmt.Println("------")
	fmt.Println(plat)

	fmt.Println(plat.Load())
}

func (plat *Platform) SpinCycle() {
	for rot := 0; rot >= -3; rot-- {
		plat.rotation = (rot + 4) % 4
		plat.Float()
	}
}

func (plat *Platform) Float() {
	for {
		movedSomething := false

		for y := 0; y < plat.Height(); y++ {
			for x := 0; x < plat.Width(); x++ {
				thisPos := Pos{y, x}
				abovePos := Pos{y - 1, x}
				if plat.At(thisPos) == 'O' && plat.At(abovePos) == '.' {
					plat.Swap(thisPos, abovePos)
					movedSomething = true
				}
			}
		}

		if !movedSomething {
			break
		}
	}
}

func (plat *Platform) Load() int {
	sum := 0

	defer func(orig int) {
		plat.rotation = orig
	}(plat.rotation)
	plat.rotation = rotN

	for y := 0; y < plat.Height(); y++ {
		for x := 0; x < plat.width; x++ {
			if plat.At(Pos{y, x}) == 'O' {
				sum += (plat.Height() - y)
			}
		}
	}
	return sum
}

type Platform struct {
	data     [][]byte
	rotation int
	width    int
	height   int
}

func (plat *Platform) Height() int {
	if plat.rotation%2 == 0 {
		return plat.height
	}
	return plat.width
}

func (plat *Platform) Width() int {
	if plat.rotation%2 == 0 {
		return plat.width
	}
	return plat.height
}

func (plat *Platform) String() string {
	var buf bytes.Buffer
	for _, line := range plat.data {
		fmt.Fprintln(&buf, string(line))
	}
	return buf.String()
}

type Pos struct {
	Y, X int
}

func (plat *Platform) translate(p1 Pos) Pos {
	switch plat.rotation {
	case 0:
		return p1
	case 1:
		return Pos{Y: p1.X, X: plat.Height() - p1.Y - 1}
	case 2:
		return Pos{X: plat.Width() - p1.X - 1, Y: plat.Height() - p1.Y - 1}
	case 3:
		return Pos{Y: plat.Width() - p1.X - 1, X: p1.Y}
	default:
		panic("invalid rotation")
	}
}

func (plat *Platform) At(p Pos) byte {
	p0 := plat.translate(p)
	if p0.X < 0 || p0.X >= plat.width || p0.Y < 0 || p0.Y >= plat.height {
		return '#'
	}
	return plat.data[p0.Y][p0.X]
}

func (plat *Platform) Swap(a, b Pos) {
	a0, b0 := plat.translate(a), plat.translate(b)
	plat.data[a0.Y][a0.X], plat.data[b0.Y][b0.X] = plat.data[b0.Y][b0.X], plat.data[a0.Y][a0.X]
}

func ReadInput() *Platform {
	data := aocutil.ReadLinesAsBytes("input.txt")
	return &Platform{data: data, width: len(data[0]), height: len(data)}
}
