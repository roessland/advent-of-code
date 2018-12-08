package main

import "fmt"
import "math"
import "strconv"
import "strings"
import "io/ioutil"

type Pos struct {
	X, Y int
}

func Dist(x0, y0, x1, y1 int) int {
	return abs(x1-x0) + abs(y1-y0)
}

func FindClosest(points []Pos, x, y int) int {
	minDist := math.MaxInt32
	minName := -1
	for i, p := range points {
		d := Dist(x, y, p.X, p.Y)
		if d == minDist {
			minName = -1
		} else if d < minDist {
			minDist = d
			minName = i
		}
	}
	return minName
}

func FindAreas(closest [][]int) map[int]int {
	areas := map[int]int{}
	for j := range closest {
		for i := range closest[j] {
			areas[closest[j][i]]++
		}
	}
	return areas
}

func SumDist(x, y int, points []Pos) int {
	sum := 0
	for _, p := range points {
		sum += Dist(x, y, p.X, p.Y)
	}
	return sum
}

func main() {
	// Read input
	buf, _ := ioutil.ReadFile("input.txt")
	points := []Pos{}
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		xy := strings.Split(line, ", ")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		points = append(points, Pos{x, y})
	}

	// Get bounds
	minX := points[0].X
	minY := points[0].Y
	maxX := points[0].X
	maxY := points[0].Y
	for _, p := range points {
		minX = min(minX, p.X)
		minY = min(minY, p.Y)
		maxX = max(maxX, p.X)
		maxY = max(maxY, p.Y)
	}

	// Make 2D map
	closest := make([][]int, maxY-minY+1)
	for y := range closest {
		closest[y] = make([]int, maxX-minX+1)
		for x := range closest[y] {
			closest[y][x] = FindClosest(points, x, y)
		}
	}

	// Find invalid points on the edges
	invalidPoints := map[Pos]bool{}
	for y := minY; y <= maxY; y++ {
		pLeft := closest[y-minY][0]
		if pLeft != -1 {
			invalidPoints[points[pLeft]] = true
		}
		pRight := closest[y-minY][maxX-minX]
		if pRight != -1 {
			invalidPoints[points[pRight]] = true
		}
	}
	for x := minX; x <= maxX; x++ {
		pTop := closest[0][x-minX]
		if pTop != -1 {
			invalidPoints[points[pTop]] = true
		}
		pBottom := closest[maxY-minY][x-minX]
		if pBottom != -1 {
			invalidPoints[points[pBottom]] = true
		}
	}
	areas := FindAreas(closest)
	maxArea := 0
	for i, p := range points {
		if invalidPoints[p] {
			continue
		}
		if areas[i] > maxArea {
			maxArea = areas[i]
		}
	}
	fmt.Println(maxArea)

	pt2 := 0
	padding := 200
	for y := minY - padding; y <= maxY+padding; y++ {
		for x := minX - padding; x <= maxX+padding; x++ {
			if SumDist(x, y, points) < 10000 {
				pt2++
			}
		}
	}
	fmt.Println(pt2)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
