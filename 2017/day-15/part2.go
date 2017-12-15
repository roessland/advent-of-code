package main

import "fmt"

func main() {
	a := 883
	b := 879
	count := 0
	for i := 0; i < 5000000; i++ {
		a = (a * 16807) % 2147483647
		for a%4 != 0 {
			a = (a * 16807) % 2147483647
		}
		b = (b * 48271) % 2147483647
		for b%8 != 0 {
			b = (b * 48271) % 2147483647
		}
		if (a & 0xffff) == (b & 0xffff) {
			count++
		}
	}
	fmt.Println(count)
}
