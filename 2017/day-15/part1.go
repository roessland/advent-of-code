package main

import "fmt"
import "strconv"
import "log"

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func main() {
	fmt.Println("vim-go")
	a := 883
	b := 879
	count := 0
	for i := 0; i < 40000000; i++ {
		a = (a * 16807) % 2147483647
		b = (b * 48271) % 2147483647
		if (a & 0xffff) == (b & 0xffff) {
			count++
		}
	}
	fmt.Println(count)
}
