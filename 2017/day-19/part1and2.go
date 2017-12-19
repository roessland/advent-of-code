package main

import "fmt"
import "unicode"
import "bufio"
import "os"

type Vec struct {
	X, Y int
}

func (v Vec) Add(d Vec) Vec {
	return Vec{v.X + d.X, v.Y + d.Y}
}

func (v Vec) Neg() Vec {
	return Vec{-v.X, -v.Y}
}

func (v Vec) On(diagram [][]rune) rune {
	return diagram[v.Y][v.X]
}

func (v Vec) IsHorizontal() bool {
	return v.Y == 0
}

func (v Vec) IsVertical() bool {
	return v.X == 0
}

func LoadDiagram() [][]rune {
	diagram := [][]rune{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		diagram = append(diagram, []rune(scanner.Text()))
	}
	return diagram
}

func GetStartingPos(diagram [][]rune) Vec {
	x := 0
	for diagram[0][x] != '|' {
		x++
	}
	return Vec{x, 0}
}

func main() {
	diagram := LoadDiagram()
	pos := GetStartingPos(diagram)

	north := Vec{0, -1}
	south := Vec{0, 1}
	west := Vec{-1, 0}
	east := Vec{1, 0}

	numSteps := 0
	dir := south
	for {
		done := false

		// Go straight until corner
		for {
			pos = pos.Add(dir)
			numSteps++
			curr := pos.On(diagram)
			if curr == ' ' {
				done = true
				break
			}
			if unicode.IsLetter(curr) {
				fmt.Printf("%c", curr)
			}
			if curr == '+' {
				break
			}
		}
		if done {
			break
		}
		// Turn in the correct direction
		if dir.IsHorizontal() {
			if pos.Add(north).On(diagram) == '|' {
				dir = north
			} else if pos.Add(south).On(diagram) == '|' {
				dir = south
			} else {
				fmt.Println("we done boys")
			}
		} else {
			if pos.Add(west).On(diagram) == '-' {
				dir = west
			} else if pos.Add(east).On(diagram) == '-' {
				dir = east
			} else {
				fmt.Println("we done boys")
			}
		}
	}
	fmt.Println()
	fmt.Println("Number of steps:", numSteps)
}
