package main

import "fmt"

import "github.com/roessland/gopkg/disjointset"
import "github.com/roessland/advent-of-code/2017/day-14/knothash"

func GenerateGrid(input string) [][]bool {
	grid := make([][]bool, 128)
	for row, _ := range grid {
		grid[row] = make([]bool, 128)
		hash := knothash.Sum([]byte(fmt.Sprintf("%s-%d", input, row)))
		for col, _ := range hash {
			grid[row][8*col+0] = 1 == ((hash[col] >> 7) & 1)
			grid[row][8*col+1] = 1 == ((hash[col] >> 6) & 1)
			grid[row][8*col+2] = 1 == ((hash[col] >> 5) & 1)
			grid[row][8*col+3] = 1 == ((hash[col] >> 4) & 1)
			grid[row][8*col+4] = 1 == ((hash[col] >> 3) & 1)
			grid[row][8*col+5] = 1 == ((hash[col] >> 2) & 1)
			grid[row][8*col+6] = 1 == ((hash[col] >> 1) & 1)
			grid[row][8*col+7] = 1 == ((hash[col] >> 0) & 1)
		}
	}
	return grid
}

func PrintGrid(grid [][]bool) {
	for row := 0; row < 28; row++ {
		for col := 0; col < 28; col++ {
			if grid[row][col] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func UsedSquares(grid [][]bool) int {
	sum := 0
	for row := 0; row < 128; row++ {
		for col := 0; col < 128; col++ {
			if grid[row][col] {
				sum += 1
			}
		}
	}
	return sum
}

func Used(grid [][]bool, i, j int) bool {
	max := len(grid)
	if i < 0 || j < 0 || i >= max || j >= max {
		return false
	}
	return grid[j][i]
}

func Id(i, j int) int {
	return j*128 + i
}

func ConnectedComponents(grid [][]bool, numUnused int) int {
	ds := disjointset.Make(128 * 128)
	for row := 0; row < 128; row++ {
		for col := 0; col < 128; col++ {
			if !Used(grid, col, row) {
				continue
			}
			this := Id(col, row)
			if Used(grid, col+1, row) {
				ds.Union(this, Id(col+1, row))
			}
			if Used(grid, col-1, row) {
				ds.Union(this, Id(col-1, row))
			}
			if Used(grid, col, row+1) {
				ds.Union(this, Id(col, row+1))
			}
			if Used(grid, col, row-1) {
				ds.Union(this, Id(col, row-1))
			}
		}
	}
	return ds.Count - numUnused
}

func main() {
	input := "jxqlasbh"
	grid := GenerateGrid(input)
	PrintGrid(grid)
	used := UsedSquares(grid)
	fmt.Println("Used squares:", used)
	fmt.Println("Regions:", ConnectedComponents(grid, 128*128-used))
}
