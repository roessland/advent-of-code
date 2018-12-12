package main

import "fmt"

const SerialNumber int = 1133

type Cell struct {
	X, Y, RackId, PowerLevel int
}

func NewCell(X, Y int) Cell {
	c := Cell{}
	c.X = X
	c.Y = Y
	c.RackId = c.X + 10
	c.PowerLevel = c.RackId*c.Y + SerialNumber
	c.PowerLevel = c.PowerLevel * c.RackId
	c.PowerLevel = (c.PowerLevel % 1000) / 100
	c.PowerLevel = c.PowerLevel - 5
	return c
}

func main() {
	grid := make([][]Cell, 300)
	for y := range grid {
		grid[y] = make([]Cell, 300)
		for x := range grid[y] {
			grid[y][x] = NewCell(x+1, y+1)
		}
	}

	maxLevel := 0
	maxX := -1
	maxY := -1
	maxSize := -1
	for size := 1; size <= 60; size++ {
		fmt.Println(size)
		for y := 0; y < len(grid)-size; y++ {
			for x := 0; x < len(grid[y])-size; x++ {
				X, Y := x+1, y+1
				sum := 0
				for j := 0; j < size; j++ {
					for i := 0; i < size; i++ {
						sum += grid[y+j][x+i].PowerLevel
					}
				}
				if sum > maxLevel {
					maxLevel = sum
					maxX = X
					maxY = Y
					maxSize = size
				}
			}
		}
	}
	fmt.Printf("%d,%d,%d\n", maxX, maxY, maxSize)

}
