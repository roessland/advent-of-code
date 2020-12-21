package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var K int = 0

type Edge []rune

func (e Edge) Equal(o Edge) bool {
	if len(e) != len(o) {
		return false
	}
	for i := 0; i < len(e); i++ {
		if e[i] != o[i] {
			return false
		}
	}
	return true
}

type BaseTile struct {
	N         int
	Id        int
	Data      [][]rune
	Rotations []*Tile
}

func NewBaseTile(id int, data [][]rune) *BaseTile {
	b := BaseTile{}
	b.Id = id
	b.N = len(data)
	b.Data = data
	b.ComputeRotations()
	return &b
}

func (b *BaseTile) GetRotation(corner, eastAxis, southAxis Pos) *Tile {
	t := Tile{Base: b, Corner: corner, EastAxis: eastAxis, SouthAxis: southAxis}

	// Cache the edges
	t.EdgeN = make(Edge, b.N)
	t.EdgeE = make(Edge, b.N)
	t.EdgeS = make(Edge, b.N)
	t.EdgeW = make(Edge, b.N)
	for i := 0; i < b.N; i++ {
		t.EdgeN[i] = t.At(i, 0)
		t.EdgeE[i] = t.At(b.N-1, i)
		t.EdgeS[i] = t.At(i, b.N-1)
		t.EdgeW[i] = t.At(0, i)
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
	Base                       *BaseTile
	Corner                     Pos
	EastAxis, SouthAxis        Pos
	EdgeN, EdgeE, EdgeS, EdgeW Edge
}

func (t Tile) N() int {
	return t.Base.N
}

func (t Tile) At(east, south int) rune {
	dataPos := t.Corner.Mul(t.N() - 1).Add(t.EastAxis.Mul(east)).Add(t.SouthAxis.Mul(south))
	return t.Base.Data[dataPos.I][dataPos.J]
}

func (t Tile) Set(east, south int, r rune) {
	dataPos := t.Corner.Mul(t.N() - 1).Add(t.EastAxis.Mul(east)).Add(t.SouthAxis.Mul(south))
	t.Base.Data[dataPos.I][dataPos.J] = r
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

func (t Tile) Print(removeBorder int) {
	fmt.Println("Base id: ", t.Base.Id)

	for s := removeBorder; s < t.N()-removeBorder; s++ {
		for e := removeBorder; e < t.N()-removeBorder; e++ {
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
	if i > 0 && b[i-1][j] != nil && !tile.EdgeN.Equal(b[i-1][j].EdgeS) {
		return false // North edge
	}
	if i < K-1 && b[i+1][j] != nil && !tile.EdgeS.Equal(b[i+1][j].EdgeN) {
		return false // South edge
	}
	if j > 0 && b[i][j-1] != nil && !tile.EdgeW.Equal(b[i][j-1].EdgeE) {
		return false // West edge
	}
	if j < K-1 && b[i][j+1] != nil && !tile.EdgeE.Equal(b[i][j+1].EdgeW) {
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
	for i = 0; i < K; i++ {
		for j = 0; j < K; j++ {
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
	if !solvedPuzzle.Solved {
		log.Fatal("whoops")
	}

	prod := 1
	prod *= solvedPuzzle.Board[0][0].Base.Id
	prod *= solvedPuzzle.Board[K-1][0].Base.Id
	prod *= solvedPuzzle.Board[0][K-1].Base.Id
	prod *= solvedPuzzle.Board[K-1][K-1].Base.Id

	fmt.Println("Part 1:", prod)
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
			data := make([][]rune, 0)
			for {
				scanner.Scan()
				line := scanner.Text()
				if line == "" {
					break
				}
				data = append(data, []rune(line))
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

func Combine(puzzle Puzzle) *BaseTile {
	N := puzzle.Board[0][0].N()
	M := N-2
	L := len(puzzle.Board) * M
	// Allocate empty data
	data := make([][]rune, L)
	for i := range data {
		data[i] = make([]rune, L)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}
	// Copy each tile, removing border.
	for i0 := 0; i0 < len(puzzle.Board); i0++ {
		for s := 0; s < M; s++ {
			for j0 := 0; j0 < len(puzzle.Board); j0++ {
				for e := 0; e < M; e++ {
					data[i0*(M)+s][j0*(M)+e] = puzzle.Board[i0][j0].At(e+1, s+1)
				}
			}
		}
	}

	t := NewBaseTile(1337, data)
	return t
}

var monster = [][]rune{
	[]rune("                  # "),
	[]rune("#    ##    ##    ###"),
	[]rune(" #  #  #  #  #  #   "),
}

func (t *Tile) MarkMonsters() {
	mS := len(monster)
	mE := len(monster[0])
	for s0 := 0; s0 < t.N()-mS; s0++ {
		NoMonsterFoundTryNextLocation:
		for e0 := 0; e0 < t.N()-mE; e0++ {
			// See if a monster is found at this location. Allows for overlapping sea monsters.
			for s := 0; s < mS; s++ {
				for e := 0; e < mE; e++ {
					val := t.At(e0+e, s0+s)
					if monster[s][e] == '#' && val != '#' && val != 'O' {
						continue NoMonsterFoundTryNextLocation
					}
				}
			}
			// Complete sea monster found! Mark those tiles.
			for s := 0; s < mS; s++ {
				for e := 0; e < mE; e++ {
					if monster[s][e] == '#' {
						t.Set(e0+e, s0+s, 'O')
					}
				}
			}
		}
	}
}


func Part2(puzzle Puzzle) {
	tile := Combine(puzzle)

	for _, t := range tile.Rotations {
		t.MarkMonsters()
	}
	//tile.Rotations[0].Print(0)

	roughness := 0
	for i := range tile.Data {
		for j := range tile.Data[i] {
			if tile.Data[i][j] == '#' {
				roughness++
			}
		}
	}
	fmt.Println("Part 2:", roughness)

}

func main() {
	baseTiles := ReadInput("input.txt")
	solvedPuzzle := Part1(baseTiles)
	Part2(solvedPuzzle)
}
