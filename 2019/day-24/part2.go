package main

import (
	"bufio"
	"fmt"
	"os"
)

const MaxLevel = 500
const StartLevel = 250
const N = 5

// Use arrays to avoid nested allocations.
type State [MaxLevel][N][N]byte

// Count number of neighbors for each location.
func (s State) Count() State {
	var counts State
	for level := 1; level < MaxLevel-1; level++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if s[level][i][j] != '#' {
					continue
				}
				// Skip the center, need more advanced logic.
				if i == 2 && j == 2 {
					continue
				}
				// Going out to the left
				if i > 0 {
					counts[level][i-1][j]++
				} else {
					counts[level-1][1][2]++
				}
				// Going out to the right
				if i < N-1 {
					counts[level][i+1][j]++
				} else {
					counts[level-1][3][2]++
				}
				// Going out upwards
				if j > 0 {
					counts[level][i][j-1]++
				} else {
					counts[level-1][2][1]++
				}
				// Going out downwards
				if j < N-1 {
					counts[level][i][j+1]++
				} else {
					counts[level-1][2][3]++
				}
				// Going in downwards
				if i == 1 && j == 2 {
					counts[level+1][0][0]++
					counts[level+1][0][1]++
					counts[level+1][0][2]++
					counts[level+1][0][3]++
					counts[level+1][0][4]++
				}
				// Going in upwards
				if i == 3 && j == 2 {
					counts[level+1][N-1][0]++
					counts[level+1][N-1][1]++
					counts[level+1][N-1][2]++
					counts[level+1][N-1][3]++
					counts[level+1][N-1][4]++
				}
				// Going in to the right
				if i == 2 && j == 1 {
					counts[level+1][0][0]++
					counts[level+1][1][0]++
					counts[level+1][2][0]++
					counts[level+1][3][0]++
					counts[level+1][4][0]++
				}
				// Going in to the left
				if i == 2 && j == 3 {
					counts[level+1][0][N-1]++
					counts[level+1][1][N-1]++
					counts[level+1][2][N-1]++
					counts[level+1][3][N-1]++
					counts[level+1][4][N-1]++
				}
			}
		}
	}
	return counts
}

func (s State) Evolve() State {
	counts := s.Count()
	for level := 1; level < MaxLevel-1; level++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				// Center state is never evolved. It is kept as 0 or '.'
				if i == 2 && j == 2 {
					continue
				}
				if s[level][i][j] == '#' {
					if counts[level][i][j] == 1 {
						continue
					}
					s[level][i][j] = '.'
				} else {
					if counts[level][i][j] == 1 || counts[level][i][j] == 2 {
						s[level][i][j] = '#'
					}
				}
			}
		}
	}
	return s
}

func (s State) Print() {
	for level := MaxLevel/2-5; level <= MaxLevel/2+5; level++ {
		fmt.Println("Level", level)
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if i == 2 && j == 2 {
					fmt.Printf("?")
				} else {
					fmt.Printf("%c", s[level][i][j])
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func (s State) CountAll() int {
	count := 0
	for level := 0; level < MaxLevel; level++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if s[level][i][j] == '#' {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	var state State
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Fill the middle level only.
		copy(state[MaxLevel/2][i][:], line)
		i++
	}

	for i := 0; i < 200; i++ {
		state = state.Evolve()
	}

	fmt.Println("Part 2:", state.CountAll())
}