package day15

import (
	"embed"
	"fmt"
	"strings"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

type Map [][]byte

func (m Map) Print(robotPos Pos) {
	for y, row := range m {
		for x, c := range row {
			p := Pos{Y: y, X: x}
			if p == robotPos {
				fmt.Printf("@")
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Println()
	}
}

type Pos struct {
	Y, X int
}

type Move byte

func (move Move) String() string {
	return string(move)
}

type State struct {
	Map      Map
	Moves    []Move
	Robot    Pos
	CurrMove int
}

func (s *State) Progress() (done bool) {
	if s.CurrMove >= len(s.Moves) {
		return true
	}

	move := s.Moves[s.CurrMove]
	if move == '\n' {
		s.CurrMove++
		return
	}

	nextPos := NextPos(s.Robot, move)
	if s.Bump(nextPos, move) {
		s.Robot = nextPos
	}

	s.CurrMove++
	return false
}

func NextPos(pos Pos, move Move) Pos {
	switch move {
	case '^':
		return Pos{pos.Y - 1, pos.X}
	case 'v':
		return Pos{pos.Y + 1, pos.X}
	case '<':
		return Pos{pos.Y, pos.X - 1}
	case '>':
		return Pos{pos.Y, pos.X + 1}
	}
	panic("Invalid move " + string(move))
}

func (s *State) Bro(pos Pos) *Pos {
	block := s.Map[pos.Y][pos.X]
	if block == '#' {
		return nil
	}
	if block == '.' {
		return nil
	}

	if block == '[' {
		p := Pos{pos.Y, pos.X + 1}
		return &p
	}

	if block == ']' {
		p := Pos{pos.Y, pos.X - 1}
		return &p
	}

	panic("Invalid block " + string(block))
}

func (s *State) BumpTree(tree map[Pos]byte, bumpPos Pos, move Move) map[Pos]byte {
	if _, ok := tree[bumpPos]; ok {
		fmt.Println("Already visited", bumpPos)
		return tree
	}
	tree[bumpPos] = s.Map[bumpPos.Y][bumpPos.X]

	bumpBlock := s.Map[bumpPos.Y][bumpPos.X]
	broPos := s.Bro(bumpPos)

	if bumpBlock == '[' || bumpBlock == ']' {
		s.BumpTree(tree, *broPos, move)
		thisNext := NextPos(bumpPos, move)
		broNext := NextPos(*broPos, move)
		s.BumpTree(tree, thisNext, move)
		s.BumpTree(tree, broNext, move)
	}

	fmt.Println(len(tree))
	return tree
}

func (s *State) Bump(bumpPos Pos, move Move) bool {
	bumpBlock := s.Map[bumpPos.Y][bumpPos.X]

	if bumpBlock == '#' {
		return false
	}

	if bumpBlock == '.' {
		return true
	}

	if bumpBlock == '[' || bumpBlock == ']' {
		fmt.Println("Bumping", bumpPos, string(bumpBlock))
		tree := make(map[Pos]byte)
		s.BumpTree(tree, bumpPos, move)
		movable := true
		for _, block := range tree {
			fmt.Printf("bump tree contains %c\n", block)
			if block == '#' {
				movable = false
				break
			}
		}
		if !movable {
			return false
		}

		for pos, block := range tree {
			if block == '#' {
				panic("meh")
			}
			s.Map[pos.Y][pos.X] = '.'
		}

		for pos, block := range tree {
			if block == '[' || block == ']' {
				nodeNext := NextPos(pos, move)
				s.Map[nodeNext.Y][nodeNext.X] = block
			}
		}

		return true
	}

	return false
}

func ReadInput(inputName string) (Map, []Move, Pos) {
	input := aocutil.FSReadFile(Input, inputName)
	parts := strings.Split(input, "\n\n")
	m := Map{}
	var robot Pos
	for y, line := range strings.Split(parts[0], "\n") {
		row := []byte{}
		for _, c := range line {
			switch c {
			case '#':
				row = append(row, '#', '#')
			case '.':
				row = append(row, '.', '.')
			case 'O':
				row = append(row, '[', ']')
			case '@':
				row = append(row, '@', '.')
			}
		}
		m = append(m, []byte(row))
		for x, c := range row {
			if c == '@' {
				robot = Pos{y, x}
				m[y][x] = '.'
			}
		}
	}
	return m, []Move(parts[1]), robot
}

func (s *State) GPS() int {
	sum := 0
	for y, row := range s.Map {
		for x, c := range row {
			if c == '[' {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func Part12(inputName string) (int, int) {
	m, moves, robotPos := ReadInput(inputName)
	state := &State{
		Map:      m,
		Moves:    moves,
		Robot:    robotPos,
		CurrMove: 0,
	}

	for !state.Progress() {
		// fmt.Println("Robot at", state.Robot)
		// state.Map.Print(state.Robot)

		if state.CurrMove >= len(state.Moves) {
			break
		}
		// fmt.Println(state.Moves[state.CurrMove])
		// time.Sleep(200 * time.Millisecond)
	}

	return state.GPS(), 0
}
