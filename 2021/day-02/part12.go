package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(cmds []Cmd) {
	horz := 0
	depth := 0
	for _, cmd := range cmds {
		switch cmd.Dir {
		case "forward":
			horz += cmd.N
		case "up":
			depth -= cmd.N
		case "down":
			depth += cmd.N
		}

	}
	fmt.Println(horz*depth)
}

func part2(cmds []Cmd) {
	horz := 0
	depth := 0
	aim := 0
	for _, cmd := range cmds {
		switch cmd.Dir {
		case "forward":
			horz += cmd.N
			depth += aim*cmd.N
		case "up":
			aim -= cmd.N
		case "down":
			aim += cmd.N
		}

	}
	fmt.Println(horz*depth)
}

type Cmd struct {
	Dir string
	N   int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var cmds []Cmd
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		dir := parts[0]
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("nope")
		}
		cmds = append(cmds, Cmd{
			Dir: dir,
			N:   n,
		})
	}

	part1(cmds)
	part2(cmds)
}
