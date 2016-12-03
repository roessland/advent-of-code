package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Possible(a, b, c int) int {
	if a+b > c && b+c > a && c+a > b {
		return 1
	} else {
		return 0
	}
}

func GetNumbers(line string) (int, int, int) {
	numsStr := strings.Fields(line)
	a, _ := strconv.Atoi(numsStr[0])
	b, _ := strconv.Atoi(numsStr[1])
	c, _ := strconv.Atoi(numsStr[2])
	return a, b, c
}

func main() {
	howmany := 0
	buf, _ := ioutil.ReadFile("input1.txt")
	var a, b, c, d, e, f, g, h, i int
	for idx, line := range strings.Split(strings.TrimSpace(string(buf)), "\n") {
		if idx%3 == 0 {
			a, b, c = GetNumbers(line)
			continue
		} else if (idx)%3 == 1 {
			d, e, f = GetNumbers(line)
			continue
		} else if (idx)%3 == 2 {
			g, h, i = GetNumbers(line)
		}
		howmany += Possible(a, d, g) + Possible(b, e, h) + Possible(c, f, i)
	}
	fmt.Println(howmany)
}
