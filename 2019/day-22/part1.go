package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deck struct {
	Cards []int
	StartIdx int
	Incr int
}

func (d Deck) String() string {
	cs := make([]string, len(d.Cards))
	for i := 0; i < len(d.Cards); i++ {
		cs[i] = fmt.Sprintf("%d", d.At(i))
	}
	return strings.Join(cs, " ")
}

func (d Deck) At(i int) int {
	idx := (d.StartIdx+i*d.Incr+10*len(d.Cards)) % len(d.Cards)
	return d.Cards[idx]
}

func (d Deck) Set(i int, to int) {
	idx := (d.StartIdx+i*d.Incr+10*len(d.Cards)) % len(d.Cards)
	d.Cards[idx] = to
}

func (d Deck) Deal() Deck {
	return Deck{
		d.Cards,
		(d.StartIdx-d.Incr+len(d.Cards)) % len(d.Cards),
		-d.Incr,
	}
}

func (d Deck) Increment(N int) Deck {
	c := Deck{
		make([]int, len(d.Cards)),
		0,
		1,
	}
	src := 0
	dst := 0
	for i := 0; i < len(d.Cards); i++ {
		c.Set(dst, d.At(src))
		src += 1
		dst += N
	}
	return c
}

func (d Deck) Cut(N int) Deck {
	d.StartIdx += N*d.Incr
	return d
}

func NewDeck(N int) Deck {
	s := make([]int, N)
	for i := range s {
		s[i] = i
	}
	return Deck{
		s,
		0,
		1,
	}
}

func main() {
	d := NewDeck(10007)
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cut") {
			parts := strings.Split(line, "cut ")
			N, _ := strconv.Atoi(parts[1])
			d = d.Cut(N)
		} else if strings.HasPrefix(line, "deal with") {
			parts := strings.Split(line, "deal with increment ")
			N, _ := strconv.Atoi(parts[1])
			d = d.Increment(N)
		} else if strings.HasPrefix(line, "deal into new") {
			d = d.Deal()
		}
	}

	for i := 0; i < len(d.Cards); i++ {
		if d.At(i) == 2019 {
			fmt.Println("Part 1:", i)
			break
		}
	}

	/*
	{
		d := NewDeck(10)
		d = d.Cut(6)
		d = d.Increment(7)
		d = d.Deal()
		fmt.Println(d)
	}
	{
		d := NewDeck(10)
		d = d.Increment(7)
		d = d.Increment(9)
		d = d.Cut(-2)
		fmt.Println(d)
	}
	{
		d := NewDeck(10)
		d = d.Deal()
		d = d.Cut(-2)
		d = d.Increment(7)
		d = d.Cut(8)
		d = d.Cut(-4)
		d = d.Increment(7)
		d = d.Cut(3)
		d = d.Increment(9)
		d = d.Increment(3)
		d = d.Cut(-1)
		fmt.Println(d)
	}
	*/
}