package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt input-ex*.txt
var inputFiles embed.FS

func main() {
	part1() // 38 mins
	part2() // 51 mins (+13 min), 3.6x leaderboard
}

// part 1:  17 minutes
func part1() {
	moves := ReadInput()
	m := make(map[Vec]int)
	var tailPos, headPos Vec

	m[tailPos]++
	for _, move := range moves {
		// fmt.Println(move)
		headPos.X += move.X
		headPos.Y += move.Y
		tailPos = nextTailPos(tailPos, headPos)
		m[tailPos]++
	}

	fmt.Println(len(m))
}

// part 2:
func part2() {
	moves := ReadInput()
	m := make(map[Vec]int)
	positions := make([]Vec, 10)

	m[positions[9]]++
	for _, move := range moves {
		// fmt.Println(move)
		positions[0].X += move.X
		positions[0].Y += move.Y
		for i := 1; i <= 9; i++ {
			positions[i] = nextTailPos(positions[i], positions[i-1])
		}
		m[positions[9]]++
	}

	fmt.Println(len(m))
}

func nextTailPos(t, h Vec) Vec {
	diff := Vec{h.X - t.X, h.Y - t.Y}
	// a  1  2  3  b
	// 4  0  0  0  5
	// 6  0  0  0  7
	// 8  0  0  0  9
	// c 10 11 12  d
	tailMove, ok := map[Vec]Vec{
		{0, 0}:   {0, 0}, // 0
		{1, 1}:   {0, 0}, // 0
		{0, 1}:   {0, 0}, // 0
		{-1, 1}:  {0, 0}, // 0
		{1, 0}:   {0, 0}, // 0
		{-1, 0}:  {0, 0}, // 0
		{1, -1}:  {0, 0}, // 0
		{0, -1}:  {0, 0}, // 0
		{-1, -1}: {0, 0}, // 0

		{1, 2}:  {1, 1},  // 1
		{0, 2}:  {0, 1},  // 2
		{-1, 2}: {-1, 1}, // 3

		{2, 1}:  {1, 1},  // 4
		{2, 0}:  {1, 0},  // 6
		{2, -1}: {1, -1}, // 8

		{-2, 1}:  {-1, 1},  // 5
		{-2, 0}:  {-1, 0},  // 7
		{-2, -1}: {-1, -1}, // 9

		{1, -2}:  {1, -1},  // 10
		{0, -2}:  {0, -1},  // 11
		{-1, -2}: {-1, -1}, // 12

		{2, 2}:   {1, 1},   // a
		{-2, 2}:  {-1, 1},  // b
		{2, -2}:  {1, -1},  // c
		{-2, -2}: {-1, -1}, // d
	}[diff]
	if !ok {
		log.Fatal("couldn't handle", diff)
	}
	t.X += tailMove.X
	t.Y += tailMove.Y
	return t
}

type Vec struct {
	X, Y int
}

var (
	R Vec = Vec{1, 0}
	L Vec = Vec{-1, 0}
	U Vec = Vec{0, -1}
	D Vec = Vec{0, 1}
)

var dirByName = map[string]Vec{
	"R": R,
	"L": L,
	"U": U,
	"D": D,
}

func ReadInput() []Vec {
	f, err := inputFiles.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var dirs []Vec
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		for num > 0 {
			dirs = append(dirs, dirByName[parts[0]])
			num--
		}
	}
	return dirs
}
