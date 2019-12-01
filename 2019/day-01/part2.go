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
	masses := []int{}
	for scanner.Scan() {
		massStr := scanner.Text()
		if massStr == "" {
			continue
		}
		mass, err := strconv.Atoi(massStr)
		if err != nil {
			log.Fatal(err)
		}
		masses = append(masses, mass)
	}

	fuelRequired := 0
	for _, mass := range masses {
		fuelRequired += FuelRequired(mass)
	}
	fmt.Println(fuelRequired)
}

func FuelRequired(mass int) int {
	fuel := mass/3 - 2
	if fuel < 0 {
		return 0
	}
	return fuel + FuelRequired(fuel)
}
