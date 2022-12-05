package main

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"github.com/roessland/advent-of-code/2022/aocutil"
	"github.com/roessland/gopkg/mathutil"
	"log"
)

//go:embed input.txt input-ex.txt
var inputFiles embed.FS

func main() {
	pairs := ReadInput()
	part1(pairs)
	part2(pairs)
}

func part1(pairs []Pair) {
	count := 0
	for _, p := range pairs {
		a, b, c, d := p.A, p.B, p.C, p.D
		if a > c {
			a, b, c, d = c, d, a, b
		}
		if a == c || d <= b {
			count++
		}
	}
	fmt.Println(count)
}

func part2(pairs []Pair) {
	count := 0
	for _, p := range pairs {
		m := mathutil.MaxInt(p.A, p.C)
		M := mathutil.MinInt(p.B, p.D)
		if m <= M {
			count++
		}
	}
	fmt.Println(count)
}

type Pair struct {
	A, B, C, D int
}

func ReadInput() []Pair {
	f, err := inputFiles.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var pairs []Pair
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		p := aocutil.GetIntsInString(line)
		pairs = append(pairs, Pair{
			p[0], p[1], p[2], p[3],
		})
	}
	return pairs
}
