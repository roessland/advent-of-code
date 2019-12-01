package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	nums := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	fuelRequired := 0
	for _, mass := range nums {
		fuelRequired += FuelRequired(mass)
	}
	fmt.Println(fuelRequired)
}

func FuelRequired(mass int) int {
	return mass/3 - 2
}
