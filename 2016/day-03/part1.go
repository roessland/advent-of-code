package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Possible(a, b, c int) bool {
	return a+b > c && b+c > a && c+a > b
}

func main() {
	howmany := 0
	buf, _ := ioutil.ReadFile("input1.txt")
	for _, line := range strings.Split(strings.TrimSpace(string(buf)), "\n") {
		numsStr := strings.Fields(line)
		a, _ := strconv.Atoi(numsStr[0])
		b, _ := strconv.Atoi(numsStr[1])
		c, _ := strconv.Atoi(numsStr[2])
		if Possible(a, b, c) {
			howmany += 1
		}
	}
	fmt.Println(howmany)
}
