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

func (m Mat) GetRows() [][]int {
	return [][]int(m)
}

func (m Mat) Checksum() int {
	sum := 0
	for _, row := range m.GetRows() {
		sum += DivisibleResult(row)
	}
	return sum
}

func divmod(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a - quotient*b
	return
}

func DivisibleResult(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			quot, rem := divmod(nums[i], nums[j])
			if rem == 0 {
				return quot
			}
		}
	}
	return -1337
}

func main() {
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
