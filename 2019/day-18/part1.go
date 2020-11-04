package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadInput(filename string) [][]rune {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)
	out := make([][]rune, 0)
	line := make([]rune, 0)
	for {
		r, _, err := reader.ReadRune()
		if r == '\r' {
			continue
		}
		if err == io.EOF {
			if len(line) > 0 {
				out = append(out, line)
			}
			break
		} else if err != nil {
			log.Fatal(err)
		} else if r == '\n' {
			out = append(out, line)
			line = nil
		} else {
			line = append(line, r)
		}
	}
	return out
}

func PrintMap(tiles [][]rune) {
	for _, line := range tiles {
		fmt.Printf("%s\n", string(line))
	}
}

type Pos struct {
	I, J int
}

type Edge interface {
	Id() int
	Weight() int
}

type Node interface {
	Id() int
	Edges() []Edge
}

type MapEdge struct {
	Id_ int
	To int
	Weight_ int
}

func (me *MapEdge) Id() int {
	return me.Id_
}

func (me *MapEdge) Weight() int {
	return me.Weight_
}

type MapNode struct {
	Id_ int
	Edges_ []int
	Pos Pos
	Tile rune
}

func (mn *MapNode)

func CreateMap(input [][]rune) []Node {
	height := len(input)
	width := len(input[0])
	for i := 0; i < height; i++ {
		fmt.Println("lines has len", len(input[i]))
		for j := 0; j < width; j++ {
			fmt.Printf("%c", input[i][j])
		}
	}
	return nil
}

func (m Map) String() string {
	return "TODO"
}

func main() {
	input := ReadInput("input1.txt")
	PrintMap(input)
	m := CreateMap(input)
	fmt.Println(m.String())
}