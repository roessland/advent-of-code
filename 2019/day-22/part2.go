package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lines []string

type Deck struct {
	NumCards int
	CardPos int
}

func (d Deck) Deal() Deck {
	d.CardPos = d.NumCards - d.CardPos-1
	return d
}

func (d Deck) RevDeal() Deck {
	if d.CardPos < 0 {
		panic("someone fucked up")
	}
	ret := d.Deal()
	if ret.CardPos < 0 {
		panic("someone fucked up")
	}
	return ret
}

func (d Deck) Increment(N int) Deck {
	d.CardPos = (N * d.CardPos) % d.NumCards
	return d
}


func modInv(a, n int) int {
	t, newt := 0, 1
	r, newr := n, a
	for newr != 0 {
		quotient := r / newr
		t, newt = newt, t - quotient*newt
		r, newr = newr, r - quotient*newr
	}
	if r > 1 {
		fmt.Println(a, "not invertible mod", n)
		panic("not invertible")
	}
	if t < 0 {
		t += n
	}
	return t
}

func modMult(a, b, m int) int {
	res := 0
	for b > 0 {
		// This is the remainder from division by two,
		// which disappears during integer division.
		if b & 1 == 1 {
			res = (res + a) % m
		}

		a = (2*a) % m
		b /= 2
	}
	return res
}

func (d Deck) RevIncrement(N int) Deck {
	if d.CardPos < 0 {
		panic("someone fucked up")
	}
	// find p0 such that d.CardPos == (N * p0) % d.NumCards
	// p0 = modinv(N) * d.CardPos % d.NumCards
	Ninv := modInv(N, d.NumCards)
	partial := modMult(Ninv, d.CardPos, d.NumCards)
	//partial = (Ninv*d.CardPos) % d.NumCards
	d.CardPos = partial
	if partial < 0 {
		panic("overflow")
	}
	if d.CardPos < 0 {
		panic("someone fucked up")
	}
	return d
}

func (d Deck) Cut(N int) Deck {
	d.CardPos = (d.CardPos - N + d.NumCards) % d.NumCards
	return d
}


func (d Deck) RevCut(N int) Deck {
	if d.CardPos < 0 {
		panic("someone fucked up")
	}
	d.CardPos = (d.CardPos + N + d.NumCards) % d.NumCards
	if d.CardPos < 0 {
		panic("someone fucked up")
	}
	return d
}

func NewDeck(numCards int, card int) Deck {
	return Deck{
		NumCards: numCards,
		CardPos:  card,
	}
}

// FastIteratedAffineTransform returns the iterated application of A = a*n + n, A^p = A^(p-1)(A(n)) mod m
func FastIteratedAffineTransform(a_, b_, n, p int) int {
	m := 119315717514047

	A := func(x int)int {
		return (modMult(a_, x, m) + b_) % m
	}

	if p == 1 {
		return A(n)
	}

	A2_1 := A(A(1)) // A2_1 = a + b,    2*A2_1 = 2a+2b
	A2_2 := A(A(2)) // A2_2 = 2a + b,    a = A2_1 - b
	b2 := (2*A2_1 - A2_2 + m) % m
	a2 := (A2_1 - b2 + m) % m

	if p & 1 == 1 {
		return A(FastIteratedAffineTransform(a2, b2, n, p/2))
	} else {
		return FastIteratedAffineTransform(a2, b2, n, p/2)
	}
}

func Shuffle(d Deck) Deck {
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		if strings.HasPrefix(line, "cut") {
			parts := strings.Split(line, "cut ")
			N, _ := strconv.Atoi(parts[1])
			d = d.RevCut(N)
		} else if strings.HasPrefix(line, "deal with") {
			parts := strings.Split(line, "deal with increment ")
			N, _ := strconv.Atoi(parts[1])
			d = d.RevIncrement(N)
		} else if strings.HasPrefix(line, "deal into new") {
			d = d.RevDeal()
		}
	}
	return d
}

func main() {
	times := 101741582076661
	numCards := 119315717514047

	// Read input lines
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Assume the shuffle is an affine transformation, nextCardPos = a * cardPos + b.
	// Find a and b.
	s1 := Shuffle(NewDeck(numCards, 1)).CardPos // = a + b
	s2 := Shuffle(NewDeck(numCards, 2)).CardPos // = a*2 + b
	b := (2*s1 - s2 + numCards) % numCards
	a := (s1 - b + numCards) % numCards

	// Apply the iterated affine transformation a crazy amount of times.
	fmt.Println("Part 2:", FastIteratedAffineTransform(a, b, 2020, times))

}