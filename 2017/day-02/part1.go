package main

import "os"
import "fmt"
import "strconv"
import "encoding/csv"
import "log"

type Mat [][]int

func NewMat(input [][]string) Mat {
	mat := [][]int{}
	for _, line := range input {
		nums := make([]int, len(line))
		for i, s := range line {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			nums[i] = num
		}
		mat = append(mat, nums)
	}
	return mat
}

func (m Mat) NumRows() int {
	return len(m)
}

func (m Mat) NumCols() int {
	return len(m[0])
}

func (m Mat) GetRows() [][]int {
	return [][]int(m)
}

func MinMax(xs []int) (int, int) {
	min, max := 123456789, -123456789
	for _, x := range xs {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return min, max
}

func (m Mat) Checksum() int {
	sum := 0
	for _, row := range m.GetRows() {
		min, max := MinMax(row)
		sum += max - min
	}
	return sum
}

func main() {
	fmt.Println("vim-go")
	r := csv.NewReader(os.Stdin)
	r.Comma = ' '
	r.TrimLeadingSpace = true
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	mat := NewMat(records)
	fmt.Println("Checksum: ", mat.Checksum())
}
