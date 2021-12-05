package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/roessland/gopkg/mathutil"
)

const N = 1000

type Board [N * N]int

type Line struct {
	X0, Y0, X1, Y1 int
}

type Input []Line

func main() {
	in := ReadInput()
	part12(in)
}

func ReadInput() Input {
	var in Input
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		p0Str := strings.Split(parts[0], ",")
		p1Str := strings.Split(parts[1], ",")
		x0, err1 := strconv.Atoi(p0Str[0])
		y0, err2 := strconv.Atoi(p0Str[1])
		x1, err3 := strconv.Atoi(p1Str[0])
		y1, err4 := strconv.Atoi(p1Str[1])
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			panic("nah")
		}
		in = append(in, Line{
			x0, y0, x1, y1,
		})
	}
	return in
}

func part12(lines Input) {
	board1 := &Board{}
	board2 := &Board{}

	for _, line := range lines {
		if line.X0 == line.X1 || line.Y0 == line.Y1 {
			line.Draw(board1)
		}
		line.Draw(board2)
	}

	fmt.Println("Part 1:", CountOverlaps(board1))
	fmt.Println("Part 2:", CountOverlaps(board2))
}

func (l Line) Draw(b *Board) {
	dx, dy := mathutil.SignInt(l.X1-l.X0), mathutil.SignInt(l.Y1-l.Y0)
	n := mathutil.MaxInt(mathutil.AbsInt(l.X1-l.X0), mathutil.AbsInt(l.Y1-l.Y0))
	for i := 0; i <= n; i++ {
		x, y := l.X0+i*dx, l.Y0+i*dy
		b[y*N+x]++
	}
}

func CountOverlaps(b *Board) int {
	count := 0
	for _, n := range b {
		if n >= 2 {
			count++
		}
	}
	return count
}
