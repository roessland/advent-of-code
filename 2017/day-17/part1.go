package main

import "fmt"

func main() {
	step := 335
	//step = 3
	buf := make([]int, 2018)
	pos := 0
	size := 1
	for i := 1; i <= 2017; i++ {
		pos = (pos + step) % size
		copy(buf[pos+2:size+1], buf[pos+1:size])
		size++
		buf[(pos+1)%size] = i
		pos++
	}
	fmt.Println(buf[pos+1])
}
