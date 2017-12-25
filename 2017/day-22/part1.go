package main

import "fmt"
import "os"
import "bufio"

type Vec struct {
	X, Y int
}

func (p Vec) Add(v Vec) Vec {
	return Vec{p.X + v.X, p.Y + v.Y}
}

// CCW=1 is counterclockwise, CCW=-1 is clockwise
func Rotate(v Vec, CCW int) Vec {
	if v.Y == -1 {
		return Vec{-CCW, 0}
	}
	if v.Y == 1 {
		return Vec{CCW, 0}
	}
	if v.X == -1 {
		return Vec{0, CCW}
	}
	if v.X == 1 {
		return Vec{0, -CCW}
	}
	return Vec{}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	infected := map[Vec]bool{}
	i, j := 0, 0
	for scanner.Scan() {
		row := scanner.Text()
		i = 0
		for _, c := range row {
			if c == '#' {
				infected[Vec{i, j}] = true
			}
			i++
		}
		j++
	}
	pos := Vec{i / 2, j / 2}
	dir := Vec{0, -1}
	infectionsCaused := 0
	for t := 0; t < 10000; t++ {
		if infected[pos] {
			dir = Rotate(dir, -1)
		} else {
			dir = Rotate(dir, 1)
		}
		if !infected[pos] {
			infectionsCaused++
			infected[pos] = true
		} else {
			infected[pos] = false
		}
		pos = pos.Add(dir)
	}
	fmt.Println(infectionsCaused)
}
