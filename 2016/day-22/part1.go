package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"regexp"
	"strconv"
)

const numRows int = 28
const numCols int = 32

func Atoi(nums []string) []int {
	n := make([]int, len(nums))
	for i, _ := range n {
		n[i], _ = strconv.Atoi(nums[i])
	}
	return n
}

type Node struct {
	X, Y       int
	Size, Used int
}

type ViablePair struct {
	A Node
	B Node
}

func (n Node) Id() int {
	return n.Y*numCols + n.X
}

func (n Node) Avail() int {
	return n.Size - n.Used
}

func MakeGrid() [][]Node {
	grid := make([][]Node, numRows)
	for i, _ := range grid {
		grid[i] = make([]Node, numCols)
	}
	return grid
}

func Unpack(vals []int) (a, b, c, d int) {
	return vals[0], vals[1], vals[2], vals[3]
}

// FindViablePairs returns a list of pairs (A, B) where the amount of data on A
// (non-zero) can be moved to B (there is space available).
func FindViablePairs(grid [][]Node) []ViablePair {
	pairs := make([]ViablePair, 0)
	for yA, _ := range grid {
		for xA, A := range grid[yA] {
			if A.Used == 0 {
				continue
			}
			for yB, _ := range grid {
				for xB, B := range grid[yB] {
					_, _ = xA, xB
					if A.Id() == B.Id() {
						continue
					}
					if A.Used <= B.Avail() {
						pairs = append(pairs, ViablePair{A, B})
					}
				}
			}
		}
	}
	return pairs
}

func PrintGrid(grid [][]Node) {
	for x, _ := range grid[0] {
		for y, _ := range grid {
			n := grid[y][x]
			if n.Avail() > 70 {
				color.Set(color.FgGreen)
			} else if n.Used > 86 {
				color.Set(color.FgRed)
			} else {
				color.Set(color.FgYellow)
			}
			fmt.Printf("%03d/%03d\t", n.Used, n.Size)
			color.Unset()
		}
		fmt.Println()
	}
}

func main() {
	grid := MakeGrid()

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan()
	re := regexp.MustCompile(`\d+`)
	for scanner.Scan() {
		x, y, size, used := Unpack(Atoi(re.FindAllString(scanner.Text(), -1)))
		grid[y][x] = Node{x, y, size, used}
	}

	_ = fmt.Print
	fmt.Println("Number of viable pairs:", len(FindViablePairs(grid)))
	fmt.Println("Grid: (x=0,y=0) is top left, (x=max,y=0) is bottom left")
	fmt.Println("Answer: steps to top of wall + steps to left edge + steps to bottom corner + 5*(steps to top - 1)")
	PrintGrid(grid)
}
