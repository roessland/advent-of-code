package main

import (
	"container/ring"
	"fmt"
	"log"
	"strconv"
)

func PrintRing(r *ring.Ring, player, emph int) {
	fmt.Printf("[%d]", player+1)
	r.Do(func(p interface{}) {
		val := p.(int)
		if val == emph {
			fmt.Printf("(%2d)", val)
		} else {
			fmt.Printf(" %2d ", val)
		}
	})
	fmt.Printf("\n")
}

func main() {
	P := 419
	N := 71052 * 100
	scores := make([]int, P)

	currMarble := ring.New(1)
	currMarble.Value = 0
	//zeroMarble := currMarble // Keep a reference for pretty printing
	//PrintRing(zeroMarble, -1, 0)

	p := 0
	for i := 1; i <= N; i++ {

		newMarble := ring.New(1)
		newMarble.Value = i

		if i%23 == 0 {
			scores[p] += i
			removeMarble := currMarble.Move(-7)
			scores[p] += removeMarble.Value.(int)
			currMarble = removeMarble.Next()
			removeMarble.Prev().Unlink(1)
			//PrintRing(zeroMarble, p, currMarble.Value.(int))
		} else {
			currMarble.Next().Link(newMarble)
			currMarble = newMarble
			//PrintRing(zeroMarble, p, i)
		}

		p = (p + 1) % P
	}

	maxScore := 0
	for _, score := range scores {
		maxScore = Max(score, maxScore)
	}
	fmt.Println(maxScore)
}

func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
