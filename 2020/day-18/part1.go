package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Op(op rune, a, b int) int {
	switch op {
	case '*':
		return a * b
	case '+':
		return a + b
	default:
		panic("wut")
	}
}

func Eval(rdr *bufio.Reader) int {
	var queue []int
	var op rune
	for {
		if len(queue) == 2 {
			queue[0] = Op(op, queue[0], queue[1])
			queue = queue[0:1]
		}
		r, _, err := rdr.ReadRune()
		if err == io.EOF {
			return queue[0]
		} else if r == '+' || r == '*' {
			op = r
		} else if '0' <= r && r <= '9' {
			queue = append(queue, int(r-'0'))
		} else if r == '(' {
			queue = append(queue, Eval(rdr))
		} else if r == ')' {
			return queue[0]
		} else if r == ' ' {
			// nothing
		} else {
			panic("fuck")
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		sum += Eval(bufio.NewReader(strings.NewReader(scanner.Text())))
	}
	fmt.Println("Part 1:", sum)
}
