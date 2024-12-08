package main

import (
	"fmt"
	"os"

	"github.com/roessland/advent-of-code/2024/aocutil"
	"github.com/roessland/advent-of-code/2024/day-07/day07"
)

var f *os.File

func GetID(result, fst int, rest []int, canCat bool) string {
	return fmt.Sprintf("%d_%d_%v_%v", result, fst, rest, canCat)
}

func CanBeTrue(result, fst int, rest []int, canCat bool, parent string) bool {
	id := GetID(result, fst, rest, canCat)
	color := "#ab4459"

	defer func() {
		fmt.Fprintf(f, `"%s" [label="" border="0" penwidth=0 shape="circle" style=filled fillcolor="%s"]`, id, color)
		fmt.Fprintln(f)
		fmt.Fprintf(f, `"%s"--"%s"`, parent, id)
		fmt.Fprintln(f)
	}()

	if fst > result { // Around 25% speedup
		color = "#441752"
		return false
	}
	if len(rest) == 0 {
		if fst == result {
			color = "#f29f58"
		}
		return fst == result
	}

	return CanBeTrue(result, fst+rest[0], rest[1:], canCat, id) ||
		CanBeTrue(result, fst*rest[0], rest[1:], canCat, id) ||
		(canCat && CanBeTrue(result, Cat(fst, rest[0]), rest[1:], canCat, id))
}

func Cat(a, b int) int {
	return 10*a*Order(b) + b
}

func Order(b int) int {
	n := 1
	for b >= 10 {
		b /= 10
		n *= 10
	}
	return n
}

func main() {
	var err error
	f, err = os.Create("graph.txt")
	if err != nil {
		panic("I'm a sneaky panic!")
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Fprintf(f, "graph {\n")
	// set bg color
	fmt.Fprintf(f, `bgcolor="#1B1833"`)
	fmt.Fprintln(f)

	// set edge color
	fmt.Fprintf(f, `edge [color="#ab4459"]`)
	fmt.Fprintln(f)

	input := aocutil.FSGetIntsInStringLines(day07.Input, "input-ex1.txt")
	sum1, sum2 := 0, 0
	fmt.Fprintf(f, `"%s" [label="Part 1" border="0" penwidth=0 shape="circle" style=filled fillcolor="%s"]`, "1", "#f2af58")
	fmt.Fprintln(f)
	fmt.Fprintf(f, `"%s" [label="Part 2" border="0" penwidth=0 shape="circle" style=filled fillcolor="%s"]`, "2", "#f2af58")
	fmt.Fprintln(f)
	for _, line := range input {
		res, fst, rest := line[0], line[1], line[2:]

		if CanBeTrue(res, fst, rest, false, "1") {
			sum1 += res
		}
		if CanBeTrue(res, fst, rest, true, "2") {
			sum2 += res
		}
	}

	fmt.Fprintf(f, "}\n")
}
