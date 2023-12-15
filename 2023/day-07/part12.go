package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

func main() {
	part1()
}

func part1() {
	s := aocutil.ReadLines("input-ex1.txt")
	hands := []Hand{}
	bids := []int{}
	for _, line := range s {
		parts := strings.Split(line, " ")
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("nah bruh")
		}
		bids = append(bids, bid)
		hands = append(hands, parseHand(parts[0]))
	}

	fmt.Println(hands)
}

func parseHand(s string) Hand {
	return Hand{
		ImmutableMultiSet: *aocutil.NewImmutableMultiSet[Card]([]Card(s)...),
		Fst:               Card(s[0]),
	}
}

type Card rune

func NewCardinalities(vs ...int) *aocutil.ImmutableMultiSet[int] {
	return aocutil.NewImmutableMultiSet[int](vs...)
}

func (h Hand) Cardinalities() *aocutil.ImmutableMultiSet[int] {
	return NewCardinalities(h.Multiplicities()...)
}

type Hand struct {
	aocutil.ImmutableMultiSet[Card]
	Fst Card
}

func (h Hand) Value() int {
	if h.IsFiveOfAKind() {
		return 900
	}
	if h.IsFourOfAKind() {
		return 800
	}
	if h.IsFullHouse() {
		return 700
	}
	if h.IsThreeOfAKind() {
		return 600
	}
	if h.IsTwoPairs() {
		return 500
	}
	if h.IsOnePair() {
		return 400
	}
	if h.IsHighCard() {
		return 300
	}
	panic("nah bruh")
}

func (h Hand) IsFiveOfAKind() bool {
	return h.Cardinalities().Equals(NewCardinalities(5))
}

func (h Hand) IsFourOfAKind() bool {
	return h.Cardinalities().Equals(NewCardinalities(4, 1))
}

func (h Hand) IsFullHouse() bool {
	return h.Cardinalities().Equals(NewCardinalities(3, 2))
}

func (h Hand) IsThreeOfAKind() bool {
	return h.Cardinalities().Equals(NewCardinalities(3, 1, 1))
}

func (h Hand) IsTwoPairs() bool {
	return h.Cardinalities().Equals(NewCardinalities(2, 2, 1))
}

func (h Hand) IsOnePair() bool {
	return h.Cardinalities().Equals(NewCardinalities(2, 1, 1, 1))
}

func (h Hand) IsHighCard() bool {
	return h.Cardinalities().Equals(NewCardinalities(1, 1, 1, 1, 1))
}
