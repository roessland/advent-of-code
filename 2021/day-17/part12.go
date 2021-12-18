package main

import (
	"fmt"
	. "github.com/roessland/gopkg/mathutil"
)

type TargetArea struct {
	MinX, MaxX, MinY, MaxY int
}

func (area TargetArea) Contains(x, y int) bool {
	return area.ContainsX(x) && area.ContainsY(y)
}

func (area TargetArea) ContainsX(x int) bool {
	return area.MinX <= x && x <= area.MaxX
}

func (area TargetArea) ContainsY(y int) bool {
	return area.MinY <= y && y <= area.MaxY
}

func Simulate(dx, dy int, area TargetArea) (maxHeight int) {
	x, y := 0, 0
	for {
		x += dx
		y += dy
		maxHeight = MaxInt(maxHeight, y)
		dx -= SignInt(x)
		dy -= 1
		if area.Contains(x, y) {
			return maxHeight
		}
		if y < area.MinY {
			return 0
		}
	}
}

func FindDx(area TargetArea) (dx int) {
	for dx = 1; dx <= area.MaxX; dx++ {
		hit := SimulateX(dx, area)
		if hit {
			return dx
		}
	}
	return
}

func SimulateX(dx int, area TargetArea) (hit bool) {
	x := 0
	for {
		x += dx
		dx -= SignInt(x)
		if dx == 0 {
			if area.ContainsX(x) {
				return true
			} else {
				return false
			}
		}
	}
}

func FindDy(area TargetArea) int {
	var lastDy int
	for dy := 0; dy <= -area.MinY; dy++ {
		hit := SimulateY(dy, area)
		if hit {
			lastDy = dy
		}
	}
	return lastDy
}

func SimulateY(dy int, area TargetArea) (hit bool) {
	y := 0
	for {
		y += dy
		dy--
		if area.ContainsY(y) {
			return true
		}
		if y < area.MinY {
			return false
		}
	}
}

func main() {
	targetArea := TargetArea{
		MinX: 137,
		MaxX: 171,
		MinY: -98,
		MaxY: -73,
	}
	_ = targetArea

	//targetArea = TargetArea{
	//	MinX: 20,
	//	MaxX: 30,
	//	MinY: -10,
	//	MaxY: -5,
	//}

	// 3570 too low
	// 3655 too low
	// 4095 too low
	// 4753 ?
	// 4753 !!

	dx1 := FindDx(targetArea)
	dy1 := FindDy(targetArea)
	fmt.Println("dy is", dy1)
	fmt.Println("dx is", dx1)

	for dx := dx1+2; dx <= dx1+3; dx++ {
		for dy := dy1-3; dy <= dy1+3; dy++ {
			fmt.Println(Simulate(dx, dy, targetArea))
		}
	}

}
