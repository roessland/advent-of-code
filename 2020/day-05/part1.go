package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
func GetSeatID(seat string) int {
	seat = strings.ReplaceAll(seat, "F", "0")
	seat = strings.ReplaceAll(seat, "B", "1")
	seat = strings.ReplaceAll(seat, "L", "0")
	seat = strings.ReplaceAll(seat, "R", "1")
	id, err := strconv.ParseUint(seat, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(id)
}

func main() {
	f, _ := os.Open("input.txt")
	reader := bufio.NewScanner(f)
	max := 0
	for reader.Scan() {
		seat := reader.Text()
		id := GetSeatID(seat)
		if id > max {
			max = id
		}
	}
	fmt.Println(max)
}