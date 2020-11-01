package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Vm struct {
	Mem          []int
	Ip           int
	Halted       bool
	RelativeBase int
}

func NewVm(mem []int) *Vm {
	vm := Vm{make([]int, 100*len(mem)), 0, false, 0}
	copy(vm.Mem, mem)
	return &vm
}

const (
	PositionMode  int = 0
	ImmediateMode int = 1
	RelativeMode  int = 2
)

type Param struct {
	Val  int
	Mode int
}

type Op struct {
	Code     int
	Params   []Param
	Length   int
	FullCode int
}

func (vm *Vm) ReadOp() Op {
	op := Op{}
	val := vm.Mem[vm.Ip]
	op.FullCode = val
	op.Code = val % 100
	switch op.Code {
	case 1:
		op.Length = 4
	case 2:
		op.Length = 4
	case 3:
		op.Length = 2
	case 4:
		op.Length = 2
	case 5:
		op.Length = 3
	case 6:
		op.Length = 3
	case 7:
		op.Length = 4
	case 8:
		op.Length = 4
	case 9:
		op.Length = 2
	case 99:
		op.Length = 1
	default:
		log.Fatal("ReadOp: unknown opcode", op.Code)
	}
	if op.Length >= 2 {
		op.Params = append(op.Params, Param{Val: vm.Mem[vm.Ip+1], Mode: (val / 100) % 10})
	}
	if op.Length >= 3 {
		op.Params = append(op.Params, Param{Val: vm.Mem[vm.Ip+2], Mode: (val / 1000) % 10})
	}
	if op.Length >= 4 {
		op.Params = append(op.Params, Param{Val: vm.Mem[vm.Ip+3], Mode: (val / 10000) % 10})
	}
	if op.Length >= 5 {
		op.Params = append(op.Params, Param{Val: vm.Mem[vm.Ip+4], Mode: (val / 100000) % 10})
	}
	if op.Length >= 6 {
		log.Fatal("ReadOp: instruction of length", op.Length, "not supported")
	}
	return op
}

func (vm *Vm) GetVal(val int, mode int) int {
	if mode == PositionMode {
		return vm.Mem[val]
	} else if mode == ImmediateMode {
		return val
	} else if mode == RelativeMode {
		return vm.Mem[vm.RelativeBase+val]
	} else {
		log.Fatal("GetVal: unknown position mode", mode)
		return -1337
	}
}

func (vm *Vm) SetVal(pos int, mode int, val int) {
	if mode == PositionMode {
		vm.Mem[pos] = val
	} else if mode == ImmediateMode {
		panic("Impossible to set a value in immediate mode!")
	} else if mode == RelativeMode {
		vm.Mem[vm.RelativeBase+pos] = val
	} else {
		log.Fatal("GetVal: unknown position mode", mode)
	}
}

