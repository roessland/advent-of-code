package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const N = 10

var K int = 0

type Edge [N]rune

type BaseTile struct {
	Id        int
	Data      [][]rune
	Rotations []*Tile
}

func NewBaseTile(id int, data [][]rune) *BaseTile {
	b := BaseTile{}
	b.Id = id
	b.Data = data
	b.ComputeRotations()
	return &b
}

func (b *BaseTile) GetRotation(corner, eastAxis, southAxis Pos) *Tile {
	t := Tile{Base: b, Corner: corner, EastAxis: eastAxis, SouthAxis: southAxis}
	// Cache the edges
	for i := 0; i < N; i++ {
		t.N[i] = t.At(i, 0)
		t.E[i] = t.At(N-1, i)
		t.S[i] = t.At(i, N-1)
		t.W[i] = t.At(0, i)
	}
	return &t
}

func (b *BaseTile) ComputeRotations() {
	b.Rotations = []*Tile{
		b.GetRotation(Pos{0, 0}, Pos{0, 1}, Pos{1, 0}),
		b.GetRotation(Pos{1, 0}, Pos{-1, 0}, Pos{0, 1}),
		b.GetRotation(Pos{1, 1}, Pos{0, -1}, Pos{-1, 0}),
		b.GetRotation(Pos{0, 1}, Pos{1, 0}, Pos{0, -1}),
		b.GetRotation(Pos{0, 1}, Pos{0, -1}, Pos{1, 0}),
		b.GetRotation(Pos{1, 1}, Pos{-1, 0}, Pos{0, -1}),
		b.GetRotation(Pos{1, 0}, Pos{0, 1}, Pos{-1, 0}),
		b.GetRotation(Pos{0, 0}, Pos{1, 0}, Pos{0, 1}),
	}
}

type Tile struct {
	Base                *BaseTile
	Corner              Pos
	EastAxis, SouthAxis Pos
	N, E, S, W          Edge
}

func (t Tile) At(east, south int) rune {
	dataPos := t.Corner.Mul(N - 1).Add(t.EastAxis.Mul(east)).Add(t.SouthAxis.Mul(south))
	return t.Base.Data[dataPos.I][dataPos.J]
}

type Pos struct {
	I, J int
}

func (p Pos) Mul(n int) Pos {
	return Pos{p.I * n, p.J * n}
}

func (p Pos) Add(v Pos) Pos {
	return Pos{p.I + v.I, p.J + v.J}
}

func (t Tile) Print() {
	for s := 0; s < N; s++ {
		for e := 0; e < N; e++ {
			fmt.Printf("%c", t.At(e, s))
		}
		fmt.Println()
	}
	fmt.Println()
}

type Board [][]*Tile

func (b Board) Copy() Board {
	c := make(Board, K)
	for i := 0; i < K; i++ {
		c[i] = make([]*Tile, K)
		copy(c[i], b[i])
	}
	return c
}

func (b Board) With(i, j int, tile *Tile) Board {
	c := b.Copy()
	c[i][j] = tile
	return c
}

func (b Board) CanPlace(i, j int, tile *Tile) bool {
	if i > 0 && b[i-1][j] != nil && tile.N != b[i-1][j].S {
		return false // North edge
	}
	if i < K-1 && b[i+1][j] != nil && tile.S != b[i+1][j].N {
		return false // South edge
	}
	if j > 0 && b[i][j-1] != nil && tile.W != b[i][j-1].E {
		return false // West edge
	}
	if j < K-1 && b[i][j+1] != nil && tile.E != b[i][j+1].W {
		return false // East edge
	}
	return true
}

type Puzzle struct {
	Board          Board
	RemainingTiles BaseTileSet
	Solved         bool
}

func NewPuzzle(baseTiles []*BaseTile) Puzzle {
	var puzzle Puzzle
	puzzle.Board = make([][]*Tile, K)
	for i := 0; i < K; i++ {
		puzzle.Board[i] = make([]*Tile, K)
	}
	puzzle.RemainingTiles = make(BaseTileSet)
	for _, b := range baseTiles {
		puzzle.RemainingTiles[b] = struct{}{}
	}
	return puzzle
}

func (puzzle Puzzle) Copy() Puzzle {
	c := Puzzle{}
	c.Board = puzzle.Board.Copy()
	c.RemainingTiles = puzzle.RemainingTiles.Copy()
	return c
}

func (puzzle Puzzle) With(i, j int, t *Tile) Puzzle {
	c := Puzzle{}
	c.Board = puzzle.Board.With(i, j, t)
	c.RemainingTiles = puzzle.RemainingTiles.Without(t.Base)
	return c
}

type BaseTileSet map[*BaseTile]struct{}

func (s BaseTileSet) Copy() BaseTileSet {
	c := BaseTileSet{}
	for b := range s {
		c[b] = struct{}{}
	}
	return c
}

func (s BaseTileSet) Without(baseTile *BaseTile) BaseTileSet {
	c := s.Copy()
	delete(c, baseTile)
	return c
}

func Solve(puzzle Puzzle) Puzzle {
	// If there are no tiles left to place, we are done
	if len(puzzle.RemainingTiles) == 0 {
		puzzle.Solved = true
		return puzzle
	}

	// Otherwise go to first empty tile
	i, j := 0, 0
	Free:
	for i=0; i<K; i++ {
		for j=0; j <K; j++ {
			if puzzle.Board[i][j] == nil {
				break Free
			}
		}
	}

	// Try to place a tile there, recursively
	for b := range puzzle.RemainingTiles {
		for _, t := range b.Rotations {
			if !puzzle.Board.CanPlace(i, j, t) {
				continue
			}
			result := Solve(puzzle.With(i, j, t))
			if result.Solved {
				return result
			}
		}
	}

	// Late return means unsolved puzzle
	return puzzle
}

func Part1(baseTiles []*BaseTile) Puzzle {
	puzzle := NewPuzzle(baseTiles)
	solvedPuzzle := Solve(puzzle)

	prod := 1
	prod *= solvedPuzzle.Board[0][0].Base.Id
	prod *= solvedPuzzle.Board[K-1][0].Base.Id
	prod *= solvedPuzzle.Board[0][K-1].Base.Id
	prod *= solvedPuzzle.Board[K-1][K-1].Base.Id

	fmt.Println(prod)
	return solvedPuzzle
}

func ReadInput(filename string) []*BaseTile {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var baseTiles []*BaseTile
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Tile ") {
			id, err := strconv.Atoi(line[5:9])
			if err != nil {
				log.Fatal(err)
			}
			data := make([][]rune, N)
			for i := 0; i < N; i++ {
				scanner.Scan()
				data[i] = []rune(scanner.Text())
			}
			baseTiles = append(baseTiles, NewBaseTile(id, data))
		}
	}
	K = 1 // global board side length
	for K*K != len(baseTiles) {
		K++
	}
	return baseTiles
}

func main() {
	baseTiles := ReadInput("input.txt")
	solvedPuzzle := Part1(baseTiles)
}
