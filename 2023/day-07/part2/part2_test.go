package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"testing"

	"github.com/roessland/advent-of-code/2023/aocutil"
	"github.com/stretchr/testify/assert"
)

func TestCardValues(t *testing.T) {
	assert.EqualValues(t, 1, Card('J').Value())
	assert.EqualValues(t, 2, Card('2').Value())
	assert.EqualValues(t, 3, Card('3').Value())
	assert.EqualValues(t, 4, Card('4').Value())
	assert.EqualValues(t, 5, Card('5').Value())
	assert.EqualValues(t, 6, Card('6').Value())
	assert.EqualValues(t, 7, Card('7').Value())
	assert.EqualValues(t, 8, Card('8').Value())
	assert.EqualValues(t, 9, Card('9').Value())
	assert.EqualValues(t, 10, Card('T').Value())
	assert.EqualValues(t, 12, Card('Q').Value())
	assert.EqualValues(t, 13, Card('K').Value())
	assert.EqualValues(t, 14, Card('A').Value())
}

func h(s string) Hand {
	return NewHand(s)
}

func TestType(t *testing.T) {
	assert.EqualValues(t, TypeFiveOfAKind, h("AAAAA").Type())
	assert.EqualValues(t, TypeFourOfAKind, h("AA8AA").Type())
	assert.EqualValues(t, TypeFullHouse, h("23332").Type())
	assert.EqualValues(t, TypeThreeOfAKind, h("TTT98").Type())
	assert.EqualValues(t, TypeTwoPairs, h("23432").Type())
	assert.EqualValues(t, TypeOnePair, h("A23A4").Type())
	assert.EqualValues(t, TypeHighCard, h("23456").Type())
}

func TestGreater(t *testing.T) {
	orderedHands := []Hand{
		h("32T3K"),
		h("KK677"),
		h("T55J5"),
		h("QQQJA"),
		h("KTJJT"),
	}

	assertHandsOrdered(t, orderedHands)
}

func assertHandsOrdered(t *testing.T, orderedHands []Hand) {
	t.Helper()
	for i := 0; i < len(orderedHands)-1; i++ {
		msg := fmt.Sprintf("expected %v to be stronger than %v", orderedHands[i+1], orderedHands[i])
		assert.True(t, orderedHands[i+1].Stronger(orderedHands[i]), msg)
	}
}

func TestSortByStrength(t *testing.T) {
	orderedHands := []Hand{
		h("32T3K"),
		h("KTJJT"),
		h("KK677"),
		h("T55J5"),
		h("QQQJA"),
	}

	for i := 0; i < 100; i++ {
		rand.Shuffle(len(orderedHands), func(i, j int) {
			orderedHands[i], orderedHands[j] = orderedHands[j], orderedHands[i]
		})

		sort.Sort(ByStrength(orderedHands))
		assertHandsOrdered(t, orderedHands)
	}
}

func TestBestHand(t *testing.T) {
	assert.True(t,
		aocutil.NewImmutableMultiSet(
			[]Card("QQQQA")...).
			Equals(
				BestHand(
					aocutil.NewImmutableMultiSet(
						[]Card("QQQA")...,
					),
					1,
				),
			),
	)
}

func TestNewHand(t *testing.T) {
	assert.EqualValues(t, TypeFourOfAKind, NewHand("QJJQ2").Type())
	assert.EqualValues(t, TypeFiveOfAKind, NewHand("JJJJJ").Type())
	assert.EqualValues(t, TypeFiveOfAKind, NewHand("AAAAA").Type())
	assert.EqualValues(t, TypeFourOfAKind, NewHand("33JKJ").Type())
	assert.EqualValues(t, TypeFullHouse, NewHand("3322J").Type())
}
