package main

import (
	"bufio"
	"fmt"
	"github.com/roessland/advent-of-code/2021/aocutil"
	"github.com/roessland/gopkg/mathutil"
	"log"
	"os"
	"strings"
)

type Step struct {
	Xmin, Xmax  int
	Ymin, Ymax  int
	Zmin, Zmax  int
	Instruction string
}

func ParseLine(line string) Step {
	s := Step{}
	s.Instruction = line[:strings.Index(line, " ")]
	ns := aocutil.GetIntsInString(line)
	s.Xmin, s.Xmax, s.Ymin, s.Ymax, s.Zmin, s.Zmax = ns[0], ns[1], ns[2], ns[3], ns[4], ns[5]
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
	grid := map[[3]int]bool{}
	steps := ReadInput()
	for _, step := range steps {
		xmin := mathutil.MaxInt(step.Xmin, -50)
		xmax := mathutil.MinInt(step.Xmax, 50)
		ymin := mathutil.MaxInt(step.Ymin, -50)
		ymax := mathutil.MinInt(step.Ymax, 50)
		zmin := mathutil.MaxInt(step.Zmin, -50)
		zmax := mathutil.MinInt(step.Zmax, 50)
		for x := xmin; x <= xmax; x++ {
			for y := ymin; y <= ymax; y++ {
				for z := zmin; z <= zmax; z++ {
					switch step.Instruction {
					case "on":
						grid[[3]int{x, y, z}] = true
					case "off":
						grid[[3]int{x, y, z}] = false
					default:
						panic("unknown instruction " + step.Instruction)
					}
				}
			}
		}
	}

	numOn := 0
	for x := -50; x <= 50; x++ {
		for y := -50; y <= 50; y++ {
			for z := -50; z <= 50; z++ {
				if grid[[3]int{x, y, z}] {
					numOn++
				}
			}
		}
	}
	fmt.Println(numOn)
}
