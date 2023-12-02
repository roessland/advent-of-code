package main

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"io"
	"log"
	"strings"
)

//go:embed input.txt input-ex1.txt
var inputFiles embed.FS

func main() {
	part1()
}

func part1() {
	state := readInput()
	fmt.Println(state)
}

type Value struct {
	Left, Right *Value
	Value       int
}

func readInput() (pairs []*Value) {
	f, err := inputFiles.Open("input-ex1.txt")
	if err != nil {
		log.Fatal(err)
	}
	lineScanner := bufio.NewScanner(f)

	pairs = make([]*Value, 0)
	for {
		fmt.Println("loopin")
		left, err := readValue(lineScanner)
		if err == io.EOF {
			fmt.Println("eof")
			break
		} else if err != nil {
			log.Fatal(err)
		}

		right, err := readValue(lineScanner)
		if err != nil {
			log.Fatal(err)
		}

		pairs = append(pairs, &Value{Left: left, Right: right})

		lineScanner.Scan() // skip blank line, or to end
	}
	return pairs
}

func tokenSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return 0, nil, io.EOF
	}
	if strings.ContainsRune("[,]", rune(data[0])) {
		return 1, data[:1], nil
	}
	var digits int
	for len(data) > digits && '0' <= data[digits] && data[digits] <= '9' {
		digits++
	}
	return digits, data[:digits], nil
}

func readValue(scanner *bufio.Scanner) (*Value, error) {
	scanner.Scan()
	if scanner.Text() == "" {
		return nil, io.EOF
	}
	tokenScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
	tokenScanner.Split(tokenSplitter)
	for tokenScanner.Scan() {
		fmt.Print(tokenScanner.Text())
	}
	fmt.Println()
	// 0          1   2 3   4 5
	// Operation: new = old * 19
	return nil, scanner.Err()
}
