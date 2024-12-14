package main

import (
	"fmt"
	"time"

	"github.com/roessland/advent-of-code/2024/day-13/day13"
)

func main() {
	fmt.Println("Hello, World!")
	// {
	// 	a, b := day10.Part12("input.txt")
	// 	fmt.Println("Part 1:", a)
	// 	fmt.Println("Part 2:", b)
	// }

	t0 := time.Now()
	a, b := day13.Part12("input.txt")
	fmt.Println("Part 1:", a)
	fmt.Println("Part 2:", b)
	fmt.Println(time.Since(t0))

	// 33357 too low
}
