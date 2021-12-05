package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Nums       [25]int
	Marked     [25]bool
	LastMarked int
	HasWon     bool
}

type Input struct {
	Nums   []int
	Boards []Board
}

func part12(in Input) {
	lastWin := -1
	for _, draw := range in.Nums {
		for k := range in.Boards {
			if in.Boards[k].HasWon {
				continue
			}
			for i := range in.Boards[k].Nums {
				if in.Boards[k].Nums[i] == draw {
					in.Boards[k].Marked[i] = true
					in.Boards[k].LastMarked = draw
					if hasWon(in.Boards[k]) {
						if lastWin == -1 {
							fmt.Println("Part 1:", computeScore(in.Boards[k]))
						}
						in.Boards[k].HasWon = true
						lastWin = k
					}
				}
			}
		}
	}

	fmt.Println("Part 2:", computeScore(in.Boards[lastWin]))
}

func computeScore(b Board) int {
	sumUnmarked := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.Marked[i*5+j] {
				sumUnmarked += b.Nums[i*5+j]
			}
		}
	}
	return sumUnmarked * b.LastMarked
}

func hasWon(b Board) bool {
nextRow:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.Marked[i*5+j] {
				continue nextRow
			}
		}
		return true
	}

nextCol:
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if !b.Marked[i*5+j] {
				continue nextCol
			}
		}
		return true
	}

	return false
}

func main() {
	in := ReadInput()
	part12(in)
}

func ReadInput() Input {
	var in Input

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	numsStr := scanner.Text()
	for _, numStr := range strings.Split(numsStr, ",") {
		n, err := strconv.Atoi(numStr)
		if err != nil {
			panic("nope")
		}
		in.Nums = append(in.Nums, n)
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		var numsStr string
		numsStr += scanner.Text() + " "
		scanner.Scan()
		numsStr += scanner.Text() + " "
		scanner.Scan()
		numsStr += scanner.Text() + " "
		scanner.Scan()
		numsStr += scanner.Text() + " "
		scanner.Scan()
		numsStr += scanner.Text() + " "

		var b Board

		for i, numStr := range strings.Fields(numsStr) {
			b.Nums[i], err = strconv.Atoi(numStr)
			if err != nil {
				panic("atoi")
			}
		}

		in.Boards = append(in.Boards, b)
	}
	return in
}
