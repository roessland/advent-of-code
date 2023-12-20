package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

func main() {
	part1()
	part2()
}

func part1() {
	input := aocutil.ReadFile("input.txt")
	sum := 0
	for _, str := range strings.Split(input, ",") {
		hash := 0
		for _, c := range str {
			hash += int(c)
			hash *= 17
			hash %= 256
		}
		sum += hash
	}
	fmt.Println("Sum:", sum)
}

func part2() {
	input := aocutil.ReadFile("input-ex1.txt")
	boxes := map[byte][]string{}
	sum := 0
	for _, str := range strings.Split(input, ",") {
		hash := HASH(str)
		// parse either ab-55 or ab=55 using regex
		re := regexp.MustCompile(`(\w{2})([-=])(\d+)`)
		re.FindAllStringSubmatch(str, -1)
		fmt.Println(hash)
	}
	fmt.Println("Sum:", sum)
}

func HASH(str string) byte {
	hash := 0
	for _, c := range str {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return byte(hash)
}
