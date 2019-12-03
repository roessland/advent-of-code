package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type TileType rune
const (
	Wall = '#'
	Cavern = '.'
)

type UnitType string
const (
	Elf = "E"
	Goblin = "G"
)

func (ut UnitType) String() string {
	return string(ut)
}

type Unit struct {
	HitPoints int
	AttackPower int
	Alive bool
}

type Tile struct {
	Type TileType
	Unit *Unit
}

type Pos struct {
	X, Y int
}

type Map map[Pos]*Tile

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	board := make(Map)
	for Y, line := range strings.Split(string(buf), "\n") {
		if line == "" {
			break
		}
		for X, r := range line {
			board[Pos{X, Y}] = &Tile{
				Type: TileType(r),
				Unit: nil,
			}
		}
		fmt.Println(Y, line)
	}
}