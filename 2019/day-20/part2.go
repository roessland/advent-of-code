package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const TileWall = '#'
const TileEmpty = '.'
const TileNothing = ' '



func isLetter(c byte) bool {
	if c < 'A' || 'Z' < c {
		return false
	}
	return true
}

func findDot(tiles []string, i, j int) *Pos {
	var i2, j2 int
	if i > 0 && tiles[i-1][j] == TileEmpty {
		i2, j2 = i-1, j
	} else if i < len(tiles) -1 && tiles[i+1][j] == TileEmpty {
		i2, j2 = i+1, j
	} else if j > 0 && tiles[i][j-1] == TileEmpty {
		i2, j2 = i, j-1
	} else if j < len(tiles[0]) - 1 && tiles[i][j+1] == TileEmpty {
		i2, j2 = i, j+1
	}
	if i2 != 0 && j2 != 0 {
		return &Pos{i2, j2}
	}
	return nil
}

func ReadInput(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Read entire maze into a 2D array
	var tiles []string
	scanner := bufio.NewScanner(f)
	maxLen := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > maxLen {
			maxLen = len(line)
		}
		tiles = append(tiles, line)
	}
	// Make each line have the same length
	for i := range tiles {
		if len(tiles[i]) < maxLen {
			tiles[i] = tiles[i] + strings.Repeat(" ", maxLen+-len(tiles[i]))
		}
	}

	return tiles
}

func main() {
	tiles := ReadInput("input.txt")

	// Make a lazily defined graph, using tiles as data source.
	// Nodes are created when they are accessed, based on IDs.
	var g, start, end = NewDonutGraph(tiles)

	// Find shortest path using Dijkstra's algorithm,
	// stopping early when ZZ is found, otherwise we would
	// loop forever since the lazily defined graph has
	// infinitely many layers.
	shortestPaths := Dijkstra(g, []Node{start}, func  (n Node)bool {
		return n == end
	})

	// We now have the shortest distance to any point which is at the
	// same distance or closer than ZZ.
	fmt.Println("Part 2:", shortestPaths.GetDist(start.ID(), end.ID()))
}