func (vm *Vm) Run(getInput func() int, sendOutput func(int)) {
	for !vm.Halted {
		//fmt.Println("MEM", vm.Mem)
		//fmt.Println("IP", vm.Ip)
		op := vm.ReadOp()
		//fmt.Printf("OP: %#v\n", op)
		if op.Code == 1 {
			// Add
			val := vm.GetVal(op.Params[0].Val, op.Params[0].Mode) + vm.GetVal(op.Params[1].Val, op.Params[1].Mode)
			vm.SetVal(op.Params[2].Val, op.Params[2].Mode, val)
			vm.Ip += op.Length
		} else if op.Code == 2 {
			// Multiply
			val := vm.GetVal(op.Params[0].Val, op.Params[0].Mode) * vm.GetVal(op.Params[1].Val, op.Params[1].Mode)
			vm.SetVal(op.Params[2].Val, op.Params[2].Mode, val)
			vm.Ip += op.Length
		} else if op.Code == 3 {
			// Input
			vm.SetVal(op.Params[0].Val, op.Params[0].Mode, getInput())
			// fmt.Println("Provided input", input)
			vm.Ip += op.Length
		} else if op.Code == 4 {
			// Output
			sendOutput(vm.GetVal(op.Params[0].Val, op.Params[0].Mode))
			vm.Ip += op.Length
		} else if op.Code == 5 {
			// jump-if-true
			if vm.GetVal(op.Params[0].Val, op.Params[0].Mode) != 0 {
				vm.Ip = vm.GetVal(op.Params[1].Val, op.Params[1].Mode)
			} else {
				vm.Ip += op.Length
			}
		} else if op.Code == 6 {
			// jump-if-false
			if vm.GetVal(op.Params[0].Val, op.Params[0].Mode) == 0 {
				vm.Ip = vm.GetVal(op.Params[1].Val, op.Params[1].Mode)
			} else {
				vm.Ip += op.Length
			}
		} else if op.Code == 7 {
			// less than
			if vm.GetVal(op.Params[0].Val, op.Params[0].Mode) < vm.GetVal(op.Params[1].Val, op.Params[1].Mode) {
				vm.SetVal(op.Params[2].Val, op.Params[2].Mode, 1)
			} else {
				vm.SetVal(op.Params[2].Val, op.Params[2].Mode, 0)
			}
			vm.Ip += op.Length
		} else if op.Code == 8 {
			// equals
			if vm.GetVal(op.Params[0].Val, op.Params[0].Mode) == vm.GetVal(op.Params[1].Val, op.Params[1].Mode) {
				vm.SetVal(op.Params[2].Val, op.Params[2].Mode, 1)

			} else {
				vm.SetVal(op.Params[2].Val, op.Params[2].Mode, 0)
			}
			vm.Ip += op.Length
		} else if op.Code == 9 {
			// adjust relative base
			vm.RelativeBase += vm.GetVal(op.Params[0].Val, op.Params[0].Mode)
			vm.Ip += op.Length
		} else if op.Code == 99 {
			// Halt
			vm.Halted = true
			//fmt.Println("Halt")
		} else {
			fmt.Printf("Unknown opcode %d\n", op.Code)
			vm.Halted = true
		}
	}
}

