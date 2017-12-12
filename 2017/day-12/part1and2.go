package main

import "fmt"
import "strings"
import "strconv"
import "log"
import "bufio"
import "os"

import "github.com/roessland/gopkg/disjointset"

func Atoi(s string) int {
	n, err := strconv.Atoi(strings.TrimRight(s, ","))
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func ParseLine(s string) (int, []int) {
	fields := strings.Split(s, " ")
	lhs := Atoi(fields[0])
	rhs := []int{}
	for _, id := range fields[2:] {
		rhs = append(rhs, Atoi(id))
	}
	return lhs, rhs
}

func main() {
	lhs := []int{}
	rhs := [][]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lh, rh := ParseLine(scanner.Text())
		lhs = append(lhs, lh)
		rhs = append(rhs, rh)
	}

	ds := disjointset.Make(len(lhs))
	for i, id0 := range lhs {
		for _, id := range rhs[i] {
			ds.Union(id0, id)
		}
	}

	size := 0
	for i := 0; i < len(lhs); i++ {
		if ds.Connected(i, 0) {
			size++
		}
	}
	fmt.Println("Part 1:", size)
	fmt.Println("Part 2:", ds.Count)
}
