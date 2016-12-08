package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Screen struct {
	NumRows, NumCols int
	M                [][]uint8
	rowBuf, colBuf   []uint8
}

func NewScreen(numRows, numCols int) Screen {
	s := Screen{
		numRows,
		numCols,
		make([][]uint8, numRows),
		make([]uint8, numCols),
		make([]uint8, numRows),
	}
	for j, _ := range s.M {
		s.M[j] = make([]uint8, numCols)
	}
	return s
}

func (s Screen) String() string {
	str := make([]rune, 0)
	for j := 0; j < s.NumRows; j++ {
		for i := 0; i < s.NumCols; i++ {
			if s.M[j][i] == 0 {
				str = append(str, '.')
			} else {
				str = append(str, '#')
			}
		}
		str = append(str, '\n')
	}
	return string(str)
}

func (s Screen) Rect(numRows, numCols int) {
	for j := 0; j < numRows; j++ {
		for i := 0; i < numCols; i++ {
			s.M[j][i] = 1
		}
	}
}

func (s Screen) RotateRow(j, by int) {
	for i := 0; i < s.NumCols; i++ {
		s.rowBuf[i] = s.M[j][i]
	}
	for i := 0; i < s.NumCols; i++ {
		s.M[j][(i+by)%s.NumCols] = s.rowBuf[i]
	}
}

func (s Screen) RotateCol(i, by int) {
	for j := 0; j < s.NumRows; j++ {
		s.colBuf[j] = s.M[j][i]
	}
	for j := 0; j < s.NumRows; j++ {
		s.M[(j+by)%s.NumRows][i] = s.colBuf[j]
	}
}

func (s Screen) LitPixelsCount() int {
	count := 0
	for j := 0; j < s.NumRows; j++ {
		for i := 0; i < s.NumCols; i++ {
			if s.M[j][i] > 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	s := NewScreen(6, 50)

	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if fields[0] == "rect" {
			numStrs := strings.Split(fields[1], "x")
			numCols, _ := strconv.Atoi(numStrs[0])
			numRows, _ := strconv.Atoi(numStrs[1])
			s.Rect(numRows, numCols)
		} else if fields[0] == "rotate" {
			by, _ := strconv.Atoi(fields[4])
			idx, _ := strconv.Atoi(fields[2][2:])
			if fields[1] == "row" {
				s.RotateRow(idx, by)
			} else if fields[1] == "column" {
				s.RotateCol(idx, by)
			}
		}
	}

	fmt.Printf("%v\n", s.String())
	fmt.Println("Lit pixels: ", s.LitPixelsCount())
}