func LoadString(numsStr string) []int {
	words := strings.Split(numsStr, ",")
	nums := []int{}
	for _, word := range words {
		num, err := strconv.Atoi(strings.Trim(word, "\n"))
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func LoadFile(fileName string) []int {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return LoadString(string(buf))
}

type Tile rune

const TileEmpty Tile = '.'
const TileWall Tile = '#'
const TileUnknown Tile = ' '
const TileOxygenSystem Tile = 'O'

type Direction struct {
	Code   int
	Name   rune
	Di, Dj int
}

var North Direction = Direction{1, 'N', -1, 0}
var South Direction = Direction{2, 'S', 1, 0}
var West Direction = Direction{3, 'W', 0, -1}
var East Direction = Direction{4, 'E', 0, 1}
var Directions []Direction = []Direction{
	North, South, West, East,
}

func DirectionFromDiDj(di, dj int) Direction {
	if di == North.Di && dj == North.Dj {
		return North
	}
	if di == South.Di && dj == South.Dj {
		return South
	}
	if di == West.Di && dj == West.Dj {
		return West
	}
	if di == East.Di && dj == East.Dj {
		return East
	}
	panic("fuck")
}

type State struct {
	DroidStartI int
	DroidStartJ int
	DroidI      int
	DroidJ      int
	DroidDir    Direction
	OxygenI     int
	OxygenJ     int
	Map         [][]Tile
}

func NewState() *State {
	s := State{}
	mapHeight := 43
	mapWidth := 43
	s.DroidStartI = 21
	s.DroidStartJ = 22
	s.DroidI = s.DroidStartI
	s.DroidJ = s.DroidStartJ
	s.Map = make([][]Tile, mapHeight)
	for i := range s.Map {
		s.Map[i] = make([]Tile, mapWidth)
		for j := range s.Map[i] {
			s.Map[i][j] = TileUnknown
		}
	}
	s.Map[s.DroidI][s.DroidJ] = TileEmpty
	return &s
}

func (s *State) Update(status int) {
	if s.DroidI == 0 {
		s.PrintMap()
		log.Fatal("move droid down")
	}
	if s.DroidJ == 0 {
		s.PrintMap()
		log.Fatal("move droid right")
	}
	if s.DroidI == len(s.Map)-1 {
		s.PrintMap()
		log.Fatal("increase map height")
	}
	if s.DroidJ == len(s.Map[0])-1 {
		s.PrintMap()
		log.Fatal("increase map width")
	}
	switch status {
	case 0: // hit a wall
		s.Map[s.DroidI+s.DroidDir.Di][s.DroidJ+s.DroidDir.Dj] = TileWall
		s.DroidI += 0
		s.DroidJ += 0
	case 1: // moved one step, found empty space
		s.Map[s.DroidI+s.DroidDir.Di][s.DroidJ+s.DroidDir.Dj] = TileEmpty
		s.DroidI += s.DroidDir.Di
		s.DroidJ += s.DroidDir.Dj
	case 2: // moved one step, found oxygen system
		s.Map[s.DroidI+s.DroidDir.Di][s.DroidJ+s.DroidDir.Dj] = TileOxygenSystem
		s.DroidI += s.DroidDir.Di
		s.DroidJ += s.DroidDir.Dj
		s.OxygenI = s.DroidI
		s.OxygenJ = s.DroidJ
	}
}

func (s *State) DirectionToUnknown(debug bool, tile Tile) (Direction, int, error) {
	type Pos struct{ I, J int }

	visited := map[Pos]bool{
		Pos{s.DroidI, s.DroidJ}: true,
	}

	queue := []Pos{
		Pos{s.DroidI, s.DroidJ},
	}

	prev := map[Pos]Pos{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		// visited[curr] = true
		if !(curr.I == s.DroidI && curr.J == s.DroidJ) && s.Map[curr.I][curr.J] == tile {
			// Backtrack to start find direction to go
			dist := 0
			for {
				dist++
				if prev[curr].I == s.DroidI && prev[curr].J == s.DroidJ {
					dir := DirectionFromDiDj(curr.I-s.DroidI, curr.J-s.DroidJ)
					return dir, dist, nil
				}
				curr = prev[curr]
			}
		}
		neighbors := []Pos{
			Pos{curr.I - 1, curr.J},
			Pos{curr.I + 1, curr.J},
			Pos{curr.I, curr.J - 1},
			Pos{curr.I, curr.J + 1},
		}
		for _, n := range neighbors {
			if !visited[n] && s.Map[n.I][n.J] != TileWall {
				visited[n] = true
				queue = append(queue, n)
				prev[n] = curr
			}
		}
	}

	return Direction{}, -1, errors.New("already explored everything")
}

func (s *State) PrintMap() {
	fmt.Printf("---\n")
	for i, row := range s.Map {
		for j, tile := range row {
			if j == s.DroidJ && i == s.DroidI {
				fmt.Printf("D")
			} else {
				fmt.Printf("%c", tile)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("---")
}

func (s *State) MapCopy() [][]Tile {
	c := make([][]Tile, len(s.Map))
	for i := range c {
		c[i] = make([]Tile, len(s.Map[0]))
		copy(c[i], s.Map[i])
	}
	return c
}

func (s *State) Flood() int {
	count := 0
	for {
		filled := false
		prev := s.MapCopy()
		for i := range s.Map {
			for j := range s.Map[i] {
				if prev[i][j] != TileEmpty {
					continue
				}
				if prev[i-1][j] == TileOxygenSystem {
					s.Map[i][j] = TileOxygenSystem
					filled = true
				}
				if prev[i+1][j] == TileOxygenSystem {
					s.Map[i][j] = TileOxygenSystem
					filled = true
				}
				if prev[i][j-1] == TileOxygenSystem {
					s.Map[i][j] = TileOxygenSystem
					filled = true
				}
				if prev[i][j+1] == TileOxygenSystem {
					s.Map[i][j] = TileOxygenSystem
					filled = true
				}
			}
		}
		if filled {
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	vm := NewVm(LoadFile("input.txt"))

	// Create state
	state := NewState()

	getInput := func() int {
		dir, _, err := state.DirectionToUnknown(false, TileUnknown)
		if err != nil {
			vm.Halted = true
			return East.Code // random direction, chosen by fair dice roll
		}
		state.DroidDir = dir
		return dir.Code
	}

	i := 0

	sendOutput := func(status int) {
		i++
		state.Update(status)
		if i%5 == 0 {
			time.Sleep(500 * time.Millisecond)
			state.PrintMap()
		}
	}

	vm.Run(getInput, sendOutput)
	state.DroidI = state.DroidStartI
	state.DroidJ = state.DroidStartJ
	state.PrintMap()
	_, dist, err := state.DirectionToUnknown(false, TileOxygenSystem)
	if err == nil {
		fmt.Println("Part 1:", dist)
	}

	count := state.Flood()
	fmt.Println("Part 2:", count)

}
