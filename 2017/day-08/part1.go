package main

import "fmt"
import "encoding/csv"
import "strconv"
import "log"
import "os"

func ReadInput() [][]string {
	reader := csv.NewReader(os.Stdin)
	reader.Comma = ' '
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func GetCmpOp(s string) func(a, b int) bool {
	switch s {
	case "==":
		return func(a, b int) bool { return a == b }
	case "!=":
		return func(a, b int) bool { return a != b }
	case "<":
		return func(a, b int) bool { return a < b }
	case ">":
		return func(a, b int) bool { return a > b }
	case "<=":
		return func(a, b int) bool { return a <= b }
	case ">=":
		return func(a, b int) bool { return a >= b }
	default:
		log.Fatal("No such op: " + s)
	}
	return nil
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	max := 0
	reg := map[string]int{}
	for _, record := range ReadInput() {
		name := record[0]
		op := record[1]
		rs, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal("rs atoi", err)
		}
		condLS := reg[record[4]]
		condOp := GetCmpOp(record[5])
		condRS, err := strconv.Atoi(record[6])
		if err != nil {
			log.Fatal("condRS atoi", err)
		}
		if condOp(condLS, condRS) {
			if op == "inc" {
				reg[name] += rs
			} else if op == "dec" {
				reg[name] -= rs
			} else {
				log.Fatal("No such op: " + op)
			}
		}
		max = Max(max, reg[name])
	}
	fmt.Println("vim-go")
	maxval := -9999999
	for _, val := range reg {
		if val > maxval {
			maxval = val
		}
	}
	fmt.Println(reg)
	fmt.Println("final max:", maxval)
	fmt.Println("during max:", max)
}
