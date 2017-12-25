package main

import "fmt"
import "strings"
import "log"
import "bufio"
import "os"

var rules map[string]string

func EmptyGrid(N int) [][]byte {
	grid := make([][]byte, N)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]byte, N)
	}
	return grid
}

func Rotate(p string, times int) string {
	q := []byte(p)
	for ; times > 0; times-- {
		switch len(q) {
		case 5:
			q[0], q[1], q[3], q[4] = q[1], q[4], q[0], q[3]
		case 11:
			q[0], q[1], q[2], q[4], q[6], q[8], q[9], q[10] = q[2], q[6], q[10], q[1], q[9], q[0], q[4], q[8]
		}
	}
	return string(q)
}

func Flip(p string) string {
	q := []byte(p)
	switch len(q) {
	case 5:
		q[0], q[1], q[3], q[4] = q[3], q[4], q[0], q[1]
	case 11:
		q[0], q[1], q[2], q[8], q[9], q[10] = q[8], q[9], q[10], q[0], q[1], q[2]
	}
	return string(q)
}

func Variations(p string) []string {
	qs := make([]string, 8)
	for i := 0; i < 8; i++ {
		qs = append(qs, p, Flip(p))
		p = Rotate(p, 1)
	}
	return qs
}

func ReadRules() map[string]string {
	rules = map[string]string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " => ")
		for _, variation := range Variations(fields[0]) {
			rules[variation] = fields[1]
		}
	}
	return rules
}

func Get(grid [][]byte, j0, i0, n int) string {
	s := make([]byte, 0, n*n)
	for j := j0; j < j0+n; j++ {
		for i := i0; i < i0+n; i++ {
			s = append(s, grid[j][i])
		}
		s = append(s, '/')
	}
	return string(s[0 : len(s)-1])
}

func Set(grid [][]byte, j0, i0, n int, s string) {
	r := 0
	for j := j0; j < j0+n; j++ {
		for i := i0; i < i0+n; i++ {
			grid[j][i] = s[r]
			r++
		}
		r++
	}
}

func Next(grid [][]byte) [][]byte {
	if len(grid)%2 == 0 {
		N := len(grid) / 2
		ret := EmptyGrid(3 * N)
		for j0 := 0; j0 < N; j0++ {
			for i0 := 0; i0 < N; i0++ {
				in := Get(grid, j0*2, i0*2, 2)
				out, ok := rules[in]
				if !ok {
					log.Fatal("No rule for", in)
				}
				Set(ret, j0*3, i0*3, 3, out)
			}
		}
		return ret
	} else {
		N := len(grid) / 3
		ret := EmptyGrid(4 * N)
		for j0 := 0; j0 < N; j0++ {
			for i0 := 0; i0 < N; i0++ {
				in := Get(grid, j0*3, i0*3, 3)
				out := rules[in]
				out, ok := rules[in]
				if !ok {
					log.Fatal("No rule for", in, " ", len(in))
				}
				Set(ret, j0*4, i0*4, 4, out)
			}
		}
		return ret
	}
}

func Print(grid [][]byte) {
	for _, line := range grid {
		fmt.Printf("%s\n", line)
	}
}

func Count(grid [][]byte) int {
	count := 0
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid); i++ {
			if grid[j][i] == '#' {
				count++
			}
		}
	}
	return count
}

func main() {
	rules = ReadRules()
	grid := [][]byte{
		{'.', '#', '.'},
		{'.', '.', '#'},
		{'#', '#', '#'},
	}
	for i := 0; i < 5; i++ {
		grid = Next(grid)
	}
	fmt.Println("Part 1:", Count(grid))

	grid = [][]byte{
		{'.', '#', '.'},
		{'.', '.', '#'},
		{'#', '#', '#'},
	}
	for i := 0; i < 18; i++ {
		grid = Next(grid)
	}
	fmt.Println("Part 2:", Count(grid))
}
