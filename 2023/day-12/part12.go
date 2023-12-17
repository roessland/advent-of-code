package main

import (
	"fmt"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

func main() {
	records := ReadInput()
	part1(records)
	part2(records)
}

func part1(records []Record) {
	solve(records)
}

func part2(records []Record) {
	records = Unfold(records)
	solve(records)
}

func Unfold(records []Record) []Record {
	unfolded := make([]Record, 0)
	for _, r := range records {
		u := Record{}
		for i := 0; i < 5; i++ {
			u.Data = append(u.Data, r.Data...)
			if i < 4 {
				u.Data = append(u.Data, '?')
			}
			u.Groups = append(u.Groups, r.Groups...)
		}
		unfolded = append(unfolded, u)
	}
	return unfolded
}

func solve(records []Record) {
	sum := 0
	for _, record := range records {
		// fmt.Println("Solving", string(record.Data), record.Groups)
		sum += Arrs(record.Data, record.Groups, false)
	}
	fmt.Println(sum)
}

type Record struct {
	Data   []byte
	Groups []int
}

var cache = make(map[string]int)

func Arrs(d []byte, g []int, cont bool) (ret int) {
	cacheKey := fmt.Sprintf("%s %v %v", d, g, cont)
	if cached, ok := cache[cacheKey]; ok {
		return cached
	}

	defer func() { cache[cacheKey] = ret }()

	// No more data
	if len(d) == 0 {
		if len(g) == 0 || len(g) == 1 && g[0] == 0 {
			return 1
		} else {
			return 0
		}
	}

	// Completed a group
	if cont && g[0] == 0 {
		if d[0] == '#' {
			return 0 // ...but it kept going
		}
		return Arrs(d[1:], g[1:], false)
	}

	// Continuing a group
	if cont {
		if d[0] == '#' || d[0] == '?' {
			g[0]--
			defer func() { g[0]++ }()
			return Arrs(d[1:], g, true)
		} else {
			return 0
		}
	}

	// Cannot start a group
	if d[0] == '.' {
		return Arrs(d[1:], g, false)
	}

	// Must start a group
	if d[0] == '#' {
		if len(g) == 0 {
			return 0 // But cannot
		}
		g[0]--
		defer func() { g[0]++ }()
		return Arrs(d[1:], g, true)
	}

	// Can start a group or not
	if d[0] == '?' {
		// But only if we have groups left
		start := 0
		if len(g) > 0 {
			g[0]--
			start = Arrs(d[1:], g, true)
			g[0]++
		}
		skip := Arrs(d[1:], g, false)
		return start + skip
	}
	panic("forgot a case")
}

func (r Record) String() string {
	return fmt.Sprintf("%s %v", r.Data, r.Groups)
}

func ReadInput() []Record {
	records := make([]Record, 0)
	for _, line := range aocutil.ReadLines("input.txt") {
		parts := strings.Split(line, " ")
		groups := aocutil.GetIntsInString(line)
		records = append(records, Record{[]byte(parts[0]), groups})
	}
	return records
}
