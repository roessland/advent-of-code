package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"

func MinIntSlice(s []int) int {
	min := s[0]
	for _, val := range s {
		if val < min {
			min = val
		}
	}
	return min
}

type Box struct {
	L, W, H int
}

func (b *Box) areas() []int {
	return []int{b.L * b.W, b.W * b.H, b.H * b.L}
}

func (b *Box) area() int {
	return 2*b.L*b.W + 2*b.W*b.H + 2*b.H*b.L
}

func (b *Box) perimeters() []int {
	return []int{2*b.L + 2*b.W, 2*b.W + 2*b.H, 2*b.H + 2*b.L}
}

func (b *Box) volume() int {
	return b.L * b.W * b.H
}

func (b *Box) smallestArea() int {
	return MinIntSlice(b.areas())
}

func (b *Box) smallestPerimeter() int {
	return MinIntSlice(b.perimeters())
}

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(buf), "\n")
	total_area := 0
	total_length := 0
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		dims_str := strings.Split(line, "x")
		d := [3]int{}
		for i, dim_str := range dims_str {
			dim, _ := strconv.Atoi(dim_str)
			d[i] = dim
		}
		box := Box{d[0], d[1], d[2]}
		total_area += box.area() + box.smallestArea()
		total_length += box.smallestPerimeter() + box.volume()
	}
	fmt.Printf("Total area needed: %v\n", total_area)
	fmt.Printf("Total length needed: %v\n", total_length)
}
