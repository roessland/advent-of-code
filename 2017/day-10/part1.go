package main

import "fmt"

//const N int = 5

const N int = 256

//var lengths []int = []int{3, 4, 1, 5}

var lengths []int = []int{70, 66, 255, 2, 48, 0, 54, 48, 80, 141, 244, 254, 160, 108, 1, 41}

type Circ []int

func NewCirc() Circ {
	c := make(Circ, N)
	for i := 0; i < N; i++ {
		c[i] = i
	}
	return c
}

func (c Circ) Get(i int) int {
	return c[(i+N)%N]
}

func (c Circ) Put(i, val int) {
	c[(i+N)%N] = val
}

func (c Circ) Reverse(start, length int) {
	for i, j := start, start+length-1; i < j; i, j = i+1, j-1 {
		tmp := c.Get(i)
		c.Put(i, c.Get(j))
		c.Put(j, tmp)
	}
}

func main() {
	c := NewCirc()
	curr := 0
	for i, length := range lengths {
		c.Reverse(curr, length)
		curr += length + i
	}
	fmt.Println(c[0] * c[1])
}
