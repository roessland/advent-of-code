package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Map []string

func (m Map) TreeAt(i, j int) bool {
	return m[i][j%len(m[0])] == '#'
}

func readInput() Map {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var m Map
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \n\r")
		m = append(m, line)
	}
	return m
}

/*
Right 1, down 1.
Right 3, down 1. (This is the slope you already checked.)
Right 5, down 1.
Right 7, down 1.
Right 1, down 2.
 */

func main() {
	m := readInput()
	di := []int{1, 1, 1, 1, 2}
	dj := []int{1, 3, 5, 7, 1}
	prod := 1
	for t := range []int{0, 1, 2, 3, 4} {
		treesHit := 0
		for i, j := 0, 0; i < len(m); i, j = i+di[t], j+dj[t] {
			if m.TreeAt(i, j) {
				treesHit++
			}
		}
		fmt.Println(treesHit)
		prod = prod * treesHit
	}
	fmt.Println("Prod: ", prod)
}