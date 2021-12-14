package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dots map[Pos]int

func (dots Dots) Print() {
	maxX := 0
	maxY := 0
	for pos, _ := range dots {
		if pos.X > maxX {
			maxX = pos.X
		}
		if pos.Y > maxY {
			maxY = pos.Y
		}
	}
	im := make([][]rune, maxY+1)
	for y := 0; y <= maxY; y++ {
		im[y] = make([]rune, maxX+1)
	}
	for pos, _ := range dots {
		im[pos.Y][pos.X] = '█'
	}
	for y := 0; y <= maxY; y++ {
		fmt.Println(string(im[y]))
	}
}

type Pos struct {
	X, Y int
}

type Fold struct {
	Axis rune
	Val  int
}

func main() {
	dots0, folds := ReadInput()
	part1(dots0, folds)
	part2(dots0, folds)
}

func FoldX(srcPos Pos, fold Fold) Pos {
	if srcPos.X < fold.Val {
		return srcPos
	} else {
		return Pos{2*fold.Val - srcPos.X, srcPos.Y}
	}
}

func FoldY(srcPos Pos, fold Fold) Pos {
	if srcPos.Y < fold.Val {
		return srcPos
	} else {
		return Pos{srcPos.X, 2*fold.Val - srcPos.Y}
	}
}

func FoldDot(srcPos Pos, fold Fold) Pos {
	if fold.Axis == 'x' {
		return FoldX(srcPos, fold)
	} else {
		return FoldY(srcPos, fold)
	}
}

func FoldDots(dots0 Dots, fold Fold) Dots {
	dots := make(Dots)
	for srcPos, val := range dots0 {
		var dstPos Pos
		dstPos = FoldDot(srcPos, fold)
		dots[dstPos] += val
	}
	return dots
}

func part1(dots Dots, folds []Fold) {
	dots = FoldDots(dots, folds[0])
	fmt.Println("Part 1:", len(dots))
}

func part2(dots Dots, folds []Fold) {
	for _, fold := range folds {
		dots = FoldDots(dots, fold)
	}
	fmt.Println("Part 2:")
	dots.Print()
}

func ReadInput() (Dots, []Fold) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	dots := make(Dots)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, ",")
		x, err1 := strconv.Atoi(parts[0])
		y, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			log.Fatal("ops")
		}
		dots[Pos{x, y}] = 1
	}

	var folds []Fold
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(strings.Split(line, "fold along ")[1], "=")
		axis := rune(parts[0][0])
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("nah")
		}
		folds = append(folds, Fold{
			Axis: axis,
			Val:  val,
		})
	}

	return dots, folds
}