package main

import (
	"fmt"

	"github.com/roessland/advent-of-code/2024/day-07/day07"
)

func main() {
	{
		a, b := day07.Part12("input.txt")
		fmt.Println("Part 1:", a)
		fmt.Println("Part 2:", b)
	}

	{
		a, b := day07.Part12("input-ex1.txt")
		fmt.Println("Part 1:", a)
		fmt.Println("Part 2:", b)
	}
}
