package main

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"log"
)

//go:embed input.txt input-ex.txt
var inputFiles embed.FS

func main() {
	//part1()
	part2()
}

// part 1:  17 minutes
func part1() {
	heightMap, visibleMap := ReadInput()

	// left
	for i := 0; i < len(heightMap); i++ {
		tallestSoFar := -1
		for j := 0; j < len(heightMap[i]); j++ {
			if heightMap[i][j] > tallestSoFar {
				tallestSoFar = heightMap[i][j]
				visibleMap[i][j]++
			}
		}
	}

	// right
	for i := 0; i < len(heightMap); i++ {
		tallestSoFar := -1
		for j := len(heightMap[i]) - 1; j >= 0; j-- {
			if heightMap[i][j] > tallestSoFar {
				tallestSoFar = heightMap[i][j]
				visibleMap[i][j]++
			}

		}
	}

	// top
	for j := 0; j < len(heightMap[0]); j++ {
		tallestSoFar := -1
		for i := 0; i < len(heightMap); i++ {
			if heightMap[i][j] > tallestSoFar {
				tallestSoFar = heightMap[i][j]
				visibleMap[i][j]++
			}
		}
	}

	// bottom
	for j := 0; j < len(heightMap[0]); j++ {
		tallestSoFar := -1
		for i := len(heightMap) - 1; i >= 0; i-- {
			if heightMap[i][j] > tallestSoFar {
				tallestSoFar = heightMap[i][j]
				visibleMap[i][j]++
			}
		}
	}

	// count
	visibleCount := 0
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if visibleMap[i][j] > 0 {
				visibleCount++
			}
		}
	}
	fmt.Println(visibleCount)
}

func rightScore(heightMap [][]int, I0, J0 int) int {
	houseHeight := heightMap[I0][J0]
	heights := []int{}
	for j := J0 + 1; j < len(heightMap[I0]); j++ {
		heights = append(heights, heightMap[I0][j])
	}
	return visibleTrees(houseHeight, heights)
}

func leftScore(heightMap [][]int, I0, J0 int) int {
	houseHeight := heightMap[I0][J0]
	heights := []int{}
	for j := J0 - 1; j >= 0; j-- {
		heights = append(heights, heightMap[I0][j])
	}
	return visibleTrees(houseHeight, heights)
}

func downScore(heightMap [][]int, I0, J0 int) int {
	houseHeight := heightMap[I0][J0]
	heights := []int{}
	for i := I0 + 1; i < len(heightMap); i++ {
		heights = append(heights, heightMap[i][J0])
	}
	return visibleTrees(houseHeight, heights)
}

func upScore(heightMap [][]int, I0, J0 int) int {
	houseHeight := heightMap[I0][J0]
	heights := []int{}
	for i := I0 - 1; i >= 0; i-- {
		heights = append(heights, heightMap[i][J0])
	}
	return visibleTrees(houseHeight, heights)
}

func visibleTrees(houseHeight int, heights []int) int {
	numVisible := 0
	for _, h := range heights {
		if h < houseHeight {
			numVisible++
		} else {
			numVisible++
			return numVisible
		}
	}
	return numVisible
}

// part 1 and 2: 41 minutes
func part2() {
	heightMap, _ := ReadInput()

	maxScenicScore := -1
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			right := rightScore(heightMap, i, j)
			left := leftScore(heightMap, i, j)
			down := downScore(heightMap, i, j)
			up := upScore(heightMap, i, j)

			scenicScore := right * left * down * up
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	fmt.Println(maxScenicScore)
}

func ReadInput() ([][]int, [][]int) {
	f, err := inputFiles.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var heightMap = [][]int{}
	var visibleMap = [][]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		rowHeight := make([]int, len(line))
		rowVisible := make([]int, len(line))
		for i, c := range line {
			rowHeight[i] = int(c - '0')
		}
		heightMap = append(heightMap, rowHeight)
		visibleMap = append(visibleMap, rowVisible)
	}
	return heightMap, visibleMap
}
