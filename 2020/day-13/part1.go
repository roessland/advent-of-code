package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func EarliestDeparture(earliestTime int, busId int) int {
	t := (earliestTime / busId) * busId
	if t < earliestTime {
		t += busId
	}
	return t
}

func ArgMin(f func(n int) int, ns []int) (int, int) {
	minF := math.MaxInt32
	minN := -0xB4BE5
	for _, n := range ns {
		fn := f(n)
		if fn < minF {
			minF = fn
			minN = n
		}
	}
	return minN, minF
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line1 := scanner.Text()
	scanner.Scan()
	line2 := scanner.Text()

	earliestTime, err := strconv.Atoi(line1)
	if err != nil {
		log.Fatal(err)
	}

	times := []int{}
	timesStr := strings.Split(line2, ",")
	for _, timeStr := range timesStr {
		if timeStr == "x" {
			continue
		}
		time, err := strconv.Atoi(timeStr)
		if err != nil {
			log.Fatal(err)
		}
		times = append(times, time)
	}

	busId, earliestBus := ArgMin(func(busId int) int {
		return EarliestDeparture(earliestTime, busId)
	}, times)

	fmt.Println(busId * (earliestBus - earliestTime))
}
