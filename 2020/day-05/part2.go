package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	ids := []int{}
	for reader.Scan() {
		seat := reader.Text()
		id := GetSeatID(seat)
		ids = append(ids, id)
	}
	sort.Ints(ids)
	for i := 1; i < len(ids)-1; i++ {
		if ids[i-1] == ids[i]-1 && ids[i+1] == ids[i]+2 {
			fmt.Println(ids[i]+1)
			break
		}
	}
}