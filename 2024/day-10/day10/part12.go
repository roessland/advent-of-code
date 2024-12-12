package day10

import (
	"embed"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

type Map = [][]byte

func NewBoolMap(m Map) [][]bool {
	b := make([][]bool, len(m))
	for y := range m {
		b[y] = make([]bool, len(m[y]))
	}
	return b
}

func NewIntMap(m Map) [][]int {
	b := make([][]int, len(m))
	for y := range m {
		b[y] = make([]int, len(m[y]))
	}
	return b
}

func ReadInput(inputName string) Map {
	return aocutil.FSReadLinesAsBytes(Input, inputName)
}

func Union(m1, m2 map[int]struct{}) map[int]struct{} {
	m3 := make(map[int]struct{})
	for k := range m1 {
		m3[k] = struct{}{}
	}
	for k := range m2 {
		m3[k] = struct{}{}
	}
	return m3
}

func getId(x, y, width int) int {
	return y*width + x
}

func Grow(m Map) ([][]map[int]struct{}, [][]int) {
	summitsReachable := make([][]map[int]struct{}, len(m))
	for y := range m {
		summitsReachable[y] = make([]map[int]struct{}, len(m[y]))
		for x := range m[y] {
			summitsReachable[y][x] = make(map[int]struct{})
		}
	}

	paths := NewIntMap(m)

	done := false
	for !done {
		done = true
		for y := range m {
			for x := range m[y] {
				if y == 0 || y == len(m)-1 || x == 0 || x == len(m[0])-1 {
					continue
				}
				id := y*len(m[0]) + x
				c := m[y][x]
				if c == '9' {
					paths[y][x] = 1
					before := len(summitsReachable[y][x])
					summitsReachable[y][x][id] = struct{}{}
					after := len(summitsReachable[y][x])
					if before != after {
						done = false
					}
				}

				if m[y-1][x] == c-1 {
					before := len(summitsReachable[y-1][x])
					summitsReachable[y-1][x] = Union(summitsReachable[y-1][x], summitsReachable[y][x])
					after := len(summitsReachable[y-1][x])
					if before != after {
						done = false
					}
					paths[y-1][x]++

				}
				if m[y+1][x] == c-1 {
					before := len(summitsReachable[y+1][x])
					summitsReachable[y+1][x] = Union(summitsReachable[y+1][x], summitsReachable[y][x])
					after := len(summitsReachable[y+1][x])
					if before != after {
						done = false
					}
				}
				if m[y][x-1] == c-1 {
					before := len(summitsReachable[y][x-1])
					summitsReachable[y][x-1] = Union(summitsReachable[y][x-1], summitsReachable[y][x])
					after := len(summitsReachable[y][x-1])
					if before != after {
						done = false
					}
					paths[y][x-1]++
				}
				if m[y][x+1] == c-1 {
					before := len(summitsReachable[y][x+1])
					summitsReachable[y][x+1] = Union(summitsReachable[y][x+1], summitsReachable[y][x])
					after := len(summitsReachable[y][x+1])
					if before != after {
						done = false
					}
					paths[y][x+1]++
				}
			}
		}
	}
	return summitsReachable, paths
}

var cache [][]int

func Rating(m Map, x, y int) int {
	if m[y][x] == '9' {
		return 1
	}

	if cache[y][x] != -1000 {
		return cache[y][x]
	}

	s := 0

	if m[y-1][x] == m[y][x]+1 {
		s += Rating(m, x, y-1)
	}

	if m[y+1][x] == m[y][x]+1 {
		s += Rating(m, x, y+1)
	}

	if m[y][x-1] == m[y][x]+1 {
		s += Rating(m, x-1, y)
	}

	if m[y][x+1] == m[y][x]+1 {
		s += Rating(m, x+1, y)
	}

	cache[y][x] = s
	return s
}

func Part12(inputName string) (int, int) {
	m := ReadInput(inputName)
	m = aocutil.PadMap(m, 1, '.')
	// aocutil.PrintCharMap(m)
	reachableSummit, _ := Grow(m)
	sum := 0
	for y := range m {
		for x := range m[y] {
			if m[y][x] == '0' {
				sum += len(reachableSummit[y][x])
			}
		}
	}

	cache = NewIntMap(m)
	for y := range m {
		for x := range m[y] {
			cache[y][x] = -1000
		}
	}

	sum2 := 0
	for y := range m {
		for x := range m[y] {
			if y == 0 || y == len(m)-1 || x == 0 || x == len(m[0])-1 {
				// fmt.Print("#")
				continue
			}
			// fmt.Printf("%4d", Rating(m, x, y))
			if m[y][x] == '0' {
				sum2 += Rating(m, x, y)
			}
		}
		// fmt.Println()
	}
	return sum, sum2
}
