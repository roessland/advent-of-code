package main

import (
	"fmt"
	"github.com/roessland/gopkg/mathutil"
)

// https://adventofcode.com/2021/day/21

type Args [4]int

var cache = map[Args]Pair{}

type Pair struct {
	Fst, Snd int
}

func G(s1, p1, s2, p2 int) Pair {
	if s2 >= 21 {
		return Pair{Fst: 0, Snd: 1}
	}

	cached, ok := cache[Args{s1, p1, s2, p2}]
	if ok {
		return cached
	}

	W1 := 0
	W2 := 0
	for _, outcome := range []Pair{
		{3, 1},
		{4, 3},
		{5, 6},
		{6, 7},
		{7, 6},
		{8, 3},
		{9, 1},
	} {
		roll, count := outcome.Fst, outcome.Snd
		p1Next := ((p1-1)+roll)%10 + 1
		s1Next := s1 + p1Next
		g := G(s2, p2, s1Next, p1Next)
		W1 += count * g.Snd
		W2 += count * g.Fst
	}

	ret := Pair{W1, W2}
	cache[Args{s1, p1, s2, p2}] = ret
	return ret
}

func main() {
	wins := G(0, 4, 0, 3)
	fmt.Println(mathutil.MaxInt(wins.Fst, wins.Snd))
}
