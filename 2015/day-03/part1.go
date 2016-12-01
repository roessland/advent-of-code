package main

import "fmt"
import "io/ioutil"

type House struct {
	East, North int
}

func (prev House) next(dir rune) House {
	next := prev
	switch dir {
	case '>':
		next.East++
	case '<':
		next.East--
	case '^':
		next.North++
	case 'v':
		next.North--
	}
	return next
}

func main() {
	visits := make(map[House]int)
	curr_house := House{0, 0}
	visits[curr_house]++
	buf, _ := ioutil.ReadFile("input.txt")
	for _, dir := range []rune(string(buf)) {
		curr_house = curr_house.next(dir)
		visits[curr_house]++
	}
	fmt.Printf("%v\n", len(visits))
}
