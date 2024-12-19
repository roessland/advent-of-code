package main

import (
	"fmt"
	"time"

	"github.com/roessland/advent-of-code/2024/day-19/day19"
)

func main() {
	// {
	// 	a, b := day10.Part12("input.txt")
	// 	fmt.Println("Part 1:", a)
	// 	fmt.Println("Part 2:", b)
	// }
	// {
	// 	a, b := day14.Part12("input-ex1.txt", 11, 7)
	// 	fmt.Println("Part 1:", a)
	// 	fmt.Println("Part 2:", b)
	// }
	//

	t0 := time.Now()
	a, b := day19.Part12("input-ex1.txt")
	// a, b := day18.Part12("input-ex1.txt", 6+1, 12)
	fmt.Println("Part 1:", a)
	fmt.Println("Part 2:", b)
	fmt.Println(time.Since(t0))

	// 33357 too low
}
