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

func main() {
	m := readInput()
	d := 3
	treesHit := 0
	for i, j := 0, 0; i < len(m); i, j = i+1, j+d {
		if m.TreeAt(i, j) {
			treesHit++
		}
	}
	fmt.Println(treesHit)
}