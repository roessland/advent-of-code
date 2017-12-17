package main

import "fmt"

func main() {
	step := 335
	ans := -1
	for i, pos, size := 1, 0, 1; i <= 50000000; i, size, pos = i+1, size+1, pos+1 {
		if pos+step >= size {
			pos = (pos + step) % size
		} else {
			pos = pos + step
		}
		if pos == 0 {
			ans = i
		}
	}
	fmt.Println(ans)
}
