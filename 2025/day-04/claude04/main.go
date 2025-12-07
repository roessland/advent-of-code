package main

import (
	"fmt"
	"os"
	"strings"
)

func parseGrid(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func countAdjacentRolls(grid [][]rune, r, c int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			nr, nc := r+dr, c+dc
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
				count++
			}
		}
	}
	return count
}

func findAccessibleRolls(grid [][]rune) [][2]int {
	rows := len(grid)
	cols := len(grid[0])
	var accessible [][2]int

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '@' {
				count := countAdjacentRolls(grid, r, c)
				if count < 4 {
					accessible = append(accessible, [2]int{r, c})
				}
			}
		}
	}

	return accessible
}

func countAccessibleRolls(input string) int {
	grid := parseGrid(input)
	return len(findAccessibleRolls(grid))
}

func totalRemovableRolls(input string) int {
	grid := parseGrid(input)
	total := 0

	for {
		accessible := findAccessibleRolls(grid)
		if len(accessible) == 0 {
			break
		}
		total += len(accessible)

		// Remove all accessible rolls
		for _, pos := range accessible {
			grid[pos[0]][pos[1]] = '.'
		}
	}

	return total
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	part1 := countAccessibleRolls(string(input))
	part2 := totalRemovableRolls(string(input))

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
