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
	curr_house_santa := House{0, 0}
	curr_house_robot := House{0, 0}
	visits[curr_house_santa]++
	visits[curr_house_robot]++
	buf, _ := ioutil.ReadFile("input.txt")
	for i, dir := range []rune(string(buf)) {
		if i%2 == 0 {
			curr_house_santa = curr_house_santa.next(dir)
			visits[curr_house_santa]++
		} else {
			curr_house_robot = curr_house_robot.next(dir)
			visits[curr_house_robot]++
		}
	}
	fmt.Printf("%v\n", len(visits))
}
