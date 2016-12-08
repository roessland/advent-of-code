package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Screen struct {
	NumRows int
	NumCols int
	M       [][]uint8
	buf     []uint8
}

func NewScreen(numRows, numCols int) Screen {
	s := Screen{
		numRows,
		numCols,
		make([][]uint8, numRows),
		make([]uint8, Max(numRows, numCols)),
	}
	return s
}

func (s Screen) String() string {

}

func (s Screen) Rect(numRows, numCols int) {
	for j := 0; i < numRows; j++ {
		for i := 0; i < numCols; i++ {
			s.M[j][i] = 1
		}
	}
}

func (s Screen) RotateRow(row, by int) {

}

func (s Screen) RotateCol(col, by int) {

}

func main() {
	s := NewScreen(3, 7)

	fmt.Println("vim-go")
	_ = s
}
