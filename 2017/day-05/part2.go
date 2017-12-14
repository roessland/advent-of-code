package main

import "fmt"
import "encoding/csv"
import "log"
import "os"
import "strconv"

func ReadInput() []int {
	reader := csv.NewReader(os.Stdin)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	a := make([]int, len(lines))
	for i, _ := range lines {
		num, err := strconv.Atoi(lines[i][0])
		if err != nil {
			log.Fatal(err)
		}
		a[i] = num
	}
	return a
}

func main() {
	a := ReadInput()
	ip := 0
	steps := 0
	for {
		prevIp := ip
		prevOffset := a[ip]
		ip += a[ip]
		if prevOffset >= 3 {
			a[prevIp]--
		} else {
			a[prevIp]++
		}
		steps++
		if ip < 0 || len(a) <= ip {
			break
		}
	}
	fmt.Println(steps)
}
