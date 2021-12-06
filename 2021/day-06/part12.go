package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Population [9]int

func Next(pop0 Population) Population {
	var pop Population
	for days, count := range pop0 {
		if days == 0 {
			pop[6] += count
			pop[8] += count
		} else {
			pop[days-1] += count
		}
	}
	return pop
}

func Total(pop Population) int {
	sum := 0
	for _, count := range pop {
		sum += count
	}
	return sum
}

func main() {
	pop := ReadInput()

	i := 0
	for ; i < 80; i++ {
		pop = Next(pop)
	}
	fmt.Println("Part 1:", Total(pop))

	for ; i < 256; i++ {
		pop = Next(pop)
	}
	fmt.Println("Part 2:", Total(pop))
}

func ReadInput() Population {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var pop Population
	for _, str := range strings.Split(string(buf), ",") {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		pop[n]++
	}
	return pop
}
