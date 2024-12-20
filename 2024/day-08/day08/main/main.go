package main

import (
	"fmt"
	"time"

	"github.com/roessland/advent-of-code/2024/day-08/day08"
)

func main() {
	{
		t0 := time.Now()
		a, b := day08.Part12("input.txt")
		fmt.Println("Part 1:", a)
		fmt.Println("Part 2:", b)
		fmt.Println(time.Since(t0))
	}

	{
		a, b := day08.Part12("input-ex1.txt")
		fmt.Println("Part 1:", a)
		fmt.Println("Part 2:", b)
	}
}
