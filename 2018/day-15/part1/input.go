package main

import (
	"bufio"
	"io"
	"strings"
)

func ReadInputTile(r rune, Y, X int) *Tile {
	tile := &Tile{}
	switch r {
	case '.':
		tile.Type = Open
	case '#':
		tile.Type = Wall
	case 'G':
		tile.Type = Open
		tile.Unit = &Unit{
			Type:        Goblin,
			HitPoints:   200,
			AttackPower: 3,
			Pos:         Pos{X, Y},
		}
	case 'E':
		tile.Type = Open
		tile.Unit = &Unit{
			Type:        Elf,
			HitPoints:   200,
			AttackPower: 3,
			Pos:         Pos{X, Y},
		}
	}
	return tile
}

func ReadMapRow(rowStr string, Y int) []*Tile {
	row := make([]*Tile, len(rowStr))
	for X, r := range rowStr {
		row[X] = ReadInputTile(r, Y, X)
	}
	return row
}

func ReadMap(f io.Reader) Map {
	board := make(Map, 0)
	scanner := bufio.NewScanner(f)
	Y := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		board = append(board, ReadMapRow(line, Y))
		Y++
	}
	return board
}
