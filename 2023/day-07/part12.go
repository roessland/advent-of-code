package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

type Type int

const (
	TypeInvalid = iota
	TypeHighCard
	TypeOnePair
	TypeTwoPairs
	TypeThreeOfAKind
	TypeFullHouse
	TypeFourOfAKind
	TypeFiveOfAKind
)

func main() {
	part1()
}

func part1() {
	fmt.Println()
	s := aocutil.ReadLines("input.txt")
	hands := []Hand{}
	for _, line := range s {
		hands = append(hands, NewHand(line))
	}

	sort.Sort(ByStrength(hands))

	totalWinnings := 0
	for i, h := range hands {
		rank := i + 1
		bid := h.Bid

		totalWinnings += rank * bid
		fmt.Printf("%d: %v\n", i+1, h)
	}
	fmt.Println("Total winnings:", totalWinnings)
}

func NewHand(s string) Hand {
	parts := strings.Split(s, " ")

	var bid int
	if len(parts) > 1 {
		var err error
		bid, err = strconv.Atoi(parts[1])
		if err != nil {
			panic("nah bruh")
		}
	}
	return Hand{
		ImmutableMultiSet: *aocutil.NewImmutableMultiSet[Card]([]Card(parts[0])...),
		Cards:             []Card(parts[0]),
		Bid:               bid,
	}
}

type Card rune

func (c Card) Value() int {
	switch c {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	}
	if '2' <= c && c <= '9' {
		return int(c - '0')
	}
	panic("invalid card: " + string(c))
}

func (c Card) Greater(d Card) bool {
	return c.Value() > d.Value()
}

func NewCardinalities(vs ...int) *aocutil.ImmutableMultiSet[int] {
	return aocutil.NewImmutableMultiSet[int](vs...)
}

func (h Hand) Cardinalities() *aocutil.ImmutableMultiSet[int] {
	return NewCardinalities(h.Multiplicities()...)
}

type Hand struct {
	aocutil.ImmutableMultiSet[Card]
	Cards []Card
	Bid   int
}

func (h Hand) String() string {
	return fmt.Sprintf("%v", string(h.Cards))
}

func (h Hand) Stronger(j Hand) bool {
	hType, jType := h.Type(), j.Type()
	if hType != jType {
		return hType > jType
	}

	if hType != jType {
		return hType > jType
	}

	for i := 0; i < len(h.Cards); i++ {
		if h.Cards[i] == j.Cards[i] {
			continue
		}
		return h.Cards[i].Greater(j.Cards[i])
	}

	panic(fmt.Sprintf("equal hands: %v %v", h, j))
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

func (h Hand) Type() Type {
	if h.IsFiveOfAKind() {
		return TypeFiveOfAKind
	}
	if h.IsFourOfAKind() {
		return TypeFourOfAKind
	}
	if h.IsFullHouse() {
		return TypeFullHouse
	}
	if h.IsThreeOfAKind() {
		return TypeThreeOfAKind
	}
	if h.IsTwoPairs() {
		return TypeTwoPairs
	}
	if h.IsOnePair() {
		return TypeOnePair
	}
	if h.IsHighCard() {
		return TypeHighCard
	}
	panic(fmt.Sprintf("invalid hand %v: no type", h))
}

type ByStrength []Hand

func (hs ByStrength) Len() int {
	return len(hs)
}

func (hs ByStrength) Less(i, j int) bool {
	return hs[j].Stronger(hs[i])
}

func (hs ByStrength) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
