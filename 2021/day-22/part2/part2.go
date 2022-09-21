package main

import (
	"bufio"
	"fmt"
	"github.com/roessland/advent-of-code/2021/aocutil"
	"github.com/roessland/gopkg/mathutil"
	"log"
	"os"
	"strings"
	"time"
)

var min = mathutil.MinInt
var max = mathutil.MaxInt

type Cuboid struct {
	X0, X1, Y0, Y1, Z0, Z1 int
}

func Volume(as []Cuboid) int {
	sum := 0
	for _, a := range as {
		sum += a.Volume()
	}
	return sum
}

func Subtract(as []Cuboid, b Cuboid) []Cuboid {
	var out []Cuboid
	for _, a := range as {
		out = append(out, a.Subtract(b)...)
	}
	return out
}

func (a Cuboid) Volume() int {
	if !a.HasVolume() {
		return 0
	}
	return (a.X1 - a.X0) * (a.Y1 - a.Y0) * (a.Z1 - a.Z0)
}

func (a Cuboid) HasVolume() bool {
	return a.X0 < a.X1 &&
		a.Y0 < a.Y1 &&
		a.Z0 < a.Z1
}

func (a Cuboid) Intersect(b Cuboid) Cuboid {
	c := Cuboid{
		X0: max(a.X0, b.X0),
		X1: min(a.X1, b.X1),
		Y0: max(a.Y0, b.Y0),
		Y1: min(a.Y1, b.Y1),
		Z0: max(a.Z0, b.Z0),
		Z1: min(a.Z1, b.Z1),
	}
	if !c.HasVolume() {
		return Cuboid{0, 0, 0, 0, 0, 0}
	}
	return c
}

func (a Cuboid) Subtract(b Cuboid) []Cuboid {
	c := a.Intersect(b)
	if !c.HasVolume() {
		return []Cuboid{a}
	}

	var out []Cuboid
	add := func(cub Cuboid) {
		if !cub.HasVolume() {
			return
		}
		out = append(out, cub)
	}
	add(Cuboid{a.X0, c.X0, a.Y0, a.Y1, a.Z0, a.Z1}) // x0 side
	add(Cuboid{c.X1, a.X1, a.Y0, a.Y1, a.Z0, a.Z1}) // x1 side
	add(Cuboid{c.X0, c.X1, a.Y0, a.Y1, a.Z0, c.Z0}) // y sausage 0
	add(Cuboid{c.X0, c.X1, a.Y0, a.Y1, c.Z1, a.Z1}) // y sausage 1
	add(Cuboid{c.X0, c.X1, c.Y1, a.Y1, c.Z0, c.Z1}) // y face 0
	add(Cuboid{c.X0, c.X1, a.Y0, c.Y0, c.Z0, c.Z1}) // y face 1
	return out
}

type Step struct {
	Cuboid      Cuboid
	Instruction string
}

func ParseLine(line string) Step {
	s := Step{}
	s.Instruction = line[:strings.Index(line, " ")]
	ns := aocutil.GetIntsInString(line)
	s.Cuboid = Cuboid{ns[0], ns[1] + 1, ns[2], ns[3] + 1, ns[4], ns[5] + 1}
	return s
}

func ReadInput() []Step {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	steps := make([]Step, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		steps = append(steps, ParseLine(scanner.Text()))
	}
	return steps
}

func main() {
	t0 := time.Now()
	steps := ReadInput()
	var onCubes []Cuboid
	for _, step := range steps {
		switch step.Instruction {
		case "on":
			newOnCubes := []Cuboid{step.Cuboid}
			for _, onCube := range onCubes {
				newOnCubes = Subtract(newOnCubes, onCube)
			}
			onCubes = append(onCubes, newOnCubes...)
		case "off":
			var onCubesNext []Cuboid
			for _, onCube := range onCubes {
				onCubesNext = append(onCubesNext, onCube.Subtract(step.Cuboid)...)
			}
			onCubes = onCubesNext
		default:
			panic("unknown instruction " + step.Instruction)
		}
	}
	fmt.Println(Volume(onCubes))
	fmt.Println(time.Since(t0))
}
